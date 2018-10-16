// Code generated by fileb0x at "2018-10-16 00:59:39.410021234 -0700 PDT m=+0.009297171" from config file "assets.toml" DO NOT EDIT.
// modification hash(c1b22433ad76d522dac341bf18bff4ed.e5b9c5ef4c0b7aef8593382d0449dfd6)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileCommandTfTmpl is "command.tf.tmpl"
var FileCommandTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x93\x31\x8f\xd4\x30\x10\x85\x7b\xff\x8a\x91\x45\x01\x12\x44\x14\x88\xee\x0a\xb8\x02\xe8\x10\x14\x14\xe8\x64\xe5\xe2\xd9\xdb\x11\xf1\x4c\xe4\x99\xec\xde\x29\xca\x7f\x47\x93\x90\xb0\xc5\x22\x41\x81\x9b\x64\xf4\x66\x9e\x9f\x3f\xd9\xd3\x04\x19\x0f\xc4\x08\xb1\x93\x52\x5a\xce\x11\xe6\x39\x54\x54\x19\x6b\x87\x10\x79\xec\xfb\xb4\x95\x11\xe2\x50\xe5\x44\x4a\xc2\x69\x9a\xa0\xf9\x80\x06\x71\x53\x13\xb7\x05\x7d\x3c\xa9\xe1\xb0\xcb\x5e\x24\x1e\xcb\x3d\x56\x17\x23\x4c\x01\x20\xe3\x80\x9c\x35\x09\xc3\x0d\x7c\x0f\x00\x00\x91\xee\x4b\xea\xa4\x0c\xa3\x61\x3a\x95\x44\xac\xd6\x72\x87\xcd\x9f\x37\x8a\x01\xe0\x2e\x04\x80\x3d\x15\x56\x6f\x2b\x62\xf8\x0a\x1f\xb1\x5b\x37\x03\x98\x26\xa0\x03\x34\x1f\x45\xad\xf9\xa4\xdf\x88\xb3\x9c\xd5\x0f\x0a\xcb\xea\x84\x19\x3b\x23\xe1\x5f\xfd\xbe\x8e\xa2\xb6\xfc\xdc\x40\x7c\x36\xfd\x7b\xb8\x86\x86\xd3\x9b\xd4\xe6\x5c\x51\x75\x89\xba\x2e\x7b\x1a\x70\xf3\x3d\x13\xd7\xf2\x5b\x1a\x15\xeb\x26\xbd\xcb\x85\x98\xd4\x6a\x6b\x52\x2f\xa6\xa9\xa0\x8c\xb6\xb4\xbc\x7d\x7d\x31\x3b\xb4\xaa\x67\xa9\xd9\x05\x0f\x75\x2b\x65\x40\x23\x3f\x54\xf3\x45\xc4\x3e\x6f\xfa\xbc\x67\x99\x37\x36\xd8\x2b\xfe\x2d\x8d\xff\x08\x64\xb5\x56\x3d\x5e\x21\xb2\x6a\x55\xc4\xae\xb0\x80\x2b\x38\x2a\x9d\x5a\xc3\xf4\x03\x9f\xd6\xbc\x07\xea\xf1\xb9\x93\x21\xce\xf8\x08\xcd\xfb\x91\xfa\xdc\xdc\x0a\x1f\xe8\xc1\xc3\xf6\x49\xf5\x98\x2e\xc6\x92\x4f\x2c\xb7\xec\xc5\x15\x62\xec\x20\xc3\x52\x12\xf7\xfe\x80\xb6\x7b\x0c\x3b\x7f\x7f\x4e\xdb\xf7\xab\x55\xe2\x07\x77\x7b\xb9\x74\xdd\x05\x77\x9b\xc3\xee\xf5\x33\x00\x00\xff\xff\x0d\x32\x43\xcf\x8a\x03\x00\x00")

// FileDNSRecordTfTmpl is "dns_record.tf.tmpl"
var FileDNSRecordTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x92\x3f\x4b\x04\x31\x10\xc5\xfb\x7c\x8a\x47\xb0\x3c\xc2\x81\xf5\x75\x82\xd8\x5c\xa1\x60\xa1\x48\x58\x37\x73\x6b\x30\x37\x59\x92\xac\x82\x21\xdf\x5d\xb2\xec\x9f\x3b\x0e\x1b\xb5\xcc\xbc\xc7\x9b\x99\xdf\x24\x67\x18\x3a\x58\x26\x48\xc3\x51\x07\x6a\x7d\x30\x12\xa5\x08\x11\x28\xfa\x21\xb4\x93\xd2\x4c\x9a\x8e\x94\x24\x64\x1f\xfc\x87\x8d\xd6\xb3\xce\x19\xea\x96\x12\xe4\xec\xd7\xdc\x1c\xa9\x46\xe8\x98\xa8\x5f\xe4\xfa\xd0\x3c\x1c\x5f\x29\x54\x51\x22\x0b\xe0\xcb\x33\x61\x07\x59\x5d\x37\xfb\x87\xfb\xb1\x85\x7a\xaa\xd5\x52\x94\x14\x40\x0d\xbb\x74\xec\x6b\xb5\x94\x6a\x68\x8c\x09\x14\x23\x45\xec\xf0\x2c\x00\x20\x67\xd8\xc3\xa9\xfb\x8e\xdf\x28\xd8\x44\xa6\xee\x55\x1d\xf2\x2a\x77\xde\x77\x8e\x74\xeb\x8f\xfd\x90\x48\x5b\x8e\xa9\xe1\x96\xd4\xcf\xeb\x28\xa6\xf4\xe9\xc3\xbb\xb6\x9c\x28\x1c\x9a\x96\xd4\x76\xad\xf5\x45\x6e\xe6\xee\xe4\x22\x2d\xad\xce\x07\x7f\x6c\xdc\x30\x4e\xbe\x9a\x79\x1a\xeb\x45\x00\x29\x39\xec\x70\xbd\xdd\x8a\xb3\x03\xf0\xe0\x9c\x9e\x9f\xff\x47\x3f\x05\xdb\x75\x14\xe2\xf8\x00\x66\x06\xda\x9a\x4a\xfc\x57\x8c\x4e\x32\xc6\xeb\x14\x21\x00\x43\x3d\xb1\x89\xda\xf3\x72\x23\xb9\x64\xac\xa2\x5c\xb9\x5c\x7c\x39\xf5\xf7\x95\x37\x23\xe3\x22\xc4\x0a\xfd\x3b\x00\x00\xff\xff\xf5\xf2\x7c\x23\xfe\x02\x00\x00")

// FileInfraTfTmpl is "infra.tf.tmpl"
var FileInfraTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xdc\x5a\xeb\x6f\xdb\x38\x12\xff\xae\xbf\x62\xa0\xa6\xb8\x06\xa8\xe4\x3c\x16\xd9\x36\xd8\x00\x97\x36\x69\x37\xb8\x5e\x72\x48\xb3\xb7\x1f\x76\x0b\x81\x16\xc7\x0e\x37\x12\xa9\x25\x29\xa7\x8e\xa1\xff\xfd\x40\x52\xd4\xc3\xaf\xbc\x9c\x5b\xdc\xf9\x43\x6b\x8b\x33\xc3\x79\xfe\x66\x48\xe5\xd5\xd3\x3f\xc1\x2b\xf8\x72\xfc\xe9\xe2\xf2\xf3\x29\x7c\x3e\x3d\x3f\xbd\x3c\xbe\x3a\x3d\x81\xab\xd3\xcb\x4b\xf3\xf0\x9f\xf0\xf1\xe2\xfc\xd3\xd9\xe7\x5f\x2e\x8f\xaf\xce\x2e\xce\x83\x57\x10\x45\xf0\xeb\xf1\xe5\xf9\xd9\xf9\x67\x88\xa2\xe0\x15\x5c\x5d\x33\x05\x23\x96\x21\x30\x05\xa4\xd4\x22\x27\x9a\xa5\x24\xcb\xa6\x30\x46\x8e\x92\x68\xa4\x31\x9c\x08\xe0\x42\x03\x52\xa6\x81\xe9\xbf\xa9\xe0\x15\xa4\x82\x6b\xe4\x5a\x01\x65\x12\x53\x9d\x4d\x63\xf8\x45\x21\x7c\x21\x23\x21\xc7\x08\x84\x53\x90\x08\xc3\x92\x65\x14\xb4\xdf\x24\x0e\x9e\x63\x69\xa0\x51\x4a\x23\x3f\x87\x59\x00\x30\x24\xe9\x0d\x72\x0a\x21\xea\x94\x4e\xf6\x43\xfb\x10\xa0\x90\x38\x62\xdf\xe1\x08\xc2\xd9\x0c\xb6\xe2\x2b\x24\x79\x7c\x76\x02\x55\x35\x68\xd8\x63\x3d\x52\x9a\x68\x0c\x2d\x03\x72\x5a\x08\x66\x4c\x39\x82\xdf\xec\x13\x80\xf0\x5a\xeb\x42\x1d\x0e\x06\xb3\x19\x30\x4e\xf1\x3b\x6c\xc5\x1f\x8c\x2d\xf1\x47\xc1\x47\x6c\xec\x36\x4d\x72\xa2\x34\xca\x10\xaa\x2a\x7c\xfb\x38\x4e\x95\x91\x09\x76\x18\xbf\xd9\x7f\x4b\x85\x92\x93\x1c\x6b\xed\xd7\xf0\x7b\x4a\x2b\xc2\xd9\x4d\x94\xba\x15\x92\xde\xcf\xeb\x29\x3d\x6f\x15\x54\x41\x30\x21\x92\x91\x61\x86\x10\x4e\x72\xc5\xee\xd0\xb9\x53\x4f\x0b\xab\x4c\x4e\x0a\x43\x49\x71\x44\xca\x4c\xc3\x51\xed\xeb\x50\xe5\x24\xcb\x42\x43\xc1\x77\x23\xa5\x09\xa7\x44\xd2\x68\xd7\x69\x14\xe6\x48\x59\x99\x2f\x2c\xef\xd5\xcb\x19\x91\x63\x5c\x58\xfd\xa1\x5e\xfd\xbe\x7c\x79\x7f\x6f\xa9\xca\x42\x3d\x4c\xe1\x72\x58\x72\x5d\xee\x1e\x58\xb9\xee\x47\x24\x54\x94\x66\xa2\xa4\x83\xfa\xf7\xee\xc1\xce\x0f\x51\xa6\x55\xd8\x63\x79\xb7\x96\xe5\x5d\x8f\x25\x45\xae\x85\xfa\xd1\x72\xb8\xef\x35\x79\xfd\xe3\xc7\x1e\xdd\xc1\x4a\xba\x83\x9a\x8e\xe2\x90\x11\xbe\x56\x83\x9c\x71\x96\x93\x6c\x5e\x13\xc7\xf9\xde\x72\xba\xef\x35\x5b\xfd\xe3\x7d\x4d\x77\xbb\x77\xb3\x6f\x89\x6e\x19\xa7\xe2\xd6\x0b\xf7\xbf\xf6\x76\x76\x0f\xa2\x54\x48\x6c\xc9\xdf\xad\x25\xdf\x79\x17\xc9\xbd\x96\x78\x77\x6f\xbd\xf0\xbd\x3e\xf5\xc1\x3d\xaa\x34\x29\x50\x48\x31\x61\x14\x25\x84\x63\x21\xc6\x59\x9d\xb5\xa9\x44\x8a\x5c\x33\x92\x99\x9a\x0e\xb7\x66\x06\x7b\xde\xac\xae\x8a\x71\x5a\x24\x86\x27\x31\x74\xb6\x2a\xb6\x6d\x61\x14\x52\xfc\x81\xa9\x5e\x5b\x51\x86\xb7\xa6\xf3\xf5\x24\x71\xcc\x04\xbf\x97\xcb\x91\x39\xa6\x2a\x08\x0c\x5a\x51\xae\xee\xe0\xf0\x68\x0d\x13\xe5\x2a\xb9\x13\x1c\x13\x66\xcb\x37\x08\x28\xd1\xc4\x5b\x6f\x57\x73\xc2\xc9\x18\xa9\xa5\x0a\x1d\x08\x5a\xb1\x55\xe5\x9c\xd3\xc1\x97\x66\xc1\xec\x2f\x51\x89\x52\xa6\xd8\x08\x4b\x45\x5e\x94\x1a\x13\x8e\xfa\x56\xc8\x9b\x10\xc2\x49\x91\x2e\xca\x88\x4f\xf9\x84\x49\xc1\x73\xe4\x3a\xfe\x40\x14\x42\x55\x45\xba\xc5\x5e\xf3\xcf\x79\x99\x0f\x51\x9a\x05\x23\x22\x00\xdb\x69\x8c\xcb\x89\xc6\x44\x95\xc3\x7a\x0b\x13\xad\x11\xc9\x14\xae\xd5\x67\xc4\x24\xde\x5a\xd8\x09\x49\x96\x89\xdb\x84\xa5\x79\xf1\x5c\xbd\xac\xa4\xc8\x4a\x32\x72\x9c\x3e\x2e\x79\x96\xbb\x23\x9e\x14\x69\xac\x30\x1b\x25\x19\xe3\x37\x55\x18\x18\xab\x8c\x90\xa6\x0d\x09\x2d\x52\x91\x19\x19\x5e\x6c\x65\x88\x9c\x55\x89\x24\x7c\x8c\x6d\xcb\x59\x9d\x2a\x93\x22\x4d\x52\x46\xdb\x2e\xf3\xed\x71\xde\x21\x34\x67\x7c\x33\xee\x71\xa2\x5e\xd8\x3f\xab\x88\x74\x7a\x3f\x4d\x49\x9f\xee\x67\x6b\x5c\xc2\x0a\xe3\xe7\xc1\xfe\x5e\xdd\x97\x67\x33\xb0\x12\x60\x2b\x79\x0b\x5b\x26\x0e\xa6\x3e\xfb\xde\x3b\x36\x9c\x1f\xcf\x4e\x2e\x95\xa9\x48\xbf\x8b\x23\x6e\x1a\xfc\x6c\x66\x06\x0d\x47\xf0\xc8\x08\x4e\x28\x4b\xf0\xbb\xde\x4c\x0c\x27\x94\x45\x46\xd8\xe6\xa3\x58\x07\x08\xa0\x10\xb2\x3f\x4c\xed\xed\xb5\xe3\xd1\xfe\xfe\xbb\xf7\x9d\x99\xe7\x69\x25\x41\x59\x72\x7b\xcd\x34\x66\x4c\xe9\x4e\x5d\x98\xfe\x6f\x86\x06\x9d\x68\x32\xee\x0b\x7a\x2c\x4c\x51\xf6\x94\x5a\x33\x8a\x31\xbe\xc1\x48\x19\x61\xff\xa7\xf5\xb6\x04\xd7\xfe\xc2\xf8\x51\xe4\xd3\x24\x1f\xe7\x7a\x13\xf1\x33\xc2\x22\x23\xec\x59\xf1\xb3\xe3\x2b\x9f\x3e\xa8\xd2\x20\x7c\xff\xfe\xc7\xdd\xf0\x19\x25\xb5\xb4\xcb\x04\x2d\xfc\x71\xd4\xc6\x21\x6f\xed\xb7\x45\x0c\x3c\xe3\x69\x56\x52\xa4\xe7\xbe\x93\x57\x55\x30\x18\x74\x0c\x37\xbe\xaa\x65\x98\xb5\x95\x31\x69\xa7\x81\x7a\x78\xe1\xd8\x78\xfc\xb9\x71\x99\x17\x17\x00\xb0\xc2\x9a\xed\x1c\xe5\xa5\x1a\x1a\x83\xe7\x4f\x9f\xe7\x9e\x16\xf2\x47\x61\xcd\xcb\xfa\xc6\x0f\x44\x7c\x28\x4a\x4e\xff\xc7\x30\x68\x53\xe8\x01\x9b\x48\xb2\x4d\x43\xdb\x72\xf9\x55\xd0\x0e\x18\xfe\x28\x91\xf8\x64\x4a\x6c\x5a\x1c\x1e\x41\x18\xfa\xf5\xb6\xaa\xdf\xc2\xd6\xb5\x50\x5a\x2d\xd6\xf4\xcf\x42\xe9\x0f\xd3\xba\xa2\xdd\xe4\x52\x6f\x2e\x86\x7f\x18\xf2\x37\xbe\x18\xd6\x22\x81\x61\xd8\x6e\xd8\x3b\xd3\x94\xd9\xd6\xee\xea\xf6\xaf\x67\x27\x53\x63\x23\xc0\x3f\xdd\x63\x67\x66\xc8\xd5\xce\x6e\xe8\x29\x9c\x1a\x8b\x06\x1e\x41\x21\x19\xd7\x23\x08\x5f\xab\x48\xbf\xa6\xd1\x6b\x15\xbd\x56\xe1\x32\x07\x2f\x7a\xb6\x36\xab\x5e\x6e\xb7\x6e\xb5\x6a\xc6\xb7\xf6\x7b\xdf\xe7\x7f\xb5\x4f\xcd\x96\x8c\x7e\x5f\xe7\xd9\xad\x85\x94\x78\x19\x97\xd9\xfd\x56\x62\x19\xa1\x54\xa2\x52\x35\xc0\xf7\x55\xf2\x48\x66\x3e\x5d\x34\x5b\xa0\xb2\x34\xf7\xed\xc4\xb8\xd2\x84\xa7\xb8\x81\xad\x00\x72\x92\x5e\x33\x8e\x89\xbf\x63\xda\x9a\x4d\x88\x8c\xdd\x75\xd9\x6f\x96\xd3\xfa\xe0\xac\xde\xf3\x2b\xbb\xb3\xdc\xdf\x1a\x01\xe6\x48\x7e\x6f\x1b\x71\xe7\xf6\xca\xa1\xa7\xf9\x0c\x85\xd0\x09\x65\xea\xa6\x51\x16\x80\x71\xa6\x19\xc9\xd8\x1d\x26\x05\x91\x24\x57\x9d\x35\x00\xa3\x90\x37\xc6\xaa\x74\xc2\xd4\x4d\xec\xf5\xe9\x10\x7a\x4b\x0a\x1a\x29\x45\xbb\x2b\x2c\x27\xe3\x9e\x91\x42\x75\x4c\xbc\xf8\xda\x37\xcc\x04\xc2\xff\xef\x1d\xea\x52\xd5\x8c\x53\x28\x47\x24\xc5\x8e\x86\x6d\x8f\x5f\xda\x4d\xda\xe5\xb8\xad\x0d\x9f\x5a\xbd\xfe\xe2\x05\xd6\x09\xe5\xa4\x99\x7e\x6e\xb4\x7c\x13\x76\xb8\x7d\x3b\x7f\x0b\x8d\x11\x5f\x88\xd2\x17\xa9\x46\x0d\x55\xb5\xdd\xba\x1b\x80\xa4\x29\x2a\x95\xa4\x2e\x28\x5d\xc7\x72\xa2\x13\x56\x2c\x55\xba\x56\x21\x5e\x96\x40\x71\xbd\xb8\xce\x5f\x86\xaf\xa4\xaa\x0b\x01\x56\xcb\xaf\xa9\x64\x85\x56\xfd\x87\xff\x26\x52\x41\x58\x2a\x94\x09\x25\x9a\x24\xca\x12\x25\x8c\x86\xdb\xdb\x4d\xf9\x01\xe4\xa8\x89\xbd\x1f\x6a\x6d\x30\xec\xdd\x5c\xb7\xe2\x7e\xf6\x0f\xab\xaa\xeb\xf1\xf3\xce\xa3\xf8\xa3\xc8\x0b\xd4\x4c\x33\xc1\xe3\x93\xf3\xaf\xf1\xa5\x10\xfa\x44\xe4\x84\xf1\x5e\x4e\x39\xf4\xae\x0b\x41\xfd\xea\xee\xec\x5a\xe0\x06\xf0\xd7\x78\x4a\x13\xa9\xcb\x22\x72\xaa\x47\x85\xda\xb5\x37\xbd\xa8\xed\x25\x38\xd8\xd3\x38\x53\x5a\x12\x2d\x24\x0c\x48\xaa\xd9\x04\x0f\xa7\xa8\xda\xad\xd4\x54\x15\x12\x8b\x48\x15\x98\xba\x5a\x98\x13\xd6\xb9\xf7\xf3\x85\x76\x89\xd9\xb1\x52\xa8\x3f\x09\x69\x80\x6c\x25\x84\x99\x50\xc4\x5f\x6d\x10\x7d\x9f\xdd\xee\x5b\x89\x59\xdb\x1c\xac\x36\xea\xfa\x1f\x38\x35\x29\xf8\xd3\x4f\xa7\x17\x9f\x02\x29\x84\x3e\xdc\x9a\xa5\xd7\x22\x2f\xde\xdc\x73\xff\x28\x31\x4b\x94\xba\x4e\x8a\x72\x98\xb1\x34\xb9\xc1\x69\xe7\x26\x72\xbb\x02\x23\xeb\xef\x99\x7b\xb3\xe2\x04\x3f\xe4\x3a\xd2\x8b\xbc\xc1\xa9\x91\xe4\xa4\x58\xbf\xc6\x69\xa1\xd3\xb8\xae\xb1\xc0\x68\xdb\xb3\xcc\xf7\xb9\x3a\x3f\x83\x6e\x64\x39\x2e\x04\x57\xcb\xb2\xe3\x09\x9f\x73\x49\x1d\xe0\x3a\x37\x5f\x32\x1e\x9d\x2e\x5c\xc3\x5a\x77\xc4\x7a\xde\x98\xd5\x41\x9e\xe6\x0e\x03\x20\xbc\x7f\x75\xbe\xb0\xba\xeb\xab\x36\x9d\x97\xb1\x4c\xe3\x05\x9a\x6e\x61\xda\xf7\x5c\x4b\x35\x59\xc1\xbd\xd0\xea\xfc\xfa\x37\xef\x4b\x7b\xc3\xae\x98\xe0\x28\x21\x74\x59\x39\x7b\x78\xb5\x03\xa4\x82\x73\x4c\x8d\x76\x3d\x24\x75\x48\x64\xbf\x6c\x00\x4d\x9b\x5e\x56\xcb\xbb\x65\x5c\xe6\xfd\x65\x8b\x2a\xf5\xf2\x71\x17\x5c\xe6\xa4\xb0\x1c\x45\xa9\x2d\xd9\xc1\xce\x9c\x8c\xb9\x57\x6c\x7d\xdf\x1b\x40\xfc\x97\x27\xe8\xb7\xd9\x6a\x1d\x70\xac\x71\x11\x19\x23\xd7\xf5\xf7\x23\x08\xed\xd5\x78\xb8\xdc\x89\x2f\xe0\x47\x27\x52\xa9\xeb\x15\x8e\x74\xeb\x06\x55\x56\xb8\x10\x56\x78\x51\xb2\x09\xd1\x68\x60\xee\x41\x6f\x67\x1a\x74\x6c\xf9\xe6\x5f\xd4\x2c\x77\x74\x0f\x13\xfc\xe1\xd0\x6c\x19\x0f\x16\xeb\x77\x30\x5f\x2a\x83\x1a\x70\x23\x1b\x85\xf0\x31\x39\x4f\x51\x69\xc6\x89\xae\x6f\x0d\x3e\x1e\xfe\xfe\xfb\x6a\x61\x8b\xf9\x30\xc7\x3e\x10\x85\x5e\xa3\xcc\x1c\x5c\x77\x86\x0a\x51\xea\xa2\xd4\x8b\x20\x10\x7b\x4f\xb2\xa2\x2d\xe6\x09\xc9\x4a\x5c\x9a\x43\x7e\x86\x5e\x9e\x44\x0b\x73\x5e\xbc\xd3\x3e\x2b\xfa\x63\xfa\x6a\x85\x5c\xe3\x7b\x29\x7d\x7a\x03\x9d\xd1\xcf\x8e\x71\x7d\xdd\xba\x61\xbd\x22\xe3\xd3\x3f\x4b\x92\x29\x7b\xd5\x1c\x42\x68\x5a\x5c\x73\xfc\x5c\x38\x6c\xb8\x93\x68\x2a\x24\x4d\x14\xea\x7b\x8f\x1a\xdd\x17\x74\xce\x40\xd3\x2f\xe3\x15\xaf\xf0\xe2\xce\x7b\xba\xd8\xc8\x6a\x27\xd5\xc7\xce\x71\xab\xda\x4e\xbc\xaa\xd7\xc4\x8f\x50\xcd\x2c\xd6\xea\xf5\xcf\x16\xc7\xcd\x03\x6d\xef\x6b\xf6\x77\x76\x1a\x0b\xa4\x34\xf2\xe7\x7a\xf5\x0b\x06\xbc\xed\x6e\x6d\xb1\xcc\x21\x85\x7b\x9f\xaa\x31\x2f\x32\x62\xaf\xdd\xb2\xfb\x0f\x8f\x9e\xfa\x99\x43\x4e\xd8\xf6\x5a\x9a\xd8\xc7\xba\xc8\x1a\x94\x6b\xea\x42\x76\x0f\x7b\x12\x73\x51\xe3\xfc\x7f\xa3\x58\xcc\x27\x13\x29\xc9\x36\xb9\xe3\x02\x5c\xf8\xce\x96\xb8\xb1\xdf\xde\x04\x9a\x0a\xec\x58\xdd\xbf\xf7\x09\x1f\xa3\xc3\x23\xa0\xfc\x69\x0d\x7f\x09\xa8\x33\xfb\xf7\x08\xda\x35\xae\xb5\x57\x00\x6b\xbb\xdd\xc3\x70\xbf\x05\x28\x17\xaa\x35\x39\x9c\x34\x29\xc7\xf8\x38\xe9\x8f\x77\xf5\x5f\x78\x75\x20\xaa\x57\x15\xcb\xe3\x2b\x91\x53\x94\x48\x1b\x5d\x0d\xa9\x8f\xd2\xc3\xfa\xee\x42\x15\xd4\xbd\x2f\x6c\x8a\x76\xe9\x05\xdc\x7f\x02\x00\x00\xff\xff\xe5\x75\x34\x6c\x44\x27\x00\x00")

// FileProvisionedHostTfTmpl is "provisioned_host.tf.tmpl"
var FileProvisionedHostTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x92\x3f\x8b\xdc\x40\x0c\xc5\x7b\x7f\x0a\x61\xae\xde\xc0\xc1\x41\x52\x6c\x91\x4b\x02\x49\x93\x22\x04\x52\x9a\x89\x47\x8e\xc5\x79\x46\x66\xa4\xb5\x39\x9c\xf9\xee\xc1\x33\xfe\xb3\xe6\xdc\xc4\xdd\x3c\xfd\xf4\xf4\x24\xdc\x07\x1e\x48\x88\x3d\xda\xaa\x65\x51\x28\xa7\x09\x1e\x2e\x3f\xd1\xb8\xcb\xb7\xcf\x10\xe3\xbb\xf4\xfe\x8e\x3a\x72\x78\xb9\x3c\x1b\xc1\x4d\xfc\xca\xa2\xab\x52\xc2\x54\x00\xd4\xec\x7a\x54\x52\x62\x5f\x91\x85\x2b\x24\xee\xd3\xae\xce\x9e\x7f\xa1\xad\x3b\xd1\x40\xfe\x0f\xc4\x58\x00\xa0\x1f\x28\xb0\x77\xe8\xf5\xae\xeb\xcb\xae\x9e\x76\xfd\xbe\x51\x67\x67\x3e\x7f\x4b\xd7\xf3\xac\x9e\xf2\x8a\xc6\xed\xf8\xca\xaf\x8b\xbe\xc1\x7d\xde\x78\xed\x58\xf0\xf5\x0e\x67\x1d\xf3\xf9\xde\x0e\x48\x47\x3a\xc3\x03\xe6\xc3\x9f\xe5\xff\xb1\xd6\x12\x69\x6a\xa5\x01\x61\x37\x2e\x1f\xa6\x34\x2c\x17\x62\x99\xdc\x1c\x2b\x56\xc6\xda\xb0\x31\x77\x5a\x62\x3a\xae\x4d\xb7\x21\x99\xd9\xb5\xc5\x46\xf8\x16\x6a\xac\xbc\x71\xb8\xd9\xdc\x69\xb1\x2c\x8a\x02\xe6\xac\xd4\x6c\xdb\xc9\x2f\xf2\x96\x47\xc9\x71\x47\xf2\xc1\xa5\xdf\xe1\x18\xeb\x34\x13\x40\xcf\x41\xe1\x0a\x4f\x1f\xde\x3f\xa5\x77\xab\xda\x0b\x5c\xa1\x31\x9d\x60\x52\xe4\x85\xfa\x6a\xc0\x40\xcd\xeb\x41\xbf\x09\x26\xd7\x8f\xd6\x91\x27\xd1\x60\x94\xc3\x62\x6a\x44\x46\x0e\x36\x0f\x5d\x5f\x69\x62\xcc\xf1\xb1\x4b\xff\x6d\x01\x20\xd2\xfe\x77\xda\xc7\xc7\x43\x82\xc0\xac\xb9\x4e\x16\xbd\x92\xbe\x56\x0d\x75\x98\x4d\x0e\xd2\x21\x82\xb7\x73\x82\xf8\x2f\x00\x00\xff\xff\x2b\x90\x01\x25\x84\x03\x00\x00")

// FileRemoteFileTfTmpl is "remote_file.tf.tmpl"
var FileRemoteFileTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x94\xcf\x6e\xd4\x30\x10\xc6\xef\x79\x8a\x91\xd5\x03\x48\x95\xcb\x89\x03\xd2\x1e\x4a\x11\x7f\x2e\x08\xf5\xc2\x01\x21\xcb\xac\x67\xb3\x23\x12\x4f\xe4\x99\x74\xa9\xa2\xbc\x3b\xb2\xb3\x49\x53\x11\xe0\x42\x25\x7c\x8a\xf5\xf9\x9b\xcc\xfc\xc6\x9e\x61\x80\x80\x07\x8a\x08\x26\x61\xcb\x8a\xee\x40\x0d\x1a\x18\xc7\x2a\xa1\x70\x9f\xf6\x08\x26\xf6\x4d\xe3\xe6\xad\x01\xd3\x25\xbe\x23\x21\x8e\x6e\x18\xc0\xbe\x43\xcd\xe6\x49\x75\xd1\xb7\xc5\xee\x44\xb1\x5b\xe4\xbc\x71\xb1\x6f\xbf\x61\x2a\x62\xdf\x35\xec\x83\x81\xa1\x02\xd0\x44\x75\x8d\x49\xca\x06\x80\xa2\xa8\x8f\x7b\x74\x14\x60\x07\xe6\x62\xa8\x99\xeb\x06\xdd\x9e\xdb\xae\x57\x74\xb3\x6e\x7f\xff\x6f\xbb\x8a\x31\x9a\x0a\x60\xac\x2a\x80\x80\x1d\xc6\x20\x8e\x23\xec\xe0\x4b\xf9\x97\x59\x62\x3c\x88\x39\x40\xf6\x7c\xcd\x9e\x61\x80\x8b\x0c\x24\x47\x86\x57\x3b\xb0\xb7\x05\xd2\x5b\x6a\xd0\x5e\x8b\xa0\x7e\xcc\xc2\x58\xe2\x2f\x54\x30\x81\x99\x28\x4e\x15\x0d\x03\xd0\x01\xec\x7b\x16\xb5\x1f\xe4\x33\xc5\xc0\x27\xc9\x26\x28\x6b\xcf\x31\xe2\x5e\x89\xe3\xf9\x7c\x5e\x47\x16\x2d\x1f\x5b\x08\x7c\x08\x09\x45\xfe\x44\xe0\x7c\xa4\x54\x32\x2d\xbd\xef\x70\x8e\x78\xa2\x98\xda\x07\xa9\x17\x4c\xb3\x74\x1d\x5a\x8a\x24\x9a\xbc\x72\x5a\xb9\xa9\x45\xee\xb5\x1c\x79\xf9\x62\xe5\xed\xbc\xc8\x89\x53\xe9\x55\xce\xe7\x86\xdb\x0e\x95\x72\x39\xf6\x96\x59\x3f\xcd\xfa\xb8\xe4\x32\xce\x54\xb0\x11\xfc\x0b\x07\x5f\x63\xd4\xf3\xf7\x0e\xcc\xc1\x37\x82\xe6\x57\x4a\x4f\x02\x6a\x0a\x2a\x72\xdc\x20\x35\x69\x89\x59\x37\x18\xc1\x06\xa6\x44\x77\x5e\xd1\x7d\xc7\xfb\x29\xd3\x7c\x41\x9e\x65\x62\x14\x03\xfe\x00\xfb\xba\xa7\x26\xd8\x1b\x8e\x07\xaa\x73\x9e\x8d\x13\x39\xba\x95\x6d\x79\x98\xe6\xf9\x06\xc9\x18\x66\x90\xe7\x47\x3b\x67\x61\xed\x95\xb5\x57\xc1\xab\xbf\x7a\x74\x99\xe7\x76\x04\x14\xa5\xe8\x0b\xf6\x73\x0b\x57\x77\xfc\xcd\x4a\xcd\x8e\xf2\x96\xc6\xea\x29\x67\xc3\xff\x3b\x14\x2e\x27\xf9\x51\xc5\xf6\x9f\xcd\xc2\xcb\x32\x73\xc6\xaa\x5a\xda\xf9\x33\x00\x00\xff\xff\xbf\xe6\x1b\x66\x9c\x05\x00\x00")

// FileRootModuleTfTmpl is "root_module.tf.tmpl"
var FileRootModuleTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xca\x3b\x0a\x43\x21\x10\x05\xd0\x7e\x56\x71\x11\xcb\x44\xfb\x80\x4d\xd6\x90\x0d\x48\x1c\x82\xe0\x07\xfc\x54\xc3\xec\x3d\xf8\xaa\x57\x1e\x38\x22\x18\xb1\xfd\x18\x76\xe5\xf4\x80\x5d\x1c\x2b\x5e\x01\xd6\xbd\x77\x2e\xc9\x7d\x38\xd6\x09\x55\xaa\x3d\xed\xc2\x30\x27\x3c\x45\xae\x0f\x55\x03\x21\x60\xf6\x3d\xbe\x8c\x00\xe3\xfc\x09\xd3\xdf\x06\x29\x91\x08\xb8\x1d\xfd\x03\x00\x00\xff\xff\x66\x90\xa0\x30\x70\x00\x00\x00")

// FileScriptTfTmpl is "script.tf.tmpl"
var FileScriptTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x57\x4b\x8f\xdb\x36\x10\xbe\xeb\x57\x0c\x88\x14\x48\x50\x5b\xdb\xa0\x68\x0f\x01\xf6\x90\x5d\xf4\x05\x04\xc5\xa2\x3e\xf4\x50\x07\x04\x23\x8d\x65\xa2\x14\xa9\x72\xa8\xd8\x0b\x45\xff\xbd\x18\xca\x92\xe5\x67\xd2\xee\x23\x41\xb0\x3c\x18\xa6\x87\xc3\x99\xf9\xf8\xf1\xe3\xb8\x69\x20\xc7\x85\xb6\x08\x82\x32\xaf\xab\x20\xa0\x6d\x13\x8f\xe4\x6a\x9f\x21\x08\x5b\x1b\x23\xfb\xa9\x00\x51\x79\xf7\x5e\x93\x76\x56\x36\x0d\xa4\xbf\x60\x00\xd1\x5b\xa5\x55\x25\xb2\xbb\xa4\x80\xd5\x60\xe6\x89\xb4\x75\xf9\x0e\x7d\x34\xd6\x95\x71\x2a\x17\xd0\x24\x00\xc1\xeb\xa2\x40\x4f\x71\x02\xa0\x2d\x05\x65\x33\x94\x3a\x87\x4b\x10\xcf\x9a\xc2\xb9\xc2\xa0\xcc\x5c\x59\xd5\x01\x65\x6f\x4f\x4f\xc7\x4e\x47\x7b\xb4\x22\x01\x68\x93\x04\x20\xc7\x0a\x6d\x4e\xd2\x59\xb8\x84\xbf\x62\x2c\x31\xec\xb1\x35\xf2\x06\xec\xf3\x96\x7d\x86\x4a\xd1\x83\x58\x68\x83\x62\x93\x65\xd3\x80\x5e\x40\xfa\xab\xa3\x90\xfe\x46\x7f\x6a\x9b\xbb\x15\x31\x6a\x10\x47\xe6\xac\xc5\x2c\x68\x67\x37\xeb\x79\x2c\x1d\x85\xf8\xe5\x58\x59\x2a\xcf\x3d\x12\x9d\xab\x6a\xb3\x24\x66\xd7\x8d\x70\x5b\x61\xbf\xe3\x4a\x5b\x5f\x6e\x4d\x35\xa1\xef\x4d\xaf\xf3\x52\x5b\x4d\xc1\xab\xe0\xfc\xc8\x5b\x97\xe8\xea\x10\x97\xfc\xf8\xdd\xc8\xb7\x52\x44\x2b\xe7\x23\xfe\x9c\xcf\xb5\x2b\x2b\x0c\x9a\xcb\x49\xff\x70\x2e\xdc\xf4\xf6\x76\xc8\x25\x22\xcc\x63\xc3\x99\x3e\x34\xbb\x47\x90\xf8\x83\x2b\x81\xb6\xbd\x50\x44\x18\xe8\x82\x6d\xb3\xc8\xb7\xf4\x4a\x11\x8e\x76\xeb\xd0\xc5\x7f\x06\xfb\x1b\x65\x8b\x5a\x15\x4c\xd0\x25\x1a\x23\xb6\x48\xf3\xb9\x52\xd0\x56\x45\xb0\xbb\x88\x95\xd7\x36\x2c\x40\x5c\xbf\x9a\xcf\xe7\x73\xa3\x16\xce\x17\x28\x3b\x6a\xcb\x6f\x28\x7d\xa7\x82\x80\xe7\x87\xcc\x7c\xb1\x9b\x02\x9a\x98\xd4\x1d\x02\x55\xf4\xf2\x93\x02\xd9\xbc\x8f\x73\x10\xb6\x69\xe0\x59\x4e\x81\xb9\x07\xaf\x2e\x87\x88\x17\xde\xb9\x70\x71\x18\x91\x96\x27\x03\x9e\xa3\xa6\x2a\xd0\x06\x18\x4e\x6d\xa1\x0c\xa1\x38\x24\xee\x83\x70\xb7\xdb\x94\x68\x79\x84\xbc\x9d\x8d\x6b\x3d\x42\x5b\x38\xc2\x5c\xaf\xdf\xab\x80\xf2\x6f\xbc\xed\x32\x65\xdc\x9e\xf3\x51\x69\x9b\xe3\x1a\xd2\xab\x5a\x9b\x3c\xbd\x76\x76\xa1\x0b\xce\xd3\x48\xa2\xa5\x1c\xb9\xc9\xee\x96\xb7\xad\x78\x31\x22\xf7\x3d\x72\xfb\x90\x45\xc3\xf9\xf6\x8b\xc6\x94\x68\x93\x36\x79\x50\x29\xc6\x35\x66\x9f\x5f\x88\x77\xea\x4a\xef\xed\x81\x99\x1c\x15\x72\x8f\xa5\x0b\x38\xdd\x56\xfe\xa4\xe7\x27\xf5\x5c\x5b\xc3\x8d\x41\x7f\x50\xff\x59\x9d\xef\xac\xc7\x93\x71\xdc\x7d\x49\x16\x95\x5b\xa1\x8f\x51\x61\xfa\xbb\xbb\xf1\x2e\x5e\xa4\xe9\x4f\x6b\xcc\x6a\xae\xef\xc6\x19\x9d\xdd\xc2\xd5\x2d\x63\x00\xd3\x9f\xd9\x3a\xbf\xab\x74\xcf\xc5\x5e\x56\x5b\xfd\x66\xba\x3d\xe9\xf8\x17\xab\xe3\xa7\x49\x2d\xb2\x65\xe9\x72\xf8\x76\x0d\x7b\x8a\x3c\xd9\xa1\xf2\xc8\x32\x92\xf8\xd7\xbe\x98\x05\xaf\x6d\x31\x76\x78\xfb\xf8\x5a\xbe\x52\x3a\x7c\x7d\x5a\x1e\x75\x7a\x50\x72\x3e\x04\xa3\x6d\xbd\x96\x64\x10\x2b\xc9\x34\xf2\x7c\xa5\xbe\xef\x40\x66\xfb\xaa\xd3\xef\xfd\x15\x2f\x7f\xe0\x25\xc9\xa0\xf6\x6f\x5c\xa6\xcc\xa1\xdc\xef\xbc\x15\x86\xd7\xec\x3c\x15\x1f\x7d\x2c\xf8\x6e\x96\xa5\xb2\x11\xea\x59\x50\x3e\x4c\x67\x9c\xc8\xc9\xd4\xce\x76\x9d\x67\xf6\x3a\x84\xe1\x54\x5b\xc9\x0c\x08\xe8\x2b\x8f\x01\x3d\x9f\x94\xb8\x61\xe1\x9c\x45\xb9\x9e\x80\x98\x5e\x77\x51\x44\x47\xda\x0d\x8e\xe3\x54\xee\x15\x14\xba\x17\x38\xe8\x7f\x00\xd1\x57\xd6\xfd\xf2\xb0\xd7\xf1\x71\x6e\xe2\xe4\x31\xaf\x62\x14\x98\xa7\xa6\xea\x73\x37\x55\x00\x22\x47\x03\x0f\xda\x5a\x7d\x7a\x8c\xb3\xff\x72\x9f\x1a\xa5\xaf\xad\x51\xf2\x25\x4c\x17\x7b\x5d\x12\x7c\xf8\x00\xc1\xd7\xf8\x91\xe6\x67\x3b\x4f\xfe\x0d\x00\x00\xff\xff\x3e\x2f\xbd\x02\x74\x14\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileCommandTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "command.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileDNSRecordTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "dns_record.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileInfraTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "infra.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileProvisionedHostTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "provisioned_host.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRemoteFileTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "remote_file.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRootModuleTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "root_module.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileScriptTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "script.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
