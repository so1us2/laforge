package cli

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/hashicorp/packer/packer"
	"github.com/sirupsen/logrus"
)

type loggerWriter struct {
	logger *logrus.Entry
}

type loggerClient struct {
	logger *logrus.Entry
}

func (w *loggerWriter) Write(p []byte) (n int, err error) {
	w.logger.Info(string(p))
	return len(p), nil
}

func NewUI(progname string) packer.Ui {
	log.SetOutput(ioutil.Discard)
	return &TaskUI{
		Name: progname,
	}
	// return &packer.ColoredUi{
	// 	Color:      packer.UiColorCyan,
	// 	ErrorColor: packer.UiColorYellow,
	// 	Ui: &packer.TargetedUI{
	// 		Target: progname,
	// 		Ui: &packer.BasicUi{
	// 			Reader:      os.Stdin,
	// 			Writer:      intLogger.writer,
	// 			ErrorWriter: intLogger.writer,
	// 			StackableProgressBar: packer.StackableProgressBar{
	// 				Bar: packer.BasicProgressBar{},
	// 			},
	// 		},
	// 	},
	// }
}

type TaskUI struct {
	Name string
	packer.StackableProgressBar
}

// Ask implements the Ui interface
func (t *TaskUI) Ask(msg string) (string, error) { return "", errors.New("user input not implemented") }

// Say implements the Ui interface
func (t *TaskUI) Say(msg string) {
	Logger.Infof("%s => %s", Boldwhite("%s", t.Name), Boldgreen("%s", msg))
}

// Message implements the Ui interface
func (t *TaskUI) Message(msg string) {
	// Logger.Debugf("%s => %s", Boldwhite("%s", t.Name), Boldcyan("%s", msg))
}

// Error implements the Ui interface
func (t *TaskUI) Error(msg string) {
	Logger.Errorf("%s => %s", Boldwhite("%s", t.Name), Boldyellow("%s", msg))
}

// Machine implements the Ui interface
func (t *TaskUI) Machine(m1 string, ms ...string) { return }

// ProgressBar implements the Ui interface
func (t *TaskUI) ProgressBar() packer.ProgressBar { return &t.StackableProgressBar }