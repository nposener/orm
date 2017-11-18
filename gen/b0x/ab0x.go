// Code generated by fileb0x at "2017-11-18 23:21:44.055738538 +0200 IST m=+0.008467142" from config file "b0x.yml" DO NOT EDIT.

package b0x


import (
  "bytes"
  "compress/gzip"
  "io"
  "log"
  "net/http"
  "os"
  "path"

  "golang.org/x/net/webdav"
  "golang.org/x/net/context"


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
type HTTPFS struct {}



// FileCommonGo is "common.go"
var FileCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x34\x8d\xb1\x8e\x83\x30\x10\x44\x6b\xef\x57\x8c\x5c\x81\xc4\x1d\xa7\x2b\x91\xac\x7c\x41\x9a\xb4\x51\x0a\x0b\x2d\xc8\x02\x16\xc7\x5e\xd2\x44\xfc\x7b\x04\x84\x6e\x57\xf3\xde\x4c\xf4\xed\xe0\x7b\x86\xc6\x31\x13\x85\x29\xce\x49\x61\xb3\xa6\x20\x7d\xb6\x44\xdd\x22\x2d\x9e\x57\x9f\x86\x5c\x08\x82\x68\x89\x23\xc4\x9b\x4c\xe8\x20\x70\x0e\x7f\xdb\x63\x12\xeb\x92\x04\xd6\x92\x59\xc9\xec\x0e\x1a\xf7\xc5\xf3\xef\x8d\x23\x7b\x2d\xec\xa5\x82\xad\x20\xe5\x89\xb8\xa3\xfe\xde\x8c\x2c\xc5\x7e\x96\x3f\xff\x0f\xd4\x35\x12\x4f\xf3\x8b\x31\xfa\xac\x9b\x62\xe9\x9c\xd8\x29\x5a\xe9\x13\x00\x00\xff\xff\x56\x56\x3e\xf3\xbd\x00\x00\x00")

// FileCreateGoTpl is "create.go.tpl"
var FileCreateGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\x90\x41\x6b\x32\x31\x10\x86\xcf\x9b\x5f\xf1\x1e\x16\xbe\x5d\xd0\x95\xef\x5a\xe8\xc1\xca\x16\x8a\xb6\x58\xdd\x8b\x37\xa7\xeb\xac\x86\x66\x47\x49\xb2\x2d\x12\xf2\xdf\x4b\xd4\x52\x6f\x93\x37\xcf\x43\xde\xcc\x89\xda\x4f\xda\x33\x42\xa8\x96\xd7\x31\x46\xa5\x42\xc8\x05\x0f\x8f\x30\x2c\xa8\x9a\xf3\x89\xab\x67\xcd\x66\xe7\xd2\x5d\x37\x48\x8b\xa2\x45\x33\xb3\x4c\x9e\x4b\xac\xbd\xd5\xb2\x2f\x4a\xb8\xcb\x80\xa0\x00\x60\x32\xc1\x15\x80\xf3\xe4\xb9\x67\xf1\x38\x90\x03\xc1\x68\x61\x74\x47\x0b\xa6\xf6\x80\x2f\xb2\x9a\x3e\x0c\xe3\x5b\xfb\x03\xb4\xff\xe7\x20\xd4\x33\x48\x76\xd7\x93\x4f\xcf\xab\xcc\xb2\x1f\xac\x60\x3b\x5b\xd5\xd3\xa6\x46\x33\x7d\x5a\xd4\xa9\x75\x93\xe4\x18\x51\xa8\x2c\x84\x31\x2c\xc9\x9e\x91\xeb\x11\xf2\x2e\xfd\xe0\xbe\x3d\x62\x54\x59\x16\x42\xde\x55\xb3\xa3\x19\x7a\x79\xa3\x3e\xa9\x97\x64\xfd\xbe\xb8\xb0\x37\x66\x0c\xdd\xe1\x16\x2f\xad\xee\xc9\x9e\xe7\x7c\x46\x8c\x58\xae\x5e\x5e\xa7\xab\x0d\xe6\xf5\x26\x04\xb0\xec\x70\xaf\x18\x8f\xe2\x64\x06\xf7\x1f\xb9\x2e\x91\x0b\x62\x1c\xfd\x61\x69\x31\x09\xfc\xb5\xca\xad\x8a\xea\x27\x00\x00\xff\xff\xa0\x0e\x64\x2c\x83\x01\x00\x00")

// FileCreateCommonGo is "create_common.go"
var FileCreateCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x64\xd0\xc1\x6a\xc3\x30\x0c\x06\xe0\x73\xf4\x14\xc2\x30\x6a\x8f\x92\xde\x0b\x3b\x6c\x5d\xef\x63\xeb\x7d\x38\x8e\x92\x86\x39\x76\x2a\xab\x65\xa3\xf4\xdd\x87\x53\xb7\x0c\x76\x0a\x28\xfa\xbf\x5f\xc9\x64\xdd\x97\xed\x09\x65\xf2\x09\x60\x18\xa7\xc8\x82\x1a\x2a\xd5\x5a\xb1\x8d\x4d\xb4\x4a\x07\xaf\xa0\x52\xdd\x28\xf9\xe1\x63\xaf\xc0\x00\xac\x56\xb8\xdb\x30\x59\x21\x1c\x12\x5a\x4c\xc2\x47\x27\x28\x7b\x2b\xb8\x8f\xbe\x4d\x98\x01\xec\x22\xa3\xec\x09\x37\xef\xdb\xe7\xdd\x16\x93\x58\xa1\x91\x82\x80\xfc\x4c\x74\x17\x4a\xf8\x0c\x55\x37\x4a\xfd\x21\x3c\x84\x9e\x18\x2e\x73\x4d\xd9\x61\x92\x23\x87\x3f\x55\x99\xb6\xff\xe1\xee\x18\x5c\xc9\x68\x73\x6f\x38\x43\x75\x05\x6e\x93\xf3\xa5\xf0\xdb\x6f\x72\xe8\xe6\x59\xc6\xc5\x36\x9e\xee\x67\xf7\xc3\x89\x42\x29\xbc\xca\xda\xdd\x04\x33\x47\x75\xdb\xe0\x63\x3a\xf8\xfa\xf5\xc5\x20\x31\x47\xce\x5d\x49\x46\xc1\xf5\x13\xba\xf2\x31\xda\x40\xe5\x63\x5f\xbf\xf1\x10\xa4\xd3\xea\x2a\xac\x71\xf1\x70\x5a\xa8\x25\xe6\x75\x03\xd5\xe7\x32\x0b\x39\xd7\x36\xf5\x8c\x97\x17\xe5\x74\xe2\xfc\x4f\x7e\x03\x00\x00\xff\xff\x30\xd5\xb8\xb9\xb2\x01\x00\x00")

// FileInsertGoTpl is "insert.go.tpl"
var FileInsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x90\xd1\x4a\xc3\x30\x14\x86\xaf\x9b\xa7\x38\x84\x09\x89\xd4\x3c\x80\xb0\x0b\x91\x09\x15\x9d\xc3\x56\x6f\x5d\x5c\xd3\x99\x2d\x4d\x6b\x92\x0e\x24\xe4\xdd\x25\x4d\xeb\x85\x4c\xbc\x4b\xfe\x73\xfe\xff\x7c\xe7\xf4\x7c\x77\xe4\x7b\x01\xde\xb3\x4d\x7a\x86\x80\x90\x6c\xfb\xce\x38\x20\x08\x00\x00\x37\xad\xc3\xe9\x65\x9d\x91\x7a\x6f\x31\x4a\x5f\xef\x59\xf5\xd5\x0b\x56\x8c\xed\x1b\xee\x3e\x42\xc0\x88\x22\xd4\x0c\x7a\x07\x44\x42\x55\x68\x2b\x8c\xa3\x50\x8e\x46\x42\x21\x25\x80\x47\x99\x11\x6e\x30\x1a\x9a\xd6\xb1\xb2\x37\x52\xbb\x86\x6c\x8b\x75\xb9\x7a\xae\xa0\x58\x57\x4f\x91\xa8\xe2\xef\x4a\x84\x00\xe4\xc2\x52\x78\xbd\x79\x78\x59\x95\xe3\x7b\x9b\xa3\x2c\x9b\x58\xd8\x7d\x27\x35\x91\x6c\xd7\x29\x9b\x03\xce\x01\xd3\x58\xfd\x7c\xe4\xe6\x68\x89\x12\xb1\x76\xe2\x6a\x10\x96\xc6\x02\x45\xe1\x0c\xde\xbc\xc8\x9a\xb7\x22\x04\xd2\xc3\xe5\xac\xdc\x0d\x4a\x25\x95\xce\xed\x11\xfe\xc4\x0d\x1c\x60\x09\x12\x65\xde\x5f\x81\xe1\x7a\x2f\x60\xf1\x96\xc3\xa2\x81\xeb\x25\x4c\x5e\x29\x54\x6d\x43\x40\x59\x6c\x3d\x30\x5e\xd7\x04\x7b\xbf\x68\xd8\x6d\xa7\x86\x56\xa7\x5c\x9c\x43\xcf\x46\x75\x9a\x93\x22\x85\xae\xa3\x73\xba\xd2\x21\x62\x7b\xff\xcf\x9c\x33\x7b\xfd\xc4\x92\xf1\x08\x49\x89\xae\x5f\x0b\x4d\x73\xe4\xdf\x90\xa3\x3f\x9e\xcf\xfb\xc4\xf6\x1d\x00\x00\xff\xff\xb8\x88\xa2\x91\x3b\x02\x00\x00")

// FileInsertCommonGo is "insert_common.go"
var FileInsertCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x64\x51\xc1\x6a\xdc\x30\x14\x3c\xeb\x7d\xc5\x60\x08\x91\x8a\xf1\xf6\x1c\xd8\x4b\xe9\x1e\x72\x29\xa5\xc9\x2d\x84\xa2\xb5\x9f\xbd\xa2\xb6\xe4\x48\x6f\xb7\x2d\x9b\xfd\xf7\x22\xdb\xeb\x26\xe4\x64\x33\x7a\x33\x6f\x66\xde\x68\xeb\x5f\xb6\x63\xc8\xd8\x27\x22\x37\x8c\x21\x0a\x34\xa9\xa2\xb1\x62\xf7\x36\xf1\x26\xbd\xf4\x05\xa9\xa2\x1d\x24\x7f\xfa\xd0\x15\x64\x88\x36\x1b\x3c\xde\xfb\xc4\x51\xe0\x12\x2c\x92\xc4\x63\x2d\x90\x80\x43\xe8\x1b\x38\xdf\x86\x38\x58\x71\xc1\xa3\x0d\x11\xd6\xe3\xfe\xdb\xc3\xee\xc7\x23\x92\x58\xe1\x81\xbd\x90\xfc\x1d\x79\x15\x59\xf8\x67\x52\xed\x20\xd5\x83\x44\xe7\x3b\x8e\xa4\xea\xd0\x27\x00\x4f\xcf\x69\x82\x48\x9d\x6c\x7f\xe4\x84\xa7\x67\xe7\x85\x63\x6b\x6b\x3e\x5f\xe8\x32\x19\x5a\xa4\x22\xcb\x31\xfa\x6c\xca\xf3\xef\x8f\x6b\xdb\xa3\xaf\x97\x51\x6d\xd6\xfd\x67\x52\x33\xef\x8a\xac\xaa\xbb\x3f\x5c\xc3\x4d\x58\x82\x1c\x18\xb9\x99\x1c\x34\xff\x77\xee\xc4\x1e\xd7\xae\x66\x6d\xed\xae\x1a\x66\x22\xeb\x66\x8f\x4f\xe9\xa5\xaf\xbe\x7e\x31\xe0\x18\x43\xcc\xdb\x5c\x8b\x9e\xbd\x76\x55\x4e\x68\xb0\xdd\xe2\x33\x5e\x5f\x17\x6c\x0e\xb9\xa0\x67\x52\x57\x6f\xb9\x9b\x5d\x56\x68\x75\xe1\x83\x1c\x9c\xef\xb2\x93\xd9\x5d\x61\x48\x5d\x88\x54\x92\x41\x70\xb7\x85\x5b\x6a\xd4\x86\x54\x1f\xba\xea\x7b\x74\x5e\x5a\x5d\xcc\xde\xee\x70\x7b\x73\xba\x85\xbe\x39\x99\xa2\x44\xe6\x94\x58\x17\x93\xfa\x59\x66\xab\x59\xa6\xd9\x57\x53\x8a\xf7\x23\x55\x55\x99\xb5\x32\x8e\x31\x97\xf5\x21\xbd\x6d\x1a\xed\xed\xc0\x98\x8f\x57\x62\xe2\xe2\xcd\xe5\xde\x1d\x60\xee\x02\x5b\xd8\x71\x64\xdf\x2c\xdd\x94\xc8\x12\x26\x3f\x2f\xb7\x7f\x33\x30\x23\x8b\xf0\x7f\x43\x8e\x2e\xf4\x2f\x00\x00\xff\xff\xc5\xd1\x28\x8e\xda\x02\x00\x00")

// FileOpGo is "op.go"
var FileOpGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x44\xcb\x41\xaa\x83\x40\x0c\xc6\xf1\xf5\xcb\x29\x82\xab\xd7\x95\x17\xd0\x2e\x84\x50\xa4\xa2\x94\x16\xba\x1e\x24\xc8\x60\x89\xa9\x13\x28\xbd\x7d\xd1\x91\x99\x2c\x7f\xf9\x7f\xea\xc6\xd9\x4d\x8c\xa6\xaf\x00\x50\x96\x38\x28\xfa\x80\x4e\xf0\x7e\xeb\x70\x51\x5e\x9d\xf9\x45\xc0\xbe\xca\xdb\x2f\xd8\xea\x65\x02\x18\x17\x09\x86\xff\xf0\x37\x28\xbd\x71\xbf\x1a\x8b\xba\xd8\xa0\xe7\x04\xd5\x79\x97\x8b\x25\x39\x80\x32\xc4\x51\x97\x93\x2a\x42\x4e\xaa\x98\x34\x6c\x1f\x66\xd9\xa4\xa1\xc7\x93\xa8\x8f\xa1\x9f\xf9\x08\xbb\xf6\x4a\xbb\xb5\x92\xc6\x6d\x5f\xc0\x09\x7e\x01\x00\x00\xff\xff\x18\x53\xc5\x2f\xe8\x00\x00\x00")

// FileQueryGoTpl is "query.go.tpl"
var FileQueryGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x92\xc1\x6e\xdb\x3c\x10\x84\xcf\xe4\x53\xcc\x4f\x20\x7f\xa8\xc4\xa0\xef\x29\x7c\x68\x53\x07\x28\x90\xa6\x49\xed\x5b\xd1\x03\x2d\xaf\x65\xa1\x34\x65\x2d\x29\xc7\x81\xa1\x77\x2f\x48\xc9\x09\xd0\xe4\x26\x2c\x67\x96\xdf\x0c\xb5\xb7\xe5\x1f\x5b\x11\x4e\x27\xf3\x38\x7c\xf6\xbd\x94\xf5\x6e\xdf\x70\x84\x96\x42\xad\x6d\xb4\x2b\x1b\x68\x1a\x5a\xa7\xa4\x50\xae\xa9\x94\x94\x42\x9d\x4e\x66\xf9\xb2\x27\xf3\x2d\x4b\x1f\x6d\xdc\xf6\xbd\x92\x85\x94\xd3\x29\x16\x91\x6b\x5f\x81\x29\x76\xec\x03\xe2\x96\xb0\x78\xba\x47\xdb\x11\xbf\x20\xe4\x43\xb9\xe9\x7c\x09\xdd\xe2\x6a\xf9\x94\xc6\xc5\x68\xd2\xc5\x28\xc0\x49\x8a\x61\x01\xd4\x62\x7e\x3f\xbf\x5d\x42\xe1\x1a\xad\x09\xe4\xcc\xab\xf6\x1a\x0a\x77\x3f\x7f\x7c\x4f\xfc\x4b\xbb\x72\xd4\xf7\xa3\xec\x79\x4b\x4c\xaf\x42\xd9\x67\xb0\xf9\x91\x4a\x70\x37\x32\xe5\x8b\xd1\x78\x58\x54\xf5\x81\x3c\xce\x59\xcd\x7b\xba\xe4\xd4\xeb\x15\xae\x42\xeb\xcc\xd7\x2f\x05\xf4\xaf\xdf\xe7\x0a\xee\x3a\xe7\x1e\xec\x8e\xfa\x7e\x02\x62\x6e\xb8\x48\xf0\xd3\x29\x4a\x26\x1b\x09\x81\x1c\x95\x11\x21\xda\x48\x3b\xf2\x51\x8a\x10\x77\x11\x37\x33\xb4\x6f\x80\xc2\x72\x15\x86\xd9\x80\xfe\x99\xab\x90\xe6\xae\xa9\xcc\x23\xd7\x3e\x6e\xb4\xca\x30\x37\xb8\xbc\x38\x5c\xe2\xe2\xa0\x26\x48\x8b\x26\x48\xd6\x42\x0a\x6e\x9e\x43\x26\x48\x6b\xd6\x2b\x93\xd5\xfa\x4d\x62\x8c\x29\xa4\xa8\x37\x59\xf2\xdf\x0c\xbe\x76\x09\xf4\x5c\xb3\xaf\x5d\x76\x4b\xd1\x4b\xb1\xa6\x0d\x31\xd2\x46\x73\xeb\x9a\x40\xba\x90\x39\x12\x1d\x23\xdb\x32\xe6\x13\xc4\x26\x3d\x56\x57\xc6\x8e\x29\x48\x71\xb0\x0c\xeb\x1c\x3e\x6a\x46\x8a\x4d\x33\xee\x7b\xa0\x63\xd4\xb9\xa2\xec\xa8\xf1\x91\xfa\x8c\x79\x33\x1b\x4c\x8b\xd2\x7a\x3d\x3c\x7e\x28\xad\xcf\xe5\xfc\x5f\x17\x29\xd2\xa7\x7f\xf3\xbc\x0f\x94\x12\x89\x84\x36\x83\xdd\xef\xc9\xaf\xb5\x75\x6e\x82\xba\xc8\x59\x47\x79\x1e\xe5\xcb\xe6\xcc\xf9\x97\xf9\x1b\x00\x00\xff\xff\xe5\x6c\x31\x9a\x1d\x03\x00\x00")

// FileQueryCommonGo is "query_common.go"
var FileQueryCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x90\x31\x6e\xc3\x30\x0c\x45\x67\xeb\x14\x7f\x74\x96\xe6\x04\x39\x44\xd1\x00\x9d\x05\x85\xa9\x85\x2a\x92\x2c\x52\x30\x02\x23\x77\x2f\x2c\x2a\x41\xbd\x64\xb1\x01\x92\xff\x3d\x52\xd9\xba\x5f\xfb\x43\x90\x1c\xd8\x98\xe3\x11\x9f\x95\xca\x1d\x85\xa4\x96\xc8\xb0\x60\x29\xd5\x09\x24\x21\x53\xb9\xa6\x72\xc3\xdc\x26\x52\xa6\x62\xc5\xa7\xc8\xe6\x5a\xa3\xd3\xdc\x78\xc0\x59\x01\xab\x19\x94\xd1\x0b\xeb\xc3\x3c\x1a\xbf\xf7\xfd\x9e\x3d\xa5\x70\xe9\x60\x1f\x37\x4b\x43\x1b\xb9\x67\x7a\x26\xfa\xf0\x6a\x06\xa6\x00\xe0\xfc\x45\x81\x9c\x98\x61\x99\xa8\x10\xbe\xb7\x6f\x77\x68\x07\x36\xe7\xe0\x89\xe1\x52\xa8\xb7\xc8\x60\x2d\xa7\x08\x99\x48\x6d\xba\xfb\x38\x77\xc9\xa1\x47\x47\x7e\xe2\xff\x1f\x34\x7f\x6c\xe6\x13\xf8\x75\xdb\xdc\x85\x4d\xfe\xf2\xe9\x42\x2e\xc5\x8b\x6f\x0f\xf4\xde\xd8\xb2\xe3\xa2\xff\xbd\x4e\x41\x27\x2c\x3b\xe1\x5f\x00\x00\x00\xff\xff\x46\xdf\xa5\x07\xb3\x01\x00\x00")

// FileSelectGoTpl is "select.go.tpl"
var FileSelectGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x92\x4f\x6b\xdc\x30\x10\xc5\xcf\xd6\xa7\x78\x0d\x4b\xb1\x5b\xc7\xdb\x5c\x5b\xf6\x10\x0a\x81\x5e\x4a\x20\xb9\x2d\x4b\x51\xed\xb1\x57\x44\x7f\x1c\x49\xdb\xb2\x08\x7d\xf7\x22\xd9\x4e\xbc\x6d\x0f\xd1\xc9\x32\x33\x4f\xbf\xf7\x66\x46\xde\x3e\xf1\x81\x10\x42\x73\x3f\x7d\xc6\xc8\x98\x50\xa3\xb1\x1e\x57\x21\x34\x8f\xe7\x91\x9a\x6f\xf9\x7e\xcf\xfd\x31\xc6\x2b\xc6\x42\x80\xe5\x7a\x20\x6c\x7e\xd4\xd8\xf4\xf8\xbc\xc3\x54\x77\x27\x48\x76\x0e\xd7\x31\xb2\xed\x16\x21\x6c\xfa\xe6\x3b\x57\x14\x23\x6e\xbb\xee\xe2\xee\x0d\xfc\x91\xe0\x48\x52\xeb\xa9\x43\x6b\xe4\x49\x69\x98\x1e\x1c\xcf\x27\xb2\x67\xd6\x9f\x74\x8b\xd2\xe1\xf1\x21\xd7\x54\xeb\xf6\xb2\x5a\x7e\x23\x30\x00\xb0\xe4\x4f\x56\x83\x8f\x23\xe9\xae\x74\x75\x42\xdf\xf4\xcd\xd7\xac\x3a\xf5\x5c\x55\x2c\x26\x72\xd2\x5d\x06\x4c\x84\xae\xe5\xfa\xd6\x0e\x0e\xdc\x12\xa4\x70\x3e\x01\xf4\x93\x09\x6f\xf0\x93\x30\x88\x5f\xa4\x5f\x68\x9f\x25\x1e\x5a\xae\xd1\x1a\xa5\xb8\xee\xfe\x65\x5c\x04\xcb\x11\x1f\x96\xf0\xee\x4e\x52\x4e\x08\x15\xf6\x07\xa1\x3d\xd9\x9e\xb7\x14\x22\x02\x2b\x44\x0f\x49\xba\x74\x15\x76\x3b\x7c\x9a\xdd\xa4\xb3\xdd\x82\x77\x5d\x7a\x99\x67\x40\x29\x33\xc2\x0c\x67\x7a\x8c\x2f\xa5\xb3\xf9\x0b\xed\x57\xa1\x74\xde\x32\xb0\x75\xfd\xfb\xb1\x59\xa5\x5d\xff\xad\x95\x22\x5c\x75\x44\x56\x44\xc6\x66\xe8\x69\xa2\xf8\xcd\xdd\x94\x5d\x8d\xf6\x68\x8c\x23\x18\x2d\xcf\x70\x46\x2d\x1e\x58\xa1\x12\x87\xe2\x4f\x54\x2a\x3e\xee\x9d\xb7\x42\x0f\xc9\x43\x3d\x47\x52\x65\xcd\xde\x58\x88\x3a\x2d\x48\x2a\x9f\x6c\xb8\x55\x50\x6a\xdf\x1a\x79\xc0\x0e\x02\x1f\x71\xc3\x66\xa0\x1c\xda\x22\x7f\x91\xcc\xab\x7a\xf1\x96\x58\xd2\x84\x44\x56\xda\xff\x67\xa7\x0e\x5f\x20\xf0\x6e\x9a\x5c\x91\x1f\xdd\x8b\xeb\x9b\x44\x73\x19\x61\x4a\xa8\x58\x2d\x5f\xb1\x2c\xac\x1d\x1c\x8b\xec\x4f\x00\x00\x00\xff\xff\x60\xf3\xd3\x38\x86\x03\x00\x00")

// FileSelectCommonGo is "select_common.go"
var FileSelectCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x90\x3d\x6e\xeb\x30\x10\x84\x6b\xf3\x14\x03\x56\xd4\x83\xf1\x9c\x0b\xb8\x0a\xdc\x04\xe9\xe4\x2e\x48\xb1\x91\xd6\x96\x12\x89\x24\xb8\xab\x22\x30\x74\xf7\x80\xa6\xe8\x26\x69\x67\xe7\xe7\x23\x23\x75\x5f\x74\x65\x68\x9c\xc4\x98\x71\x8e\x21\x29\xac\x68\x1a\xfd\x55\xac\x31\x87\x03\x5a\x9e\xb8\x53\x24\xd6\x25\x79\x01\x79\x84\x8f\xcf\xac\x68\x40\x97\x98\x94\x41\x68\x4f\xaf\xa7\xe7\x33\x44\x49\x79\x66\xaf\xe6\xb2\xf8\x6e\x8b\xba\x06\xe7\xad\xe4\x66\x76\xa5\xa7\x2a\xb7\xd5\xac\xf7\x95\xea\x18\x05\x3a\x30\x44\xd3\x92\x27\x06\x52\x0c\x61\xea\x8b\xba\xad\xf4\xa4\x64\xf4\x3b\xf2\x23\xf5\xf6\x5e\x90\x4d\xd9\x75\x52\x2f\x0d\xda\xfb\xc1\x35\x28\x8e\x8c\x30\x5e\x30\xb1\x77\xd2\xe0\x78\xc4\x53\x56\x2a\x95\xfd\x67\xcd\x6e\x7d\x40\x6e\xff\xf0\xff\x25\x8c\xde\xc9\x1e\x76\x0f\xdb\x64\xe0\x5f\x2b\xd4\xf7\xae\x0b\xd3\x32\xd7\xd0\x5f\x6f\xa6\x18\xd9\xf7\xb9\xa8\x58\x73\xd5\x4f\x00\x00\x00\xff\xff\xfc\xff\x7f\x6a\x81\x01\x00\x00")

// FileWhereGoTpl is "where.go.tpl"
var FileWhereGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x91\xb1\x6a\xc3\x30\x10\x86\x67\xeb\x29\x8e\x90\xc1\x06\xa3\xec\x85\x2c\x29\x2e\xf5\xd0\xa4\x94\x40\x86\x10\x8a\xb0\xcf\xb6\x88\x7c\x32\xb2\x9c\x50\x84\xde\xbd\xd8\x4a\x4b\xa0\x09\x5d\xb2\x99\xf3\xdd\xa7\xff\xbe\xeb\x44\x71\x14\x35\x82\x73\xfc\x3d\x7c\x7a\xcf\x98\x73\x46\x50\x8d\x30\xff\x4c\x61\x5e\xc1\xd3\x12\xf8\xf6\xab\x43\xfe\x22\x51\x95\xbd\xf7\x6c\xb1\x80\x5d\x83\x06\x9d\x9b\x57\x7c\x2d\x5a\xf4\x1e\x44\x59\xf6\x20\xa0\xd0\x54\x4a\x2b\x35\x81\x26\xb8\xfe\x6f\x35\xd8\x06\x61\xf7\x9a\x7d\x64\xd0\x5b\x61\xb1\x45\xb2\xac\x1a\xa8\xf8\x03\x8b\x75\x07\x9b\x2e\x85\x93\x50\x81\x31\x3e\xef\x7d\x12\x1a\xc1\xb1\xc8\xa0\x1d\x0c\x01\xe1\x79\x2a\xc5\xba\x4b\x61\x36\xb5\x3e\x6b\x35\xb4\x14\x38\xb3\x09\x91\x30\xcf\x6e\x45\xce\xe9\x12\x9a\x20\x5f\x3f\x26\x78\x4e\xf1\x49\xa8\x1e\x38\xe7\xb7\x73\x0b\x53\xf7\xa3\xd0\x56\x1c\x31\xde\x1f\x24\x59\x34\x95\x28\xd0\xf9\x14\x14\x86\xe9\x24\x61\x51\xa5\x0d\xc8\xb1\x31\x5c\x62\x82\x3a\x16\x4d\xf3\x7b\x79\x80\xe5\x54\xda\xcb\x03\x8b\xfc\xb5\x8d\xb7\x41\x05\x21\x9b\x2e\xa7\x3b\x4a\x46\x06\xe7\xfc\x9e\x96\x15\xda\x33\xe2\x8f\x1b\x58\x65\xdb\x5d\x96\x3d\xc8\xcf\x85\x1d\x2b\x7d\x4e\xa1\x91\x75\xf3\xef\x79\xaf\x16\xba\x0c\xdf\xd9\xea\x17\x39\xee\xe5\x1c\x52\xe9\x3d\xfb\x0e\x00\x00\xff\xff\x31\x3c\xee\xa1\xe0\x02\x00\x00")

// FileWhereCommonGo is "where_common.go"
var FileWhereCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x55\xdf\x6f\xdb\x36\x10\x7e\x16\xff\x8a\x1b\x81\x20\xd2\x26\xa8\x43\x1f\x0b\x64\x43\xba\x0a\x5b\x87\xd4\xde\x9a\x02\x79\x28\xfa\xc0\x48\x27\x99\x08\x4d\x32\x24\x55\xcd\x48\xfd\xbf\x0f\xfc\x21\x5b\x76\x0d\x23\xe9\x8b\x2c\x8a\xc7\xbb\xef\xfb\xee\x3e\x5a\xb3\xe6\x81\xf5\x08\x4e\x0b\x4b\x08\x5f\x6b\x65\x1c\xe4\x24\xa3\xdd\xda\x51\x92\x51\xa1\x7a\xff\x63\x9d\xe1\xb2\xb7\x94\x14\x84\xbc\x7a\x05\x77\x2b\x34\x08\xcc\x20\x28\xed\xb8\x92\x16\x3a\x65\xe0\xf6\xdf\x1b\xb8\xfb\xab\xfe\x58\x83\x75\xcc\xe1\x1a\xa5\x23\x6e\xa3\x31\x85\x5b\x67\x86\xc6\xc1\x13\xc9\xac\x5b\x3b\xf8\xfc\x25\x26\x25\x19\x33\xbd\x85\xcf\x5f\xb8\x74\x68\x3a\xd6\xe0\xd3\x96\x6c\x43\x19\x89\x63\x3c\x6a\xd0\x0d\x46\x5a\x60\xfe\xd3\x77\x35\xba\x41\x36\xbb\xd8\x5c\x69\x58\xea\x12\xbe\x32\xc3\xd9\xbd\x08\x65\xb9\xec\xfd\x07\x31\x20\xcc\x8a\x14\x09\x97\x07\x34\x72\xd7\xac\x40\x69\xbf\x68\x98\x45\x58\xea\xfa\xb1\x84\xa5\x5e\xa0\x7f\xfe\xe9\xc2\xb3\xf6\xcf\x9b\xf0\x7e\x13\xdf\xf9\x03\xbe\x21\x59\x8b\x1d\x1b\x84\x7b\x43\xb2\x4c\xa8\xbe\xfa\x87\x49\xde\x74\x39\x5d\x6a\x34\xcc\xcb\x03\x17\x16\xb8\x05\xa9\x1c\xb4\xd8\x71\x89\x6d\xd0\x4b\x49\x8c\xa8\x68\x09\x4a\x17\x24\xdb\x92\xec\x2b\x33\x30\x46\x60\x24\x1b\xab\xa0\xd4\x15\x30\xad\x51\xb6\x79\x5c\x97\xd0\xad\x5d\x75\xab\x0d\x97\xae\xcb\xe9\x85\xf5\xe9\x7f\xa7\x7b\xca\x21\x5b\xe1\x8f\x07\x65\x67\xc7\xfd\x3a\x29\x51\x90\x2c\xaa\x0a\xe3\x5e\xed\x0f\x83\x78\x86\xe0\xbb\x66\xab\x89\xa0\x85\x91\xbb\x15\xac\x95\x41\x70\x2b\x26\x3d\x35\x9f\x32\x54\x2a\xc1\x0e\xcd\x0a\x98\x85\xcb\xf7\x8b\x4b\x60\xb2\x85\xcb\xb7\xf5\xa7\xbb\xba\x5e\x5c\x56\xbb\xe6\x4d\xa5\xcf\xf7\xcf\x42\x55\x55\xa7\x7b\x78\xa0\xdc\xa9\x86\xbe\x45\x37\x22\x4a\xdf\x25\xde\x81\x40\x99\xc7\x94\x05\xfc\x74\x05\xaf\x7d\xe0\x99\xf6\xb1\xa6\x41\xed\x2c\x28\x29\x36\xf0\x3a\x81\x29\xa1\x57\x0e\x2e\xda\xb4\x0c\x6d\x2c\xe7\x99\x0b\x92\xf9\xa6\xbe\xa8\x91\x70\xbd\x78\x77\xb2\x9d\x89\xc6\xfb\x53\x0c\x7e\x83\x5f\x23\x81\x17\x54\xca\x2f\x6c\x71\x54\xa6\x84\xc7\x0f\xcc\x3c\xd8\x7c\xce\x61\x62\xf1\x9c\x29\x6f\x15\xc6\x39\xb7\x83\x0e\x57\xc9\x7a\x10\x8e\x6b\x81\x73\x89\xe2\xa4\x9f\x9d\x4e\x5b\x55\xd5\xd1\x84\x86\x41\xc9\x47\xf8\x39\xb4\xb8\x80\xdb\x30\x17\x79\x91\x06\xc4\xd3\xe7\x1d\x8c\x70\x75\x05\x92\x0b\xf8\xf6\x2d\x08\x14\xf9\x17\xfe\x6b\x54\x28\xa5\xa4\x34\x80\x98\x56\x71\xc2\x29\xfc\x92\xb2\xd9\xea\x6f\xc5\xe5\x4e\x3d\x0a\xb4\x48\x2e\xb9\xf6\xa0\x59\x98\x73\x7f\x03\xf6\x83\x77\x44\xbc\xff\xf0\x3f\x6c\x06\xe7\xb1\x30\xb8\xad\x6f\xea\x3f\x3e\xc1\xe3\x80\x66\x13\xcd\xc1\x92\x8d\x1a\x25\x5b\xee\xf5\xfa\x8e\x91\x4f\x9d\x17\x87\x37\xe1\x31\xad\x19\x05\xc9\xc5\x9c\x43\xd4\x2f\xa1\x5c\x1a\x2f\xab\xe0\x68\xc1\xbb\xd1\xec\xab\xc2\x7d\xb4\x01\xb8\x51\xc1\x18\xcc\xb3\xdb\xb3\x3b\x48\x09\xd1\xd2\xe4\x86\xf7\x2b\x37\xad\x77\x66\x4b\x35\xef\xb9\x64\x66\x93\x8f\x25\x84\xb0\x12\xe8\xf2\xe3\x5e\x29\xd9\xce\x41\x78\xdf\xff\x18\x8a\x6b\xd9\xbe\x14\xc6\xf5\xe2\x5d\xc4\x11\x52\xa5\x7d\x11\x0f\x96\x60\xa6\x17\xa5\x53\xbf\x67\x39\xc5\x91\x83\xa6\xff\xa9\x27\x9a\xd3\x6d\x09\x71\x3b\x8e\xe7\x71\xa8\x98\xc6\xa5\x48\x57\x01\xcd\xe9\x99\x30\xf3\xbc\x54\x61\xf7\xd0\x2b\x22\x79\xc5\x84\xdf\x03\xaf\x08\xb2\x25\xff\x07\x00\x00\xff\xff\x8b\x59\x81\xea\xd5\x07\x00\x00")



func init() {
  if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}



var err error



  




  
  var f webdav.File
  

  
  
  var rb *bytes.Reader
  var r *gzip.Reader
  
  

  
  
  
  rb = bytes.NewReader(FileCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileCreateGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "create.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileCreateCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "create_common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileInsertGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "insert.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileInsertCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "insert_common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileOpGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "op.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileQueryGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "query.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileQueryCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "query_common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileSelectGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "select.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileSelectCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "select_common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileWhereGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "where.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileWhereCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "where_common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
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
	f.Close()
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


