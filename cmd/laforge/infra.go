package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gen0cide/laforge/builder/buildutil/templates"
	"github.com/masterzen/winrm"

	"github.com/emicklei/dot"
	"github.com/fatih/color"

	"github.com/gen0cide/laforge/core"
	lfcli "github.com/gen0cide/laforge/core/cli"
	"github.com/urfave/cli"
)

var (
	shouldgraph  = false
	infraCommand = cli.Command{
		Name:      "infra",
		Usage:     "Manage infrastructure deployment that has been generated with Laforge.",
		UsageText: "laforge infra",
		Subcommands: []cli.Command{
			{
				Name:   "plan",
				Usage:  "Show the delta between current deployment and the final desired state.",
				Action: performplan,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:        "graph",
						Usage:       "graphs the output in dot format",
						Destination: &shouldgraph,
					},
				},
			},
			{
				Name:            "status",
				Usage:           "Show the current build's infrastructure status.",
				Action:          performinfra,
				SkipFlagParsing: true,
			},
			{
				Name:            "apply",
				Usage:           "Provision the infrastructure to bring state in line with build blueprint.",
				Action:          performapply,
				SkipFlagParsing: true,
			},
			{
				Name:            "taint",
				Usage:           "Mark a host for re-provisioning in the laforge infrastructure state.",
				Action:          performtaint,
				SkipFlagParsing: true,
			},
			{
				Name:            "destroy",
				Usage:           "Destroy the builds infrastructure and clean the state.",
				Action:          performdestroy,
				SkipFlagParsing: true,
			},
			{
				Name:            "run",
				Usage:           "Run a host provisioner on a specific host within the infrastructure (usually for debugging).",
				Action:          performinfrarun,
				SkipFlagParsing: true,
			},
			{
				Name:            "exec",
				Usage:           "Run commands on a wildcard matched subset of hosts (usually for debugging).",
				Action:          performinfraexec,
				SkipFlagParsing: true,
			},
			{
				Name:            "graph",
				Usage:           "Generate a proposed DOT diagram of the target state.",
				Action:          performinfragraph,
				SkipFlagParsing: true,
			},
		},
	}
)

func performtaint(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	for _, x := range c.Args() {
		obj, exists := state.Current.Metastore[x]
		if !exists {
			cliLogger.Warnf("Node %s did not exist in the persisted snapshot.")
			continue
		}

		obj.Checksum = 666
	}

	plan, err := state.CalculateDelta()
	if err != nil {
		return err
	}

	plan.Base = state.Base

	tfcmds, err := core.CalculateTerraformNeeds(plan)
	if err != nil {
		return err
	}

	lfcli.SetLogLevel("info")

	_ = tfcmds

	for _, x := range plan.GlobalOrder {
		if core.IsGlobalType(x) {
			cliLogger.Debugf("%s is a global type. Not tainting...")
			continue
		}
		metaobj, found := plan.Graph.Metastore[x]
		if !found {
			cliLogger.Errorf("Node %s was not found on the graph...", x)
			continue
		}
		rev := metaobj.ToRevision().Taint()
		taintfile := rev.AbsPath(state.Base.BaseDir)
		if core.TypeByPath(x) == core.LFTypeEnvironment {
			taintfile = filepath.Join(state.Base.CurrentBuild.Dir, taintfile)
		}
		if _, err := os.Stat(taintfile); err == nil {
			cliLogger.Infof("Tainting exiting node: %s", x)
			err = ioutil.WriteFile(taintfile, []byte(rev.ToJSONString()), 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func performplan(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	plan, err := state.CalculateDelta()
	if err != nil {
		return err
	}

	plan.Base = state.Base

	tfcmds, err := core.CalculateTerraformNeeds(plan)
	if err != nil {
		return err
	}

	lfcli.SetLogLevel("info")

	for tid, cmds := range tfcmds {
		cliLogger.Infof("Terraform Commands For Team: %s", tid)
		for _, c := range cmds {
			fmt.Printf("  $ terraform %s\n", c)
		}
	}

	depthoffset := 0
	for _, k := range plan.OrderedPriorities {
		if len(plan.TasksByPriority[k]) < 1 {
			continue
		}
		cliLogger.Infof("Step #%d:", depthoffset)
		for idx, item := range plan.TasksByPriority[k] {
			tcol := ""
			tt := plan.TaskTypes[item]
			switch tt {
			case "REBUILD":
				tcol = color.HiMagentaString("[%s]", tt)
			case "MODIFY":
				tcol = color.HiYellowString("[%s]", tt)
			case "DELETE":
				tcol = color.HiRedString("[%s]", tt)
			case "TOUCH":
				tcol = color.HiCyanString("[%s]", tt)
			case "CREATE":
				tcol = color.HiGreenString("[%s]", tt)
			default:
				tcol = "[UNKNOWN]"
			}
			fmt.Printf("%s  %d) %s\n", tcol, idx, item)
		}
		depthoffset++
	}

	return nil
}

func performinfragraph(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	plan, err := state.CalculateDelta()
	if err != nil {
		return err
	}

	plan.Base = state.Base

	tfcmds, err := core.CalculateTerraformNeeds(plan)
	if err != nil {
		return err
	}

	_ = tfcmds

	snap := state.Current
	defer state.DB.Close()

	snap.AltGraph.Remove("root")
	snap.AltGraph.TransitiveReduction()
	nodemap := map[string]dot.Node{}
	g := snap.AltGraph

	di := dot.NewGraph(dot.Directed)
	di.Attr("nodesep", "0.2")
	di.Attr("compound", "true")
	di.Attr("rank", "min")
	di.Attr("rankdir", "LR")
	di.Attr("dpi", "72")
	di.Attr("smoothType", "graph_dist")
	di.Attr("mode", "hier")
	di.Attr("splines", "spline")
	di.Attr("decoreate", "true")
	di.Attr("overlap", "false")
	di.Attr("model", "subset")
	di.Attr("K", "0.6")
	di.Attr("fontname", "Helvetica")

	for _, x := range g.Vertices() {
		id := x.(string)
		nodemap[id] = di.Node(id)
		meta, ok := snap.Metastore[id]
		if !ok {
			panic(fmt.Errorf("could not find dependency for %s", id))
		}
		nodemap[id].Attr("style", meta.Style())
		nodemap[id].Attr("shape", meta.Shape())
		nodemap[id].Attr("height", "0.1")
		nodemap[id].Attr("label", []byte(meta.Label()))
		nodemap[id].Attr("fillcolor", meta.FillColor())
		nodemap[id].Attr("fontname", "Helvetica")
	}

	for _, x := range g.Edges() {
		src := x.Source().(string)
		tar := x.Target().(string)
		nodemap[src].Edge(nodemap[tar])
	}

	graphstring := di.String()
	wat := strings.Replace(graphstring, `"<`, `<<`, -1)
	wat = strings.Replace(wat, `\"`, `"`, -1)
	wat = strings.Replace(wat, `>"`, `>>`, -1)
	fmt.Println(wat)
	return nil
}

func performapply(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	plan, err := state.CalculateDelta()
	if err != nil {
		return err
	}

	plan.Base = state.Base

	err = plan.Preflight()
	if err != nil {
		return err
	}

	err = plan.SetupTasks()
	if err != nil {
		return err
	}

	diags := plan.Execute()
	if diags.HasErrors() {
		return diags.Err()
	}

	defer state.DB.Close()

	err = state.PersistSnapshot(state.Current)
	if err != nil {
		return err
	}

	return nil
}

func performtf(c *cli.Context) error {
	return commandNotImplemented(c)
}

func performinfrarun(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	plan := core.NewEmptyPlan()
	plan.Graph = state.Current
	plan.Base = state.Base

	// steps := []string{}
	for _, x := range c.Args() {
		if _, found := plan.Graph.Metastore[x]; !found {
			cliLogger.Errorf("%s not found in current graph.", x)
			continue
		}
		if core.TypeByPath(x) != core.LFTypeProvisioningStep {
			cliLogger.Errorf("%s is not of type provisioning_step.", x)
			continue
		}
		plan.Tainted[x] = true
		plan.GlobalOrder = append(plan.GlobalOrder, x)
	}

	err = plan.SetupTasks()
	if err != nil {
		return err
	}

	for k := range plan.Tainted {
		diags := plan.Orchestrator(k)
		if diags.HasErrors() {
			cliLogger.Errorf("Error in executing step %s: %v", k, diags)
		}
	}

	return nil
}

func performinfraexec(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	match := c.Args().Get(0)
	if match == "" {
		return errors.New("cannot proceed without a path matcher")
	}

	plan := core.NewEmptyPlan()
	plan.Graph = state.Current
	plan.Base = state.Base

	conns := []*core.Connection{}

	// steps := []string{}
	// for _, x := range c.Args() {
	// 	if _, found := plan.Graph.Metastore[x]; !found {
	// 		cliLogger.Errorf("%s not found in current graph.", x)
	// 		continue
	// 	}
	// 	if core.TypeByPath(x) != core.LFTypeConnection {
	// 		cliLogger.Errorf("%s is not of type connection.", x)
	// 		continue
	// 	}
	// 	plan.Tainted[x] = true
	// 	plan.GlobalOrder = append(plan.GlobalOrder, x)
	// }

	lfcli.SetLogLevel("info")

	for _, obj := range plan.Graph.Metastore {
		if obj.ObjectType != core.LFTypeConnection {
			continue
		}

		connObj, ok := obj.Dependency.(*core.Connection)
		if !ok {
			continue
		}

		parentID := connObj.ParentLaforgeID()

		if matched, err := path.Match(match, parentID); err == nil && matched {
			conns = append(conns, connObj)
		}
	}

	cnt, err := templates.NewContext(
		state.Base,
		state.Base.CurrentEnv,
		state.Base.CurrentBuild,
	)
	if err != nil {
		return err
	}

	newlib := templates.NewLibrary()
	if c.Args().Get(1) == "" {
		return errors.New("you must specify a command to run")
	}

	_, err = newlib.AddBook("command", []byte(c.Args().Get(1)))
	if err != nil {
		return err
	}

	wg := new(sync.WaitGroup)
	for _, y := range conns {
		wg.Add(1)
		go func(x *core.Connection) {
			defer wg.Done()
			dup := cnt.Clone()
			dup.Attach(
				x,
				x.Network,
				x.ProvisionedHost,
				x.ProvisionedNetwork,
				x.ProvisionedNetwork.Network,
				x.Team,
				x.Host,
			)
			cliLogger.Infof("Execution Target: %s", x.Path())
			renderedCommand, err := newlib.Execute("command", dup)
			if err != nil {
				cliLogger.Errorf("Error rendering template: %v", err)
				return
			}

			cmd := core.NewRemoteCommand()
			cmd.Init()
			cmd.Command = string(renderedCommand)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr

			baseConfig, err := core.LocateBaseConfig()
			if err != nil {
				return
			}

			baseDir := filepath.Dir(baseConfig)

			conn := &core.Connection{}
			err = core.LoadHCLFromFile(fmt.Sprintf("%s.laforge", filepath.Join(baseDir, x.Path())), conn)
			if err != nil {
				cliLogger.Errorf("Error loading job %s resource: %v", x.Path(), err)
				return
			}

			if conn.Active != true {
				cliLogger.Errorf("Host %s is not active :(", x.ParentLaforgeID())
				return
			}

			newConn, err := core.SmartMerge(x, conn, false)
			if err != nil {
				cliLogger.Errorf("Error merging connections for %s", x.ParentLaforgeID())
				return
			}

			connObj := newConn.(*core.Connection)

			if connObj.IsSSH() {
				connObj.SSHAuthConfig.IdentityFile = filepath.Join(baseDir, "envs", state.Base.CurrentEnv.Base(), state.Base.CurrentBuild.Base(), "data", "ssh.pem")
			}
			if connObj.IsWinRM() {
				cmd.Command = winrm.Powershell(cmd.Command)
			}
			cliLogger.Infof("Connecting Host: %s", connObj.RemoteAddr)
			err = connObj.ExecuteCommand(cmd)
			if err != nil {
				cliLogger.Errorf("Error Executing Command: %v", err)
				return
			}

		}(y)
	}

	wg.Wait()
	// err = plan.SetupTasks()
	// if err != nil {
	// 	return err
	// }

	// for k := range plan.Tainted {
	// 	diags := plan.Orchestrator(k)
	// 	if diags.HasErrors() {
	// 		cliLogger.Errorf("Error in executing step %s: %v", k, diags)
	// 	}
	// }

	return nil
}

func performinfra(c *cli.Context) error {
	return commandNotImplemented(c)
}

func performdestroy(c *cli.Context) error {
	state, err := core.BootstrapWithState(true)
	if err != nil {
		return err
	}
	if state == nil {
		return errors.New("cannot proceed with a nil state")
	}

	plan, err := state.CalculateDelta()
	if err != nil {
		return err
	}

	plan.Base = state.Base

	err = plan.BurnIt()
	if err != nil {
		return err
	}

	defer state.DB.Close()

	err = state.PersistSnapshot(state.Current)
	if err != nil {
		return err
	}

	return nil
}
