
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



// File00ImportGoTmpl is "00_import.go.tmpl"
var File00ImportGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x8e\xc1\xca\xc2\x30\x0c\x80\xcf\xeb\x53\x84\xb2\xc3\xff\x83\xae\x77\xc1\x07\xf0\xee\x5d\xba\x36\xce\xc2\xd2\xcc\xb4\x8a\x52\xfa\xee\xb2\xb9\x93\xa0\xc7\x2f\xdf\x97\x90\x40\x13\x4b\x86\x3f\xd5\x68\xc7\x31\xe3\x23\x6b\xd5\x68\x6f\xb3\xed\x6d\x42\x93\xae\xe3\x27\x1b\x2f\xe1\x8e\x32\x8f\xcf\xb4\xd4\x43\xc8\x97\x5b\xdf\x39\x26\x33\x71\xc2\x88\x62\x58\xe8\xbb\x31\x8e\x89\x38\xfe\x08\x7c\xb0\x23\xba\xf9\x78\x29\x20\x36\x0e\x08\xed\x69\x03\xed\xfa\xed\x6e\x0f\xdd\xf1\x39\x61\x77\x58\x38\xc1\xb6\x56\xd5\xe8\x52\xd6\xa0\xd6\xf7\x26\x46\xbf\xa8\x7f\xf5\x0a\x00\x00\xff\xff\xbb\xf7\x97\x25\xe7\x00\x00\x00")

// File10OrmGoTmpl is "10_orm.go.tmpl"
var File10OrmGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x59\x5f\x6f\xdb\x48\x0e\x7f\x8e\x3f\x05\x2f\x08\x02\xfb\xe0\xaa\xef\x06\xfc\x90\x26\x2e\x2e\xb8\x5c\x9a\x8b\x53\x74\x81\xa2\x58\x8c\x25\xca\x99\x5d\x79\xe4\xce\x8c\xf3\x07\x82\xbe\xfb\x62\xc8\x91\x2c\xc9\x92\xac\xa6\xc5\x3e\x6d\x5f\x6a\xcd\x50\x1c\xf2\x47\xf2\x47\x8e\xf2\xfe\x3d\x64\xd9\x59\x70\xa7\xe5\x93\xb0\x98\xe7\x0f\x62\x95\x20\x48\x03\xcb\xff\xdf\x80\xa5\x07\x25\x36\x38\x0a\x53\x65\x6c\x9b\xe8\x1c\x4e\xb3\x2c\x78\x78\xdd\x62\x40\x0b\x79\x7e\x3a\x1a\x35\xb5\x5e\x6a\x14\x16\x2f\xd3\x64\xb7\x51\x66\x69\x85\xc5\x0d\x2a\x6b\x40\x68\x84\x90\x57\x21\xc2\x58\x2a\x69\x65\xaa\x0c\x48\x05\x91\x8c\x63\xd4\xa8\x2c\x44\x52\x24\x18\x5a\x33\x7a\x12\x7a\x98\xda\x39\x6c\xc4\xf6\xab\xb1\x5a\xaa\xf5\x37\xfe\x2f\x1b\x01\x00\x64\x19\x68\xa1\xd6\x08\x67\xbf\x4f\xe1\x2c\x82\xd9\x1c\x82\x2b\xaf\x1e\xde\xe5\x39\x09\x9d\x66\xd9\x59\x14\xdc\x8a\x8d\xf3\x65\xe6\x1f\x9b\x87\xc0\x19\xf9\x9c\xe7\xa7\xd3\x42\x33\xaa\x88\x74\xe4\xa5\xff\x84\x0a\x2b\xfa\x74\xff\x3f\x87\xaa\x7d\x44\x90\xca\xa2\x8e\x45\x88\x90\xc6\xb4\xe0\xf6\xd2\xd5\x1f\x18\xda\x91\x7d\xdd\x62\xeb\xab\xe5\x3b\xec\xc8\x65\x92\x1a\x1c\x4f\x00\xb5\x4e\x35\xaf\x10\x16\xe3\x09\xfc\x9b\x20\xda\xad\x12\x19\x16\x08\x7d\xd8\xc9\x24\x42\x96\x5b\xa2\xf3\xb6\x29\xc7\xab\x55\xb9\x6b\x65\x50\x1f\xc8\xf1\x6a\x55\xee\xf3\x36\x6a\x39\x97\x57\xab\x72\x57\x98\xe0\xa1\x1c\xaf\x56\xe5\xbe\x3c\xa2\x3e\x10\xa3\xc5\xaa\xd4\x4d\xba\x5e\xa3\x1e\xa7\x7a\x13\xf0\xcf\x89\x87\xfd\xd3\x16\x55\x0b\x7e\xe9\x16\x5d\x8a\x09\x2b\x56\xc2\xb8\x9c\x53\x0a\x43\x97\x6b\xa3\x78\xa7\xc2\x8e\xb7\xc6\x91\x96\x4f\xa8\xdd\xf3\x94\xde\x5d\xa6\x3b\x1d\xa2\x7b\x06\xce\xaa\x09\x8c\x0f\x5f\x9b\x72\x54\x26\x90\x8d\x4e\xa2\x15\x3d\xb9\x44\x33\xdf\x93\xc0\x9d\xd3\xa3\x75\x32\x3a\x91\x31\xc9\xff\x6b\x0e\x4a\x26\x4e\xc3\x89\x46\xbb\xd3\xca\x3d\x92\xaa\xd1\x49\x3e\x2a\xd6\x6e\xf1\xf9\x98\xd5\xab\x02\x98\x56\x59\x60\x45\x06\x84\x22\x4c\x7c\x1a\x42\xac\xd3\x0d\x08\x88\x56\x20\x95\xb1\x42\x85\xc8\x38\x1d\x3b\xd0\xc3\xe2\xce\x05\x17\x9b\xab\x0f\xc7\x11\x2a\x01\xf2\x75\x1e\xdc\xe2\x73\x45\xe5\x50\x4c\x5c\x56\xf8\xe5\xf3\x2c\xab\x70\x44\xaa\x54\xe6\x55\xcf\x20\x72\xa6\xcd\x20\x5a\xe5\x53\xf7\x76\x59\xab\x75\x79\xd0\xb8\xd5\x68\x98\xa2\xe0\xea\x43\x25\x5f\x20\x4e\x35\x6c\x84\x92\xdb\x5d\x22\xac\x54\x6b\x10\xb0\x96\x4f\xa8\x9c\xeb\xbb\xd0\x06\x4e\xdf\x45\x92\x80\xc3\x8b\xc9\x4c\x3c\x09\x99\x10\x57\xda\x94\x0b\x59\x84\x16\x9e\xa5\x7d\x74\xb0\xef\x79\xd6\x3e\x0a\xeb\x38\x42\x63\x22\x2c\x46\x4e\x91\x4d\xc1\x3e\x4a\xe3\x75\x4f\x89\x2f\xa3\x54\x21\xac\x5e\xdd\xbb\x45\x70\x98\x48\xa4\xf1\xe1\x23\x1b\x1e\x52\x58\xa3\x6d\x4a\xa5\x7a\x03\x3b\x83\x94\xf0\x90\x6a\x17\xd0\xbd\xa5\x41\x41\x3f\x0d\x34\xf8\x74\x0a\x16\xe3\x58\x86\xca\x73\xa7\xcb\x73\xa0\x7f\x1c\xf4\xd1\x49\x42\x45\xc9\xcf\x5c\xa0\x0e\x6a\xca\xa1\x71\xe8\x6a\xbb\x7e\xc4\xa4\x4e\x68\xee\x28\x1f\xcb\x30\x88\x56\x81\xdf\xf4\xd1\x62\x7d\x60\x90\xa2\xe3\x4f\x22\xa4\xb8\xb0\x61\x2b\xc2\x3f\xc5\x1a\xfb\x8e\xf3\xfc\xe1\x5f\xae\xd0\x88\x3b\x3a\x0c\xfc\xfa\xdc\x6b\xf7\x07\x33\x99\xee\x6b\x06\x56\xcc\x48\x0e\x58\x1f\xc9\xcb\xfb\xc5\xc5\xc3\x02\x4c\xd1\x26\x7a\x5d\x3e\xce\xd8\x15\x20\xce\x3b\x85\x5c\x41\x6c\x85\x16\x1b\x33\x83\x30\xdd\x6c\x52\x15\xf0\xfe\x1d\x2d\xba\x6d\x17\x1a\xea\xce\xb3\x96\x16\x3e\xf5\x02\xcd\x16\x37\x1b\xd4\x6c\xbf\x86\x41\x59\xb8\x62\x83\xe3\xc9\x37\x6e\x88\xee\x5f\x3e\x1d\x91\x66\x17\x95\x19\x84\xbc\x51\x34\x48\x6e\x39\xbd\x68\x2e\x17\x37\x8b\xcb\x87\x61\x68\x0e\xe8\x6b\x0e\x4d\xe3\x98\xe6\xbc\x53\xa4\x05\x4b\xde\xf7\x58\x76\x82\xe8\x5c\x3d\x29\xfd\xf4\x7c\x64\x02\xd6\x55\x4c\x0f\x30\x87\x73\x13\x18\x52\xe8\xdb\xb6\x8f\xae\xf1\xa0\x70\x7f\xed\x05\xe5\xfa\x76\xb9\xb8\x1f\x08\xca\x80\x26\xde\x95\x62\x35\xa1\x16\x58\x78\xff\x87\x61\x61\x37\x79\x3c\xe8\x75\xf3\xf3\xdd\xd5\xe0\x4a\x1a\x30\x83\x74\xb9\x59\x13\x6a\x71\x93\xf7\x7f\xcc\xcd\x6a\x96\xf3\x80\xd3\xeb\xe9\xd5\xe2\x66\x31\xd4\xd3\x01\x53\x54\x97\xa7\x35\xa1\x16\x4f\x79\xff\xed\x9e\xd2\x8c\xd6\xeb\xe8\x97\xff\x2c\xee\x07\xfa\x79\x74\x0a\xec\xf2\xb2\x2a\x93\x15\xa6\x25\xe9\x1a\x64\xbc\x6f\x16\xcf\xc2\xb8\xf6\xd1\x67\x40\x92\xae\xc7\xa6\x1c\x67\x84\x5e\x1b\x08\x82\xa0\x1c\xc3\xb3\x9c\x3a\x85\x8c\x61\xdf\x2c\x9a\xe3\x09\x4d\x6a\xc5\xf6\xd8\xb0\x96\x20\x08\x8a\x3e\x36\x9f\xcf\xa1\xce\xf6\xf3\xf9\x7c\x7f\x77\x6a\xed\x07\x04\xaa\xe9\xea\x36\x40\x21\x45\x8b\xda\xec\x6f\x12\xed\x8a\xf6\x3d\x9d\xd3\xa0\xad\x73\x8c\x28\xd0\x00\x87\xf8\x14\x6c\x15\xdf\xa6\x76\xf1\x22\x8d\x35\xdc\x8e\xaf\x3f\xc2\xed\xa7\x07\x58\xfc\x76\xbd\x7c\x58\xd2\xa8\xe4\x9a\xb2\x37\xd3\x59\xdc\x0c\xfe\xaa\xa7\xf9\x4d\xaa\xfa\x8f\xb6\xc9\x55\xc1\xb5\x55\xa3\xe6\x60\xf5\x0e\xcb\x4c\x59\x15\x8d\x3c\x55\x16\x5f\x2c\xdb\xec\xc7\x06\x5a\x28\x2c\x76\xa6\x7e\xdf\xa1\x7e\x1d\x64\xa6\x57\x37\x0e\xed\x4b\xa1\x29\xf0\x6b\x83\xad\xbe\xb4\x2f\x30\x87\xd0\xbe\x1c\x18\xeb\xd2\xa4\xce\xd8\x6d\x69\x52\x97\xd8\xa7\x49\xb3\x5b\xf4\xa6\x48\x5d\x49\x57\x8a\x54\x99\xff\x68\x8a\xfc\x2c\xd4\x35\x93\x7e\x04\xea\x83\x26\x37\x00\x6a\x7e\xa7\xf8\x9c\xc1\xd7\x95\xea\x15\xa9\xad\xfd\x96\x3c\x17\xcb\x24\xc1\x88\xe7\xfa\x27\x91\xec\xd0\x10\xf5\xf9\xdb\x81\xbf\xdf\x0f\xf1\xb1\xcd\x8a\xf1\x96\x10\xa6\x95\xc5\x8b\xa5\xbb\x96\xbf\x55\xdd\xf1\xb4\x9b\xe7\xc7\x00\xa8\x7f\xfc\x88\xe9\xe3\x07\x69\xf8\x28\xd1\x65\xcb\xbb\x3c\x27\x21\x19\xc3\x59\x1c\x5c\x9b\x25\x5a\xbe\x98\x54\x36\x54\x6a\x79\xf3\x1e\xe9\xeb\x4c\xe8\x77\x4b\x70\x2f\x8c\x91\x6b\x45\x93\x61\x70\x11\x45\xe3\xd3\x2c\x3b\x8b\xfd\xfc\x93\xe7\xa7\x53\xd8\x06\xb4\xc2\x6e\x4d\x48\x2f\x26\x06\x1b\xa7\xb3\x67\x29\xf1\x2d\x6f\xc9\xb8\xfe\x6a\xe5\x3a\x58\xf9\xf6\xf2\x56\x43\xf8\x37\x9f\xaa\xe5\x46\xe8\xd7\xff\xe2\x6b\xcd\xc8\x2e\xbb\xf2\xfa\xf9\x43\x7e\xb7\x14\x78\x7d\x56\x69\x2b\xf0\xba\x84\x2f\xf0\xd6\x84\xec\xab\xf0\xba\x96\xae\x0a\xaf\x0e\x3d\x47\x2b\x9c\xfb\x7e\x59\xdf\x8d\x26\x5f\x5c\xcc\x8e\x94\x79\xcd\xae\xa2\xfd\x3f\x93\x66\x6f\x13\x2d\x1d\x9b\xf0\xca\xe8\xb3\x51\x73\x20\x15\xbf\x9c\xff\x1b\xd6\x0e\x27\xa5\x6e\x83\xbb\x49\x89\xdf\x69\x90\xd2\x8e\x07\x68\xcf\x33\xce\x6c\x91\x24\x45\x3c\x63\xaa\xe7\x41\xa6\xb7\x29\x7f\x03\xd7\x1c\xf8\x35\x84\x6b\xfc\x67\xd3\x7f\xe8\xe6\xef\xa7\x9b\xfa\x85\xa1\x8d\x6e\xea\x12\x15\xba\x69\xde\x56\x7a\xe9\xa6\xae\xa5\x8b\x6e\xaa\x37\x8f\x81\x74\x23\xb6\xdb\x44\xa2\x81\x82\x22\x54\xe4\xff\x72\x90\xaa\x21\x05\x5c\xb3\xab\xa4\x9b\x3e\xaa\x39\xb8\x62\x1d\x52\xcd\x2f\xa7\x99\x86\x95\xc3\x69\xa6\xdb\xd8\xfe\x31\x93\xcb\xf8\xbd\xff\x08\xc1\x34\xc2\xd9\x91\x65\xef\xfa\x0b\x3a\xcf\x9d\x50\x7b\x31\xd3\xc7\x1e\x5b\x2d\x26\x42\x84\xc8\x8b\xb0\xe0\x3f\x01\x41\xad\x64\x40\x72\x24\x3b\x3e\x77\xf4\x8f\x50\xf5\xe3\xc6\x7c\x12\xad\x2c\xd1\xfe\xd4\x10\x35\xbc\xe0\xe9\xd0\xc9\x01\xce\x6f\x47\xa2\xe3\x8b\x48\x3f\xc1\xff\x32\x24\xba\x5b\xd7\x9b\x90\xa8\x10\x54\xe5\xe7\x5f\x01\x00\x00\xff\xff\x34\x75\xa1\xff\x8e\x1c\x00\x00")

// File30ExecGoTmpl is "30_exec.go.tmpl"
var File30ExecGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x58\x5b\x6f\xdb\x36\x14\x7e\x96\x7e\xc5\x69\xe0\xb6\x52\xe0\xd2\xed\x6b\x0a\x3d\xac\x59\x8a\x15\xdb\xda\xae\xcd\xf6\x12\x04\x03\x4d\x1d\xd9\x44\x28\xd2\x25\xa9\xc4\x86\xa0\xff\x3e\xf0\x22\x47\xf3\x25\xf1\xb0\x74\xd9\x43\xf2\x12\x89\x87\x3c\x97\xef\x7c\xfc\x8e\x92\xc9\x04\xce\x96\xc8\x80\x69\xa4\x16\x0d\x50\xb0\x74\x2a\x10\x2a\xa5\xc1\xce\x11\x66\xfc\x1a\x25\x18\xab\x1b\x66\xd3\xaa\x91\x0c\xb2\x29\x1c\xb7\xed\x88\x7c\x6e\xa6\x82\xb3\xae\x3b\xf5\x27\xdf\x35\x5c\x94\xa8\x73\xef\x2d\xcb\x21\x33\xdf\x04\xf9\x82\xa6\x11\x76\x0c\xa8\xb5\xd2\x39\xb4\x69\x62\x6c\x6d\xc7\x40\xf5\xcc\xc0\x49\x01\x53\xc2\x94\x94\xa4\xe4\x54\x20\xb3\x24\x78\xca\x5e\x4c\xc9\x82\x6a\x5a\x9b\x3c\x4d\xe2\x0e\xa1\x66\xd9\x51\x30\x9f\xc0\xcb\xe7\xd7\x2f\xe1\xf9\xf5\xd1\x18\x6e\xbd\xe5\x69\xa2\xd1\x36\x5a\xae\x7d\x4e\x89\xcb\xe4\x54\x49\x8b\x4b\x9b\x31\x55\xd7\x4a\x92\xf8\xfa\x49\xbf\xa3\xec\x6a\xa6\x55\x23\xcb\xac\x8f\x46\x4e\xed\x32\x1f\xfa\x24\x84\xe4\x69\x97\xa6\x93\x09\x7c\x6b\x50\xaf\x80\x1b\x68\x0c\x96\x30\x5d\x79\x68\xbe\xa2\xcf\xfa\x37\x6f\xa3\xb2\xec\x17\x7e\xe1\x35\xb7\xe0\xb0\xb2\x5c\x49\xb3\x1b\xb5\xb0\x77\x8d\x9a\x0f\x90\x31\xbb\x04\x16\x52\xec\x53\xcd\x21\x3b\xf6\x50\xaa\x1b\x73\x30\x90\xc1\xf9\x5e\x20\x7d\xc6\x87\xe3\xe8\xb7\xaf\x81\xb4\xcb\x3d\x18\x79\x1a\x71\x69\x50\x5b\xe3\xf1\x29\xa9\xa5\x60\xd5\x80\x46\x6e\x65\x4a\x0d\xee\x86\xe4\x83\x3f\x7b\x20\x91\x78\x05\x02\xe5\x6d\xf3\x7e\x30\x86\xcf\x64\x8d\xd2\x9a\x1c\x8a\x02\x5e\xbb\x4d\x7d\x2d\x92\x8b\x31\x54\xb5\x25\x67\xee\x7c\x95\x1d\x49\x65\xe7\x5c\xce\x5c\x76\x21\xe3\xa3\x3c\x4d\xba\x7b\x50\x0d\xf9\xed\x45\x35\x98\x1f\x85\x9e\xff\x1a\xfa\xdf\x17\xe5\xe1\x77\xf8\xc1\xa0\x6f\x7c\xd4\x43\xa0\x0f\xf9\xed\x85\x3e\x98\x1f\x0f\x7a\xdd\xc8\x88\x3b\x0a\xb4\x08\xc6\x52\x8b\x0e\x10\x50\x12\xe8\x46\x07\xc8\xee\x16\xfc\xe8\x8f\x3e\x84\x8c\x06\x4f\x7b\xc1\x0a\xe6\x47\x01\x2b\x48\x65\x4f\xd0\xfd\x74\xdc\x10\x47\x7f\xcc\x81\x71\x71\xe9\x76\x9d\xaf\x16\x48\xce\x96\xf6\x23\xad\x11\xe2\xeb\x67\xca\xae\xe8\x0c\xbb\x6e\x80\x12\x00\x80\x53\xd4\x93\x02\x0e\xcd\xdb\x9f\xd1\xbd\xd4\x06\x6c\xd7\xd2\x9c\x7b\xea\xbb\xf5\x67\x85\xe3\xf5\x16\xcd\x51\x6b\x4f\xe5\x12\x2b\xd4\xde\x0d\x39\x15\xca\x60\x96\xa7\x69\x72\x4d\x35\x64\x69\xe2\x02\x70\x8b\xb5\x81\x03\x8a\xf1\xe9\xb8\x9f\xb6\x05\x5e\x41\xb0\xfe\x44\xcd\x27\x89\xe7\xea\x57\x2a\x57\x5f\x50\x50\x37\x64\xe0\xd5\x60\xf3\x64\x02\xb8\xe4\xc6\x1a\x37\xb0\x28\xd4\x74\xb1\x70\x17\xae\xd2\xaa\x86\x85\xe6\x35\xd5\x2b\xb8\xc2\x95\xbb\x81\x54\x68\xa4\xe5\x0a\x16\x54\xbb\xc1\x16\xc6\xbc\x59\x7b\x8a\x6e\x0a\xa8\xe9\x15\x66\x35\x5d\x5c\xb4\x6d\x4c\x31\xf8\xf9\x19\x57\x77\x57\x70\x79\xdc\x9f\xd8\xb7\x23\x1f\x16\x89\xb2\x5c\x97\x92\xa7\x89\xfb\x08\xf1\x30\x7e\x74\xe4\xf3\xcc\x8f\x05\xb2\x39\xb2\xab\x7e\x56\x02\xa3\x92\xa1\x08\x50\x44\x88\xab\xbe\x81\xcc\x2e\x9d\xf6\x64\xf9\xdb\x61\xeb\x7a\x57\xbe\xdf\x9b\x2d\x74\x8b\x5d\x9a\x24\xae\x4f\x03\x22\x18\xcf\x4a\xa5\xc9\x7b\xae\x8d\xcd\x36\x6e\x9d\x2b\x2e\xcb\xc7\x3d\xd5\xfe\xa0\xa2\x41\x93\x1d\xbb\xf4\xf3\xb6\xbd\xb3\x7b\x9e\xb5\x1e\xea\xb6\x45\x59\x0e\x31\xd9\x26\xdc\x36\xe3\x92\x2e\xfd\x67\x44\x49\x3c\x8f\x5d\x40\x57\x58\x88\x7c\xe1\x8a\x25\x3b\xda\xeb\xea\xea\xba\xcb\xb7\xf1\xc0\x20\x93\x18\x4f\x53\x39\x43\x18\xfd\x39\x86\x51\xe5\xfc\x05\x0f\x5f\xdc\x1d\x40\xc9\xd0\xc4\x90\xb7\xd9\x8d\xaa\xb0\xe5\xab\xe0\x0c\xa3\x35\xf1\xde\x49\xdb\x8e\xaa\x18\x11\x0a\xa0\x8b\x05\xca\x32\xdb\x32\x8d\x21\x26\xbb\x5e\xf1\x1a\x93\x24\xc9\x80\x42\x9b\x6f\x1d\xa0\x30\x18\x20\x0c\x17\x70\x1d\xc0\xbf\x8e\xe1\xd8\xfd\xce\xd7\xc9\xdc\x0b\x09\x14\xf0\xc2\x1f\xbd\x70\xe3\xd0\x3f\xe5\xaf\xde\x5c\xfa\x86\x84\xe0\x2e\x60\x04\xfc\xee\x88\xc3\x4c\xbb\xb5\x06\xc7\x5d\xfe\x06\x78\x0a\x47\x19\x3d\x55\x8d\xb4\x40\xcb\x12\x28\x30\xff\xcc\x94\x68\x6a\xd9\x4f\x7c\xaf\x58\x07\x29\xab\xf7\xd4\x2b\x6b\xa8\x33\x14\xe7\x0d\x0f\xa0\xa5\x83\x5b\x13\x32\x2d\xc0\xea\x06\xff\x6b\x9d\xdd\x2a\xed\x49\x59\xff\xbf\xca\x1a\x48\xf9\x24\xaf\x4f\xf2\xba\x4b\x5e\xc9\xfa\xab\x29\xec\xfb\xfe\x72\xeb\x49\x19\xb9\x1c\x3e\xf3\xab\xb0\xa2\x6e\xc0\xce\xa9\x85\x9a\x5a\x36\x47\x73\xab\xbd\xc4\x1d\xfb\x50\x81\x54\x7e\xd3\x96\x7d\x0c\x54\xc2\x99\xd6\x1f\x95\x7d\xef\x84\x13\x6e\xb8\x10\x30\xc5\x18\x04\x4b\xef\xe0\x7c\xce\x0d\x30\x2a\x44\xbc\x85\x06\xa8\x74\x82\x32\xf3\x7f\x3e\xb9\xc0\x37\xd4\x80\x41\x0b\x37\xdc\xce\x9d\x77\x77\x6a\xaf\xda\xc3\x42\xe3\x35\x57\x8d\x11\x2b\x72\xd0\x70\x08\x9f\x39\x39\x64\xf7\x2a\xcd\x83\x0c\x8a\xb8\xf2\x99\xce\x30\xfe\x0b\xa5\x80\x37\x3b\x6c\x9f\xaa\xca\xd5\x5c\xc0\xeb\xef\x38\x45\x2a\xdf\x96\x93\x62\x28\x92\x69\x54\x8c\x67\xc1\xd8\xa6\xbb\x64\x4e\xe9\x9a\x0c\x3a\x9b\x06\xc5\xeb\xa7\xd1\xc3\x7d\x4d\x8e\xee\xd6\x3b\xc9\xc5\xdf\xc4\x6e\x13\x8a\x7d\x0a\x1d\x05\x7a\x70\x23\x36\x2e\xc4\x5f\x01\x00\x00\xff\xff\x08\xf1\x36\x0f\x30\x14\x00\x00")

// File40SelectGoTmpl is "40_select.go.tmpl"
var File40SelectGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x55\x4d\x6f\xe3\x36\x10\x3d\x47\xbf\xe2\x61\xe1\x83\x1c\xb8\xf2\xa5\xe8\xa1\x80\x0f\x5d\x37\xdb\xaf\xb4\x49\xd7\x41\x7b\x08\x82\x82\xa6\x46\x5e\x36\x12\xe9\x92\x54\xd6\x86\xa0\xff\x5e\xf0\xcb\x6b\xc7\xce\xae\x9c\xa2\x27\x9b\xa3\x99\xe1\x7b\x6f\x3e\x68\xb7\x6b\x42\xd7\x8d\x8a\xdb\x76\x59\x0b\xde\xf7\x0b\xce\xa4\x24\x0d\x21\x2d\xe9\x8a\x71\x42\x97\x01\xc0\x5c\xd5\x6d\x23\x4d\x3e\xc6\xfd\x83\xb1\x5a\xc8\x95\x37\xbf\x13\xda\xd8\xbc\x14\xac\x26\x6e\x11\x3e\x4c\xf0\xc4\xea\x96\x0c\xee\x1f\x4a\x2d\x9e\x48\x17\x7f\xb8\x73\xd7\x89\x0a\xc5\xdd\x76\x4d\xc5\x8f\xcc\xdc\x48\xba\x53\xbf\x32\xb9\x7d\x4f\x35\xb3\x42\xc9\xbe\x9f\x80\x36\xc2\x58\x83\x86\xad\xef\xbb\x2e\xb8\xde\x6a\xd1\x30\xbd\xfd\x85\xb6\xe1\x7c\xb5\xb1\xbf\xb1\x86\x30\x8a\x9f\x19\x7f\x64\x2b\xea\xfb\x87\xcb\x14\xf1\x92\x47\xd7\x91\x2c\xfb\x7e\x8c\xfc\xd2\x11\xfe\xe4\xfb\x48\xe5\x91\xf3\x04\xa4\xb5\xd2\xe3\xac\xcf\xb2\xae\x83\x66\x72\x45\x18\xfd\x35\xc1\xa8\xc2\xb7\xb3\xe4\xfe\x9e\x2a\xd2\x24\x39\x19\xf4\x7d\xf6\x49\x4b\x2d\x9e\x98\xf5\x57\x8e\xaa\xc2\x81\xf9\xdf\x75\x1d\x55\x67\x28\x9b\x9c\xcf\xd1\x36\xc5\x0c\x53\xb7\x1a\x2e\x6f\xd7\x81\x64\xe9\xf4\xcb\xa6\x53\xa4\x22\x06\xd1\xe6\xaa\x95\x16\xc2\x80\x39\x09\x5a\x6e\x51\x29\x0d\xee\xac\x42\xae\xa0\xd5\x47\x03\x55\x21\x0a\xbf\x1f\x99\x9d\xb0\x85\x6c\x31\x51\x50\xff\x4b\x3d\x13\x4b\xe4\x51\x48\xfb\xcd\xd7\x59\x42\xb9\x37\x30\xe4\x4a\xf4\xb6\x15\x75\x49\x1a\x4b\xf7\x6b\xc0\x24\x16\xbf\x5f\x63\x71\x75\x7d\x35\xbf\x83\xb1\xcc\x52\x43\xd2\x62\xcd\x34\x6b\xc8\x92\x36\xd9\xf1\xe4\x1d\x24\xda\xc1\xbc\xf0\x31\x06\x5c\x35\x8d\x92\x45\xf0\xba\xf5\xb6\xec\x82\x2b\x29\xe1\x3a\x7f\xd7\x72\x73\x25\x65\x76\x61\xbc\x97\xd2\x87\xed\xb8\x88\x56\xc7\xa2\x6a\x25\x47\xbe\xc4\xe5\x8b\x10\xc6\x88\x3d\x9b\x8f\x4f\x6e\x88\xa0\xa0\x26\xdb\x6a\x89\x65\x11\x60\x16\xb1\x9f\x8b\x30\x64\x47\x57\x8f\xa3\x82\x7f\x7e\x20\x4d\x60\xeb\x75\x2d\xc8\xe0\xa3\x3f\x71\x25\x4b\xe1\x1a\xd6\x40\x49\xd8\x0f\x84\x7f\x5a\xd2\xdb\x41\x50\x7d\xbe\x3c\xe5\xf1\x4a\x79\xd3\xf8\x33\x51\x4e\xdc\x1d\xee\x00\x68\x16\xa0\x64\x17\x89\x56\x84\x7b\x2d\x1a\x61\x77\x70\x7d\xdf\xd5\xde\xb4\x0f\x14\x9a\xcc\x5a\x49\x43\x83\x10\xfb\x94\x79\xc8\xe2\x7b\x6b\x30\xd4\x5b\xb6\xa2\x22\x20\x9a\x05\x18\x47\x78\x9d\xcb\x21\x5c\x55\x55\x86\x2c\x98\x2c\xff\x33\x72\x97\x3c\x0f\xf9\x26\x78\x35\x81\x9b\x00\x68\x16\x91\x9d\xc1\xef\xc4\x3e\x0e\x63\xfb\x4e\x90\x1b\xbe\xaf\x7a\xbf\x55\x44\x05\xa9\xac\xdb\x8c\x3f\x99\xdd\xa2\x76\x7b\x66\x3a\x45\xc0\xb6\xb7\x9e\xc1\xca\xd2\x60\xdf\x60\x95\xd7\x27\x4c\x12\x95\xe0\xbe\xb1\xdd\xba\x61\x67\xf4\xe5\xd1\x45\xf9\x17\x54\x72\x33\xb5\x2c\xd2\xfc\x16\xc7\x40\x67\xb0\xba\xa5\x83\xd9\x8b\x5b\xb4\x36\x89\xde\xcf\x4a\xc8\x67\xe4\xc0\xf0\xb7\x12\x32\xd6\xbb\x0a\xab\x21\x39\x0c\xa2\xf2\x2c\x69\x6e\xd2\x1e\xf8\xfc\x8b\x77\x1e\xdf\xe7\xc8\x67\x88\xd7\x9c\x22\x1c\x9f\x8d\x83\x5a\xfb\x46\x58\xd4\x82\x93\xef\x83\xe9\x14\x37\xba\x24\xfd\x76\xbb\x9f\xd5\x35\x9e\x72\xe6\x54\xe5\xdd\x14\xb4\xb5\x35\x60\x9c\x2b\x5d\xba\x17\xc6\xaa\x54\x77\x1f\x1e\x96\xdb\x40\xbd\x8e\xef\xcd\x4b\xa1\xd3\x76\xf2\x5f\xbf\x17\xc3\xe4\x89\x83\xe1\x63\x4c\xf1\x5d\x59\xe6\x6f\x0e\xf0\xbc\x99\xa0\x14\x7a\xfc\x5c\x23\xc7\xfe\x07\xad\xda\xf5\x21\xfb\x86\x3d\xd2\x1e\xeb\x95\xf3\xc0\x72\xfb\x7a\xa6\xc7\x77\xe4\xe7\xd0\xf2\xe1\x27\x69\x8d\x5f\xa8\x7a\x9c\xf0\xf4\xd7\xf1\x9c\x2b\x69\x69\x63\x5d\x69\x8d\x27\xc7\xa3\xc1\xb5\xba\x3b\xbb\xf7\x78\xf8\xdc\xc6\x74\x39\xb7\x9b\x94\xa9\x88\xb6\xc1\x6b\x6e\x6e\x37\x98\x81\xdb\xcd\xc1\xfa\xfa\x37\x00\x00\xff\xff\x0c\xe6\x94\xeb\x72\x0b\x00\x00")

// File50SelectorGoTmpl is "50_selector.go.tmpl"
var File50SelectorGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x58\x4b\x73\xdb\x38\x12\x3e\x47\xbf\xa2\xa3\xe2\x3a\x92\xc3\x50\x97\xad\x3d\x68\xcb\x87\x2d\xaf\x53\x9b\x7d\x24\xde\xd8\xbb\x73\x50\xa9\x26\xb0\xd8\xb4\x60\x53\x00\x03\x80\x52\x54\x0c\xff\xfb\x14\x5e\x24\x28\x91\x96\x3d\x99\x9a\xaa\xf1\xc5\x02\xd0\x68\x34\xbe\xfe\xd0\x0f\xce\x66\x50\x55\x51\x72\x2d\xe8\x96\x28\xac\xeb\x1b\xcc\x71\xa5\xb8\x00\x69\x7e\x48\x58\xf1\xbc\xdc\x30\x09\x19\x17\x70\xf3\xdf\x7f\xc3\xd7\x12\x05\x45\x09\x84\xa5\x66\xae\x20\x42\x52\x76\x6f\xd6\x04\xdf\xc9\x91\xda\x17\x38\xa4\x53\x89\x72\xa5\xa0\x1a\x01\x00\x54\x15\x08\xc2\xee\x11\x22\x1a\x43\x94\xc1\xfc\x02\x92\xdb\x7d\x81\xc9\x7b\x8a\x79\x2a\xe1\x5d\x5d\x7b\x39\x9a\x01\xe3\x0a\xa2\x2c\xf9\x20\x3f\x63\x86\x02\xd9\x0a\x1b\x01\xab\xbe\xaa\xa2\x2c\xf9\x48\x36\x58\xd7\x70\xc7\x79\xee\xf7\x62\x2e\x5b\xd1\x7f\x72\xca\x42\xc1\x8e\x99\xc1\xc2\xcd\x8a\x30\x86\xa2\xd1\xc1\xd2\xd0\x9c\x70\xb8\xe2\x25\x53\xe6\x40\x98\xcd\xa0\x94\x68\x51\x91\x5f\x73\xb8\xfc\xf4\xbf\x8f\xb7\x93\xf3\xa9\x83\x70\x54\x8f\x46\xb3\x19\x5c\x3a\x3c\x89\x40\x50\x6b\x04\x46\x36\x28\x81\x67\x0e\x70\x4c\x3d\xe2\xa3\xac\x64\x2b\x98\x48\x38\xef\x05\x73\xea\x35\x4d\xa6\xb0\x58\x4a\x25\xb4\x13\xaa\xd1\xab\x2d\x11\x5a\x83\x6c\x26\x0f\xc0\xfe\xf9\xc7\xc0\xa6\x19\xc8\xe4\x18\x71\xeb\x52\x0b\x48\x2e\xe1\x02\x48\x51\x20\x4b\x27\x7a\x14\xc3\xd8\x88\x5a\x7b\xeb\x7a\x3c\x35\xc2\xf5\x09\x74\x5f\x09\x54\xa5\x60\x46\xa1\xc3\x4e\xbb\xcf\x22\xf7\xc0\x29\x03\x5e\x28\xca\x99\x01\x4f\x23\xa9\x99\xb9\x3f\x89\x9a\xd1\x61\x30\x5b\xf1\xcd\x86\xb3\x44\x4f\x5c\x13\x41\x36\xd2\xc3\xf7\x60\x8e\xe9\x11\x78\x1a\xca\x06\xac\x43\x38\xa3\xcc\x0a\xdc\xe4\xf4\x00\x49\xff\x30\xe6\x17\x20\x93\x03\x72\xfe\xb5\x5d\x7e\x7d\x01\x8c\xe6\x01\xc8\xb3\x19\x90\x34\xd5\x1e\xb7\x40\x64\xc0\x19\x82\xe2\xb0\x21\x6c\x0f\x02\x73\xa2\x91\x89\x1d\x95\x40\x8b\xa8\x35\x0a\x30\x6f\xb3\xe0\x94\x29\xa9\xa5\xd5\x9a\x4a\x33\xd7\x28\xb6\x57\x6f\xdc\x67\x86\x5a\xcd\x01\x10\xad\x25\xfa\xef\x3d\x17\x48\xef\xd9\xbf\x70\x3f\xf7\xa2\xed\x54\x57\xd4\x19\xdf\xda\xd5\x98\x00\x6a\x4d\x94\xf1\x63\x60\x2a\x95\xd6\x5a\x7d\x53\xce\x8e\x34\x59\x46\xcd\x1d\xbf\xda\x33\x5b\xaa\xc5\x7d\xa7\x07\x07\x28\x72\x97\xe3\x91\xcc\x67\xcc\x6e\xf5\x42\x8f\x66\xbf\xf4\x0c\xdd\xee\x92\xe6\x5e\x4f\x61\x1e\x9c\x3a\x78\xa1\x66\xed\xe8\xdc\xba\x3b\xb4\x44\x77\xa1\x61\xde\x50\x28\x69\x82\x45\x2b\x5e\x1f\xbe\xc3\x30\x52\xfe\x20\x3b\x0d\x2d\x9d\x47\xbd\x83\x5b\x0c\x08\x6b\x61\x7a\x23\xa1\x10\x74\x43\xc4\x1e\x1e\x71\x1f\xea\x68\x76\x4a\x58\x58\xf3\x97\x81\x8e\x85\xf7\xc4\x32\x59\x34\xe8\x2c\x7f\x57\x22\x77\xbd\xf5\x04\xe7\x0e\xf9\x64\xa2\xc1\x30\x8d\x8e\x88\x60\xe4\xaf\x2d\x4a\x4f\xd0\xfb\x37\xa4\x02\x4b\xa1\x3e\x1c\xf9\x80\x6c\xb0\x6c\xb2\x99\x4e\x7f\xda\x4f\xa2\x44\xd8\xad\x91\x01\x39\x4c\x7c\x20\xd7\xbc\xcc\x53\xb8\x43\x1d\xb4\x30\xb5\xaf\xe0\xb9\xf1\xda\x9c\x30\x99\xda\x14\x6b\x7d\xe0\xec\x90\x89\x49\xbe\xce\x92\xf7\x54\x48\x65\xcd\x91\x2b\xa2\xd3\x04\xf3\x35\x89\xa1\x1c\x54\x95\xc5\xd1\xf2\xd7\x49\x9a\x82\xe4\xa4\x0d\xad\xee\x49\x4a\x89\x9e\x05\x9b\x5b\x63\xd8\x12\x93\x6a\x53\x41\xb7\x28\x92\xff\x93\xbc\xc4\xaa\xa2\x99\x4b\x09\xff\x20\xf2\x13\xc3\x5b\xfe\x1f\xc2\xf6\x9f\x5d\x54\xae\xeb\x18\xf0\x1b\x95\x4a\xc2\x86\x14\x0b\x6f\x56\xe0\x5e\x33\xbe\xfa\xa6\xb4\xa5\x10\xb9\x65\xb2\x7a\x24\xf7\x58\xd7\xcb\x73\xbf\x63\x48\xa2\xaa\x90\xa5\x75\x3d\x85\xc9\x79\xcf\x9d\x63\x40\x21\xf4\xa5\x2c\x96\x72\x47\xd5\x6a\x0d\xfe\x5a\xbe\x34\x7b\x17\xe6\x38\xbf\x38\xbf\x80\x28\xf9\xbb\x1d\x48\xcf\x8f\x15\x91\x68\x68\xea\xa4\xdc\x59\xe3\x79\x43\xae\xc6\x5d\xda\x2f\x47\x82\x13\x8d\xa0\x81\x2c\x7a\x0e\x66\xfe\x72\x7d\xc5\x43\x8a\x19\x29\x73\x75\x74\x32\xa3\x79\x0c\xd9\x46\x25\x57\xfa\xe6\xd9\x64\x5c\x32\x59\x16\x05\x17\xba\xda\xf2\x97\xfb\x93\x1c\xc7\x7e\xe0\x5f\x43\xdd\x10\xeb\x59\x9c\x7a\x11\x9d\xfe\xa0\x4c\x1a\x12\x3d\xa0\x15\x55\xb8\x31\x53\x36\x79\x1c\x3f\x20\x7b\xdf\x67\x5f\xb0\xe3\x76\x9a\x19\xcd\x47\x79\x27\x74\x37\x0a\x11\x84\x34\xb7\x72\xa6\xcd\x4a\xba\x4e\x8b\xb5\xbc\x8e\x21\xdd\xba\x6e\x90\xf3\xb3\x19\xf4\xf2\xd8\x31\xe4\x78\x3e\x20\x4b\xf4\x6b\xd8\xd2\x7f\xda\x64\x88\x2e\xd1\x0b\xf8\x12\xbd\x9c\x30\xd1\x0b\x18\x13\x9d\x0a\x3e\xba\xd8\x9e\xb4\xee\xe3\x3b\xe8\xdb\xd4\x08\x90\x3c\x07\xcd\x26\x9b\xc9\xfe\x96\xe7\x93\x69\xb3\x46\x81\x06\x92\x82\xef\xae\xec\x45\x9b\x56\x70\x3a\x54\xbc\x47\x9d\x46\xe8\x54\x1f\xd4\x96\x47\xda\x9a\xef\xdf\x4f\x74\x43\x34\x33\x3c\x5f\xd0\xa5\xa7\xeb\xd9\x19\xbc\x6e\xcd\xab\x46\x81\x6b\x2f\x39\xdb\xa2\x50\xc6\x97\x97\x3c\xf5\xe8\x42\x94\xb9\x53\x5b\x42\xb7\x36\xea\x3e\xfc\x09\x9f\x37\xc1\xb1\xdd\x81\x5f\xc1\x99\x0a\xc7\x04\x30\xd3\xe1\x1e\x5d\xad\xaf\x71\xf5\xa8\x77\xee\xd0\xb0\x91\xe9\x1c\xae\x8b\x33\xfd\x9c\x74\x11\x5f\x08\xdc\x52\x5e\x4a\xf3\x01\x20\x81\x0f\x46\x32\xa5\x69\x0c\x12\x15\x7c\x69\xae\xfb\x25\x0e\xd5\x4a\xee\xaa\x65\xff\x89\xa1\x69\x07\x74\x65\x04\x3b\xce\xde\x28\x5d\x34\xe0\x96\xe4\x25\x51\x98\x26\xf0\x13\x02\x67\xf9\x1e\x18\x62\xaa\x81\x2d\x51\x86\x1a\x33\xc1\x37\xbe\x02\xd7\x3a\x64\x12\x3a\xc2\x52\x7f\x21\xf8\x2e\xe9\xa3\xbe\xf5\xdd\xf2\x38\xac\x74\x09\x75\x61\x4a\x9d\x7e\x77\x84\xc9\x68\x60\x8a\xbe\x7d\xdb\x53\x75\xb7\x5c\x0b\x1a\xc5\x2e\xd5\x2a\x03\xd9\x23\x2d\x20\xb3\x25\xa9\x2e\x97\xfd\xa7\x85\x61\xed\x7d\x85\x5c\xdb\xc8\xdb\xaf\x17\xed\x55\x5d\x35\xb0\x25\xb9\x7e\x19\x8e\xb9\xc9\x44\x17\xe1\xd3\xb0\xc7\xd7\x39\x9f\x32\xf5\x97\x3f\xcf\x0f\x51\x4a\x6c\x61\x65\x36\x77\xe5\x17\xcb\xbb\xbd\xc2\xe1\x0d\xae\xe0\xbe\x26\x42\xe2\x07\xa6\x74\x80\x6b\x9f\xf7\x51\x6a\x3f\x8c\xf7\x6e\xf7\x95\x10\xee\x15\x4d\xc6\xbe\x0c\x1d\xc7\x40\x63\x7f\x99\x18\xc6\xc6\xf0\xd8\xd9\x33\x9e\xf6\xb8\xb2\x85\xf1\x44\xd0\x08\x5a\xfe\xd6\x5b\x0f\x03\x0d\xd3\xc3\x31\xb5\xd4\xa6\x68\xb2\xe4\x83\xcd\x92\x93\x9e\x5a\xca\x5b\x3f\x5f\x4e\x3b\x84\xee\x4b\x82\x43\x89\xb0\x2f\x76\xf4\x7f\x98\xf0\x8e\x09\x23\x5a\xd3\x48\x1d\x2c\xc4\x60\x92\x4e\x48\xde\x6b\xdd\xa2\xa1\xa8\xeb\x73\x97\x0c\xd4\xa6\x98\x76\x9e\x44\xd8\x67\xf6\x1f\xd6\x09\xbf\x83\x5a\x87\xde\x59\x2f\xdf\x7d\x09\x20\xf8\xce\xa7\xfc\x40\x42\xbf\x2d\x9f\x55\x9c\xa8\xeb\x69\x8c\x1d\xbe\x93\xd9\x11\x09\xb2\xc0\x15\xcd\xe8\x8a\xe4\xf9\xbe\xf9\x68\x77\x3a\x8b\xb7\x39\xab\xa7\x95\x79\xc6\xb7\xd0\x10\xe7\x6e\x42\x82\xd7\x7d\x39\xe8\xec\x0c\x1c\x52\xee\x9f\x16\xf3\x0d\xd3\x2f\x01\x00\x00\xff\xff\x54\x39\x08\xb6\xf6\x15\x00\x00")

// File60WhereGoTmpl is "60_where.go.tmpl"
var File60WhereGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x92\x41\x8b\xdb\x30\x10\x85\xcf\xd6\xaf\x18\x16\x1f\xec\x62\xb4\xf7\xc2\x5e\x52\x5c\xea\x4b\xba\x94\x40\x0e\x21\x14\xc5\x1e\xdb\x22\xb2\x64\x64\x39\x69\x10\xf3\xdf\x8b\xe5\xb8\x04\x4a\x68\x28\x9b\xa3\xde\xcc\xbc\x79\x7c\x1a\x77\xe9\x11\xbc\x8f\xf9\xfb\x78\x50\xb2\x24\xda\xb6\x68\x71\x35\x4a\x55\xa1\x85\xc1\xd9\xb1\x74\xe0\x89\x31\xef\xad\xd0\x0d\x42\xfc\x33\x83\xb8\x86\xcf\x6f\xc0\x37\x97\x1e\xf9\x57\x89\xaa\x1a\x88\xd8\xeb\xeb\xdf\x3e\xde\xc7\x35\x5f\x8b\x0e\x89\x40\x54\xd5\x00\x02\x4a\xa3\x2b\xe9\xa4\xd1\x60\x34\xdc\xd6\x9d\x01\xd7\x22\x6c\xbf\xe5\x3f\x72\x18\x9c\x70\xd8\xa1\x76\xac\x1e\x75\x09\xc9\xa7\x7b\x11\xd3\x5b\x8f\xc4\xf4\x50\x9a\xae\x33\x9a\x7f\xef\x33\x38\x09\x35\x57\x43\xd0\xfc\x97\x9b\xba\x20\x9e\x9f\xef\xa2\x3c\x8a\x06\x89\xd2\x65\x24\xd8\x82\x67\x91\x45\x37\x5a\xbd\xc8\x6b\x3c\x87\x4a\x62\xfa\x0c\x5e\xa6\x1c\x61\x7e\x23\x0e\x0a\x89\x5e\x66\xad\xe6\x5f\x8c\x1a\x3b\x1d\x84\x93\x50\x29\x23\xf6\x2f\x22\x85\xbe\x32\xd1\x50\xac\x9f\xca\xa5\xd0\xc9\x49\xa8\x01\x38\xe7\xff\xc3\x43\xd8\x66\x98\x3e\xbc\x13\x47\x4c\x76\x7b\xa9\x1d\xda\x5a\x94\xe8\x29\x03\x85\xb3\x77\x9a\xb2\xa8\x36\x16\xe4\xd4\x38\x5f\x4a\x58\xe9\x59\x14\xe6\x77\x72\x0f\x6f\x41\xda\xc9\x3d\x8b\xe8\x1e\xe5\x42\x27\x0f\x31\x9e\x3c\x39\xe7\x8f\x70\x5e\xa1\x3b\x23\x2e\xb0\x61\x95\x6f\xb6\x79\xfe\x5c\xe0\xd7\x95\x89\x32\xe7\x0c\x5a\xd9\xb4\x1f\x78\x87\x8b\xf7\x43\x98\xfe\x04\x98\x40\x79\x8f\xba\x22\x62\xbf\x03\x00\x00\xff\xff\xa5\x82\x54\xfa\xf4\x03\x00\x00")



func init() {
  if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}



var err error



  




  
  var f webdav.File
  

  
  
  var rb *bytes.Reader
  var r *gzip.Reader
  
  

  
  
  
  rb = bytes.NewReader(File00ImportGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "00_import.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(File10OrmGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "10_orm.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(File30ExecGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "30_exec.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(File40SelectGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "40_select.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(File50SelectorGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "50_selector.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(File60WhereGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "60_where.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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


