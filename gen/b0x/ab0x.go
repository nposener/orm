
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



// FileExecGoTpl is "exec.go.tpl"
var FileExecGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x98\x4d\x6f\xdb\x38\x13\xc7\xcf\xd6\xa7\x98\x1a\x68\x2b\x05\x86\xd4\xa7\xe8\x73\xc9\xc2\x87\x36\x9b\x02\x01\x16\x69\x76\xdb\xee\xa5\xe8\x81\xa6\x46\x32\x11\x8a\x74\xc9\x91\x5f\x10\xf8\xbb\x2f\x86\x94\x6c\xd9\x8d\xd1\x60\x37\xdb\xcd\xcd\xe1\x0c\x87\x33\x3f\xfe\x39\x14\xb3\x10\xf2\x56\xd4\x08\x77\x77\xf9\x4d\xfc\xb9\xdd\x26\x89\x6a\x16\xd6\x11\xa4\xc9\x68\x2c\xad\x21\x5c\xd3\x38\x19\x8d\x4b\x41\x62\x26\x3c\x16\xfe\x9b\x3e\xfe\xbb\x28\x9d\x5a\xa2\xe3\xe1\xaa\x09\xde\x0e\x2b\x8d\x32\xfc\x6c\x8d\x17\x15\xf2\xaf\x5a\xd1\xbc\x9d\xe5\xd2\x36\xc5\xc2\x7a\x34\xe8\x0a\xeb\x9a\x71\x92\x00\x00\x8c\xef\xee\xf2\x4f\x9b\x05\xe6\x57\x61\xf9\x1b\x41\xf3\xed\x76\x9c\x64\x49\x52\x14\x70\xb9\x46\x09\xd2\xa1\x20\xf4\x20\x80\xc4\x4c\x23\x54\xd6\x01\xcd\x11\x6a\xb5\x44\x03\x9e\x5c\x2b\x29\xa9\x5a\x23\x21\x9d\xc1\xd9\x45\xf0\x7e\xd7\x2a\x5d\xa2\xcb\x42\x84\x34\x83\xd4\x7f\xd3\xf9\x1f\xe8\x5b\x4d\x13\x40\xe7\xac\xcb\xe0\x2e\x19\x79\x6a\x68\x02\xc2\xd5\x1e\xce\xa7\xc0\x29\x1a\x93\x97\x4a\x70\x0d\x79\x8c\x94\xbe\x98\xe5\x0b\xe1\x44\xe3\xb3\x64\xd4\x79\x68\x5b\xa7\xe3\x68\x3e\x87\x97\xcf\x97\x2f\xe1\xf9\x72\x3c\x81\x7d\xb4\x2c\x19\x39\xa4\xd6\x99\x5d\xcc\x59\xce\x99\x5c\x44\xae\x69\xc7\xf7\x83\x7b\x27\xe4\x6d\xed\x6c\x6b\xca\xb4\x5f\x26\xbf\xa0\x75\x36\x0c\x96\xe7\x79\x96\x6c\x03\x8f\x6f\x2d\xba\x0d\x28\x0f\xad\xc7\x12\x66\x9b\xc0\xe1\x23\x86\x74\x7f\x0f\x36\x61\xca\x7e\xe0\x37\xd5\x28\x02\x06\x43\xca\x1a\xbf\x47\x14\xed\x3b\x44\x21\x68\x2a\x69\x0d\x5d\x5a\x79\x97\x66\x06\xe9\x59\xe0\x66\x57\xfe\xc1\xd4\x62\xf0\x93\xd4\x42\x96\x0f\x87\x16\xdc\x77\xd4\x68\x7d\x82\x4b\xd0\x89\x32\x1e\x1d\xf9\xc0\x84\x55\x0a\x64\x07\x3a\xe9\x75\xbb\xc7\x70\x15\xfc\x1f\xa8\x14\x55\x81\x46\xb3\xdf\xa4\xb7\xde\xab\xda\x34\x68\xc8\x67\x30\x9d\xc2\x2b\x76\xea\xf3\x37\x4a\x4f\xa0\x6a\x28\xbf\xe4\xf9\x55\x3a\x36\x96\xe6\xca\xd4\x9c\x51\xcc\x72\x9c\x25\xa3\xed\x0f\x48\xc6\xfc\x4e\x92\x8c\xe6\x9f\xab\xbf\xbf\xc5\xf9\xf3\xa2\x7c\xf8\x89\x7c\x34\xce\x6d\x58\xf5\x21\x9c\x63\x7e\x27\x39\x47\xf3\x7f\xc0\xd9\xb5\xa6\x83\x8c\x1a\x09\xc1\x93\x20\x64\x12\x60\x0d\x88\x23\xdc\xf9\x9e\xf7\xaf\xc1\xfd\x31\x3a\x60\x8c\x74\x92\x4c\x34\xff\x5c\x32\xb1\xcb\xf5\xd2\x3b\x14\xda\x51\x5f\x0b\xae\x5c\xf9\x97\xaf\xfd\x1d\x73\xb9\xa6\x6b\xd1\xe0\x76\x3b\x40\xc0\x97\x10\x37\xbf\xf3\x29\xfc\x30\xa9\xe0\xec\xfa\x76\x18\x89\xed\xda\x67\x16\xd4\xcb\xe3\xcf\xa6\x2c\xcd\xef\x94\x8a\xce\x05\x35\x96\x58\xa1\x0b\x61\xf2\x0b\x6d\x3d\xa6\x59\x92\x8c\x8a\x02\x70\x4d\x4e\x48\x0a\x16\x96\x70\xbc\xda\x5a\x87\x3e\x19\x2d\x85\x03\xa1\x35\xdc\x53\x4b\x32\xe2\x1b\x31\x84\xbb\x66\xb4\x61\x5f\x39\xd1\xa2\x00\x39\x47\x79\xdb\xd7\x05\x52\x18\x89\x5a\x0b\xbe\x0f\xa2\x4b\x97\x30\xd7\x4e\x6b\x3e\x46\x69\xf6\xcb\xb0\x84\x3e\x54\xa8\xfb\xb8\x14\x1e\xdc\x26\xa3\x91\x22\x6c\x06\x40\x7c\xd8\x06\xeb\xf2\xf7\xca\x79\x4a\x8f\x34\xc5\x39\xa7\xd9\x04\x96\x42\xb7\xe8\xd3\x33\xce\x3b\x8b\x60\x07\x09\x0d\x08\x7e\x8f\x90\x19\x8e\x18\xc6\x14\xc4\x62\x81\xa6\x4c\x85\xd6\x13\x38\xe3\x3c\xe2\x79\xef\xa6\x84\xe1\x00\x26\x54\xd6\x49\xe8\xc2\xb6\x86\x40\x94\x25\x08\x90\xe1\xb7\xb4\xba\x6d\x4c\xdf\xc7\xc2\x86\x9e\x54\x55\x98\x7d\xa8\xaa\xb8\x0d\xc1\xf0\x4f\x74\x35\x20\x17\xd3\x9a\x02\xb9\x16\x9f\x92\xe6\x06\x95\x3e\x6d\xd5\xc5\x5d\x7a\x72\xd2\x0b\xb9\x75\x25\xc5\xde\x5e\xc5\x11\xbb\x02\x9a\x0b\x82\x46\x90\x9c\xa3\xdf\xeb\x30\xe7\x69\x57\x15\x18\x1b\x9c\xbe\xb3\x4f\x40\x18\xb8\x74\xee\xda\xd2\x7b\xd6\x15\xac\x94\xd6\x30\xc3\x6e\x11\x2c\x43\x80\x4f\x73\xe5\x41\x72\xde\x71\x33\x3c\x08\xb3\x81\x85\xa8\xc3\x65\xc9\x0b\xaf\x84\x07\x8f\x04\x2b\x45\x73\x8e\xce\xb3\x0e\x94\x0f\x0b\x87\x4b\x65\x5b\xaf\x37\xf9\xc9\xc3\x11\x4f\x3c\x7f\x37\x3e\x6a\xc7\xdd\x8d\xdc\x88\x1a\xbb\xef\xda\x29\xfc\xef\x1e\xdb\x87\xaa\xe2\x2a\xa6\xf0\xea\x5f\x3c\x36\x55\x00\x7d\x3e\x1d\xaa\x3f\xe9\x14\xf4\x2c\x1a\xef\x92\xfb\xf4\x6b\x5d\x93\x0f\xf6\x2a\x89\x52\x0e\x33\x1f\xa9\x85\x1e\x57\x75\xea\x14\x75\x87\xa8\x1b\x8e\xab\x1f\xaa\x35\x6c\xf1\x7d\xdb\x73\xef\x43\xe1\x68\xa0\xfb\x92\x63\xd7\xe9\x0e\xf0\x20\x8f\xde\x7d\x10\xf6\xe0\xf4\x48\x5a\x77\x27\x26\xd6\xc8\x4f\x1e\x01\x73\x21\x6f\xfb\x26\xdd\xbf\x4b\xfa\xf7\x5f\x51\x80\x57\x46\x62\x30\xba\xbd\x01\x4a\x8b\x1e\x8c\x25\xc0\x35\x3f\x3e\x41\xd1\x4b\x0f\x5a\x78\x92\x56\xfb\x2e\x3c\xef\x0c\x08\x58\x89\x0d\x87\xe7\x0f\x2b\x8e\x27\x40\xb6\x9e\x6c\x03\x5e\x0a\x63\xd0\xf5\x4b\x7f\x94\xc2\x40\x83\x34\xb7\xf1\x70\x7d\x44\x04\xe5\x7d\x8b\x30\x27\x5a\xf8\xf3\xa2\x18\x3c\x78\x6b\xab\x85\xa9\x8b\xda\x16\xc1\xc5\x17\xaf\x5f\xff\xff\xcd\x9b\x48\xb7\xdb\x3f\xb7\x2b\x26\x83\x2f\x5f\xe3\x93\x3a\xff\x93\x6d\x4c\x8d\x0b\xb3\x0d\x42\x5b\xeb\xcd\x8e\x80\x90\x12\xfd\xa0\x8c\x4a\xa1\x2e\x93\x91\x0b\x1f\x71\xdd\x1b\x3c\x86\xf8\x50\xa5\x2f\x5c\x96\x5f\x6a\x6c\x18\xb1\xab\x82\x87\xcf\xdf\xf3\x8c\x77\x9b\x20\xa5\x71\x1f\x67\xdc\x5d\x0a\x76\x89\x4e\xf2\xa2\x0b\x61\x94\x3c\x0c\x98\x5f\x19\x42\x57\x09\x89\xe7\xdc\x4c\x98\x6c\xb7\x69\xa1\x1c\xb0\x33\x12\xca\x60\x09\x95\xb3\x0d\xb4\x86\xb1\x3b\xe2\xbf\x79\x45\x06\x1d\xd1\x85\x5c\xf6\xc9\x5e\xe3\xea\x2d\xa5\xae\x0a\x5d\x83\xc5\x1d\xff\x7d\x90\xdf\x58\xc5\xeb\xb1\xe5\x73\x18\x79\x5b\x96\x2e\xcd\xb2\x41\x49\x71\x75\x57\xed\x33\x4b\xb3\x3c\x3d\x24\xc9\x92\xfe\x2b\x00\x00\xff\xff\x09\x6c\x73\x77\xf6\x10\x00\x00")

// FileOrmGoTpl is "orm.go.tpl"
var FileOrmGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x57\x4f\x6f\xdc\xb8\x0f\x3d\xc7\x9f\x82\xbf\x20\x28\xec\x62\xea\xdc\x07\xc8\x21\x4d\xa6\x40\xf0\xcb\xa6\x69\x26\x3d\x15\xc5\x56\xb6\xe9\x89\x76\x6d\xc9\x95\xe4\x64\x03\xc3\xdf\x7d\x41\x51\x9e\xd8\xf3\xaf\xe9\x76\x2f\x9b\x4b\x64\x52\xa2\xde\x7b\x94\x48\x4d\x23\xf2\x3f\xc5\x0a\xa1\xeb\xd2\x5b\x1e\xf6\x7d\x14\xc9\xba\xd1\xc6\x41\x1c\x01\x00\x74\x1d\x18\xa1\x56\x08\x27\xbf\xcf\xe0\x24\xb8\xe6\x67\x90\xde\x3f\x37\x98\x5e\xf9\x6f\x0b\xef\xfa\xde\xcf\x3e\xee\xba\x30\xa7\xef\x8f\x87\xf5\xa8\x0a\x18\xfc\x85\x70\x22\x13\x16\x4f\xed\xf7\xea\x38\x3a\x3a\x5e\x49\xf7\xd0\x66\x69\xae\xeb\xd3\x46\x5b\x54\x68\x4e\xb5\xa9\xf7\x7b\x4e\x73\x5d\xd7\x5a\x1d\x98\x50\x48\x51\x61\xee\x8e\xa3\x24\x8a\x4e\x4f\xc1\x89\xac\x42\x90\x16\x96\x9f\xae\xc3\x87\x12\x35\x46\xb9\x56\xd6\x05\xc3\x19\xe1\x66\x42\xf7\x64\x20\xec\xb4\x36\x37\x28\x1c\x5e\xe8\xaa\xad\x95\x5d\x3a\xe1\xb0\x46\xe5\x2c\x08\x83\x90\xb3\x15\x0a\x2c\xa5\x92\x4e\x6a\x65\x41\x2a\x28\x64\x59\xa2\x41\xe5\x20\xe0\xb0\xd1\xa3\x30\x7b\x23\x9d\x41\x2d\x9a\x2f\xd6\x19\xa9\x56\x5f\xf9\x5f\xb7\x43\xf6\xc2\x2b\x7e\x19\x22\x4e\xd4\x2e\xd2\x1b\x51\x13\xe2\x79\xf8\xdc\xdc\x04\x4e\x3c\xb3\xbe\x3f\x9e\x8d\x13\x42\x31\x7a\xcf\xf2\xfc\xf6\x8a\xf4\x71\x0f\x08\x52\x39\x34\xa5\xc8\x11\x74\xe9\x0d\x1f\xef\x7e\x03\x9d\xfd\x81\xb9\x8b\xdc\x73\x83\x3c\x77\x3d\x89\xa1\x5e\x54\xda\x62\x9c\x00\x1a\xa3\x0d\x5b\x3c\xdb\x38\x81\xb7\x3c\x7a\xdf\xca\xaa\x40\xf6\x2d\x91\x38\x90\x8f\x47\x63\xdf\x95\xb2\x68\xbc\x8f\x47\x63\xdf\xe7\xa6\x08\x31\x79\x34\xf6\x5d\x62\x85\xec\xe3\xd1\xe0\xf3\xce\x6b\xbd\x5a\xa1\x89\xb5\xa9\x53\x1e\x26\x81\xf7\xa7\x16\x8d\x44\xb3\xcd\xbd\xd4\x06\x04\x2c\x17\xd7\x8b\x8b\x7b\x7f\x6c\xec\xa0\x25\x8b\xb0\x5e\xb8\x21\x04\xd9\x9f\xe3\x04\xe2\x2f\x5f\x87\xe3\xb4\xf8\xcb\x71\x7a\x66\x2c\xcf\xb0\xf7\x85\x6e\x69\xf1\xcf\xec\xed\x5d\x39\xad\x93\x6a\x05\x4d\x6b\xe8\xd0\x5b\x46\xb4\x0e\xb7\x99\x1a\xb2\x4f\x11\x31\x1c\xef\xd8\xc0\xf4\x41\x1a\xfb\x4f\x30\xad\xd0\x79\x48\x5a\x55\xcf\xb4\x92\x62\x95\x14\x0b\xa4\xc3\x3a\x05\x59\x82\xd2\x7e\x0c\xb5\x70\xf9\x03\x72\xfc\xef\xa4\xd6\x0c\x84\x82\x6f\x94\x9b\x85\x31\x37\xda\x7d\xd0\xad\x2a\xbe\xc1\x93\xac\x2a\xc8\x10\x0c\xba\xd6\x28\x2c\x52\xa6\xb9\x46\x38\xa2\x79\xe4\x8d\xc4\xf1\xed\x8f\x44\xff\xd8\xa0\x02\xdd\x20\xdd\xda\x50\x86\x20\xd7\x4a\x61\x4e\xd7\x37\x2a\x5b\x95\xfb\x39\x71\x61\xe4\x23\x1a\x8a\x31\xf3\x33\x97\xba\x35\x39\xd2\x37\xf0\x1d\x4d\x20\x3e\xbf\xbd\x1a\xc2\x13\x8a\x22\xf3\x5f\x74\x4f\xed\xf7\x2a\xfd\x41\x98\x24\x3a\x92\xa5\x9f\xff\xbf\x33\x50\xb2\xa2\x08\x47\x4c\x96\x3e\x7d\xa8\xe8\xa8\x8f\x8e\x8a\x75\xd4\x50\x4e\xd2\x1b\x7c\x1a\x05\x7e\x75\xa0\x60\x7b\x43\x7c\xbb\x10\x6b\x0e\xc5\x0c\x8a\x6c\x0e\x45\xd6\xcf\x68\x7a\xd0\xe9\x06\x9f\x82\xf2\x96\xd2\x43\x4b\x42\x11\x80\xd2\xe8\x1a\x04\x14\x19\x48\x65\x9d\x50\x39\xb2\x6c\x53\x54\x41\x25\x0a\x0e\x94\xdb\xcb\xf7\x3b\x04\xfb\x97\x98\xd1\x31\xff\x19\x72\x5c\x91\x5e\xf8\x41\xc6\xc5\x82\x2a\x9e\x50\xfe\x78\x5f\xdc\x2d\xce\xef\x17\xa3\x5b\xef\x29\xc6\x39\xbc\xa5\x0d\x92\x7d\xe5\x8d\x10\x0e\x48\x26\x0e\x42\xde\x08\x23\x6a\x3b\x07\x6e\x5f\x29\xfb\x6f\xbd\x91\xdc\xc4\xc2\xb7\x9e\x39\xb7\xa4\x59\xb0\x6d\x16\xf3\xf9\xbe\x4e\xf2\x25\x4f\xd7\x3a\x8a\x1a\xe3\xe4\x2b\x57\x7b\xfa\xeb\x67\x91\x0f\x46\xe8\xe7\x90\xb3\x63\xa8\xfe\x5c\x85\x0f\xea\x11\x6e\xfe\x5e\x3d\xf6\x94\x74\xd2\xc3\x52\x82\xdf\x4c\xcc\x3b\xd4\x60\x7f\x50\x63\x2c\x03\x21\x3f\x5a\xc3\x0e\xd9\xb6\x29\x2f\x1f\x3a\x1d\x9c\xc1\x1b\x9b\x5a\x1f\x23\x34\xa0\x90\x06\x1b\x38\x72\x37\x39\xc8\xf1\xea\x66\xb9\xb8\x3b\xc0\x71\x4f\x6b\x1a\xe7\x7c\xe2\xd8\xc1\x92\xfd\xaf\x61\xc9\xa8\xb9\xcf\x1d\x44\xfd\xf9\xf6\xf2\xe0\x49\xdd\xd3\x34\xc7\xa8\x27\x8e\x1d\xa8\xd9\xff\x43\xd4\xe3\x23\xc5\x4d\xf8\x20\xf0\xcb\xc5\xf5\xe2\x10\xf0\x3d\x1d\x7d\x0c\x7c\xe2\xd8\x01\x9c\xfd\x3f\x05\x9c\x33\x34\xed\x94\xe3\x4a\xb8\xeb\xa0\xac\xc9\x95\xb2\xaa\xb0\x80\x27\xe9\x1e\xe0\x51\x54\x2d\x5a\xcf\x17\x56\xf2\x11\x87\xfa\x19\x68\x66\x1b\xa7\x28\xd9\xb9\x73\xdc\xc0\x76\x4b\xdb\x75\x00\xbb\xee\xdd\xf8\xb1\x58\xbe\x3c\xcf\x3f\x48\xac\x0a\x4b\x8f\x6f\x3f\x49\x96\x70\x52\xa6\x57\x76\x89\x8e\x1f\xbe\xe4\xc8\x86\xeb\x74\x6e\xad\x5c\x29\x5f\x4d\xd2\xf3\xa2\x88\xe9\x39\x59\x86\x4b\x46\xaf\x47\x68\x52\x6f\x61\x20\x5d\x37\x44\xbb\x43\xff\xe2\xcd\x29\x1c\xcf\xf0\x7b\xdf\x1a\x59\x0b\xf3\xfc\x7f\x7c\x5e\xaf\x40\x55\xf4\x7d\xc2\x58\x86\x37\xe8\xfa\xa3\x7f\xe9\x52\x19\xe5\xe3\xe5\x0a\x6c\x24\xa4\xe5\x7b\x11\x34\xf6\x8f\x93\xaa\xa2\x96\xd3\x52\x83\xf2\x8c\x5f\x74\x9e\x1c\xef\x64\x67\xc0\x7d\x3a\x6f\x5d\x99\xd7\xe8\xcc\x2f\xec\xff\x98\xd4\xaf\x21\x16\x75\xdd\x36\x29\x8a\xea\x7b\x88\x1b\xc1\x05\x8b\xce\x72\x7a\xc2\x9b\x95\x68\xc1\x84\x22\xfd\x54\xa2\x47\xe0\x9e\xb2\xbb\x7d\x41\xa6\x5b\xc4\x1c\xdd\x5b\x96\xe8\xf8\xf7\xcd\xae\xab\xf1\x7a\xc5\x7d\xc4\x64\xaa\xcb\x2f\x51\xdb\x53\x9b\xb7\xcf\xe4\xeb\xa8\x6d\x9d\xc6\x5f\xa3\x36\xfa\x15\x38\x1a\xfe\x1d\x00\x00\xff\xff\xc2\x6c\x02\x83\x10\x10\x00\x00")

// FileOrmCommonGo is "orm_common.go"
var FileOrmCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x55\x4d\x6f\xe3\x36\x10\x3d\x9b\xbf\x62\xe0\x93\xbc\x28\xe8\x5f\xe0\xc3\x26\x56\x51\x03\x41\xb6\x4d\x5c\x6c\xaf\x14\x35\x96\xd9\x52\xa4\x4a\x8e\xd6\x0e\x16\xf9\xef\x05\x3f\x1c\x47\x4a\xa4\x14\x8b\xcd\x85\xf1\x70\xe6\xcd\xe3\x7b\xa4\xa6\x13\xf2\x1f\xd1\x20\x50\xa7\x3d\x63\xaa\xed\xac\x23\x28\xd8\x62\x29\xad\x21\x3c\xd3\x92\xb1\xc5\xb2\x51\x74\xec\x2b\x2e\x6d\xbb\xee\xac\x47\x83\x6e\x6d\x5d\xbb\x9c\xdc\x59\x4b\xdb\xb6\xd6\xcc\x24\xd4\x4a\x68\x94\xb4\x64\x2b\xc6\xd6\x6b\x90\xd6\x18\x70\xd8\x39\xf4\x68\xc8\x83\x80\xed\x4d\x0c\xa2\x24\x65\x0d\x1c\xac\x83\x56\x18\xd5\xf5\x5a\x90\x32\x0d\x08\x68\xd4\x37\x34\xe0\xc9\xf5\x92\x78\xc0\xf8\xac\x35\x1c\x7a\x13\x0b\x3c\x88\x6f\x42\x69\x51\x69\x04\xb2\xa0\x0c\xa1\x13\x92\xe0\xa4\xe8\x08\xc2\xc0\xe3\x1f\x77\x40\x69\xf7\x28\x08\x94\x07\x87\x5a\x10\xd6\x01\x88\x2c\xd0\x51\xf9\x8c\xfd\x0b\x08\x87\x50\x5b\x83\x50\x3d\x85\x5a\x65\x3c\x09\x23\x11\xec\x21\xe5\xd9\xea\x6f\xcc\x1c\xf6\x16\x1a\xa4\x71\x96\x75\x2d\xf4\x1e\xe1\x4b\x87\x06\xac\x83\x7b\x3c\x5d\x99\x72\x46\x4f\x1d\x26\x05\x52\x47\xf8\xce\x16\x59\x1f\xc8\x2b\xdf\xa6\x95\x2d\xea\x0a\xe2\x9f\x75\x2d\xdf\xde\xb0\x85\xb6\x4d\x83\x2e\xfd\xbe\x8b\xff\xb3\x67\xc6\x02\x3a\x14\x12\x3e\x05\xd8\x15\xdc\x6a\xeb\xb1\x58\x01\x3a\x67\x5d\x80\x77\x48\xbd\x33\x20\x79\x5d\xf1\xbc\x19\xca\xd6\x6b\x48\x18\xe0\x31\xba\x90\xd1\xa3\x22\x99\x63\xbe\x30\xe3\x16\xa9\xae\xc8\x05\x57\x36\xab\xd0\x4e\xf2\x1c\xdf\x64\xc4\xdc\xec\xd6\xa1\x20\xbc\xe9\x95\xae\xd1\x41\x15\x56\x7f\xb1\xe7\xf6\xa1\xfc\xbc\x2f\xc1\x93\x20\x6c\xd1\x10\x74\xc2\x89\x16\x09\x9d\x4f\x8a\x0d\x8b\xaf\xd2\xc5\x3c\x0f\xe9\x0a\xf2\x94\xf5\x7b\x8c\xb1\x45\x3c\x01\x24\xce\x99\xc3\xee\x70\x6f\xa9\x3c\x2b\x4f\x3e\x9d\x7a\xf7\x2b\xdc\x7f\xd9\x43\xf9\xd7\xee\x71\xff\x18\x6f\x5e\x38\x7b\xa6\x13\x98\xbd\x50\xca\x1a\x54\xf0\x69\xc0\x65\xf5\x1a\xb3\x58\x8d\x76\x03\xc5\x8a\x27\x92\xfc\x75\xf3\x0d\x90\xeb\xf1\xc5\x9b\xea\xa2\x51\x7a\x87\x89\x5b\x76\x21\x06\x2e\xcc\x02\xa5\x7f\x7b\x74\x4f\x93\x74\x32\x44\x21\xe9\x7c\xa9\xe6\x39\x36\xcb\xee\x96\xce\xb0\x01\x49\xe7\x37\xa4\x76\xc6\xa3\xa3\xb7\xc6\xed\xee\x1f\xcb\x87\xfd\x8c\x69\xc3\xc2\x29\xd3\x52\xd6\x8c\x69\x3f\x22\xca\xa0\xf5\x47\xa2\x0c\x79\xfe\x3f\x51\xfe\xec\xea\xb7\xb7\x39\x10\xf9\x50\x95\x61\xe5\x94\x2a\x29\x6b\x46\x95\xaf\x47\x74\x78\xd5\xe4\xeb\x6f\xe5\xc3\xeb\x07\x94\x5f\xf1\x3b\xd2\x0c\xfa\xaf\x12\x4e\x71\x8a\x68\xb9\x77\x0c\xad\x46\x99\x03\x5d\x52\xf3\x0d\xc4\xb2\x9f\x72\x8b\x47\xac\xe6\x0d\x9b\x26\x36\x6d\xd8\x16\x35\xbe\x6b\xd8\xb6\xbc\x2b\x67\xbf\x3d\xc3\xca\x29\xc3\x52\xd6\x87\x86\x89\xae\xd3\x0a\x3d\x5c\x04\x37\xb5\x4a\x33\xcc\x9a\x29\x69\x06\xfd\x5f\x0c\x1b\x9b\x35\x64\xf9\xae\x59\x3f\xc5\xa8\x11\x9b\x79\xa3\xa6\x49\x4d\x1b\xa5\x6d\x03\xea\x70\x9d\x47\x27\x11\xbf\xd5\xe3\x19\xa4\x6d\x53\xc4\xb1\xad\x4c\x13\xc6\x76\xe3\x81\x73\x1e\x87\xff\x41\x48\xfc\xfe\x1c\x87\x91\x3a\xc0\x75\x1e\x6d\xc0\x28\x1d\xa2\xb9\x25\x5b\x3c\x5f\xc7\x55\xe1\x13\x0a\xe7\x3c\x8c\xc7\xff\x02\x00\x00\xff\xff\x4d\x16\xa8\xf0\x2b\x09\x00\x00")

// FileParseGo is "parse.go"
var FileParseGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x92\x4d\x8b\xea\x30\x14\x86\xd7\xcd\xaf\x38\x37\x20\xb6\x58\x4b\x14\x75\x21\x78\x17\x77\x21\xdc\x9d\x0c\xe3\x4a\x5c\xd4\x4e\x5a\x0e\x93\x26\x25\x39\x1d\x18\x06\xff\xfb\x90\x58\x6a\x5d\xcc\x87\xc3\xec\x52\x78\x3f\x9e\xbc\x69\x93\x17\xcf\x79\x25\x81\x1a\xe5\x18\xc3\xba\x31\x96\x20\x66\x11\x57\xa6\xe2\x2c\xe2\x8e\x6c\x61\xf4\x4b\x77\x44\x5d\x39\x7f\x24\xac\x25\x67\x09\x63\x65\xab\x0b\x68\x72\xeb\xe4\x7f\x4d\xb1\x83\xc3\xf1\xf4\x4a\x32\x01\xd4\xb4\x5a\xc0\x1b\x8b\x30\x05\x69\x2d\xac\x37\xd0\x25\x65\xbb\x5e\x1d\xf2\x62\x97\xa4\x30\x13\x29\xac\x16\x09\x8b\xb0\x0c\xf2\x3f\x1b\xd0\xa8\xbc\x3f\x52\xa6\xca\x76\x16\x35\x95\x31\xdf\xe6\xa8\xe4\x53\xe8\x43\x5d\xc1\xc8\x01\x19\x5f\xc5\x53\xe8\xc3\x12\x16\x9d\x59\x64\x25\xb5\x56\x03\xb2\xf3\x90\x71\x7f\x0b\xd9\x7e\x45\xb9\xc7\xdf\xc3\x6c\xef\xe0\xdc\x2a\x93\x0f\x41\x4b\xff\xfd\x29\x69\xe7\xb8\xa2\xfe\x14\x33\x54\x7d\x97\xf3\x11\x6b\xd9\x63\xa6\xd0\x58\x59\xa0\x43\xa3\xfd\x9b\x24\xe0\x7f\x92\xcc\x4b\x7c\x73\x69\x6c\x9d\x93\xe7\xe6\x73\x21\x56\x53\x31\x9b\x8a\x39\xcc\x96\x6b\xb1\x58\x8b\x25\x0f\xac\x57\xff\x5f\x10\x01\xb7\x73\x4d\x36\xc0\x33\x0e\x93\x8e\xca\x65\x0f\xb2\x91\x39\xc5\x5c\xf0\x41\xeb\x85\x94\xfa\x81\x42\x7f\x58\x27\xbe\xe4\xdc\xde\xea\xae\x71\xc6\x23\x37\xf6\xf3\xf4\x77\xfa\x70\x22\xba\x9d\xe8\x9f\x31\x6a\xf0\x92\x27\x63\x42\x57\x27\x76\x07\x71\xf4\x04\x82\x9d\xd9\x7b\x00\x00\x00\xff\xff\xeb\xe7\x7c\x37\x8a\x03\x00\x00")

// FileSelectGoTpl is "select.go.tpl"
var FileSelectGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x56\x4b\x8f\xdb\x36\x10\x3e\x9b\xbf\x62\x20\xf8\x20\x2d\x1c\xf9\x52\xf4\x50\xc0\x87\xc6\xdd\x14\x29\x16\x89\xdb\x0d\xda\x43\x51\x14\x34\x39\xf2\xb2\x91\x48\x85\x1c\x65\x6d\x08\xfa\xef\x05\x1f\x72\xfd\x5e\x6c\x91\xcb\xae\x38\x9a\xf9\xe6\xfb\x38\x0f\xab\xe5\xe2\x33\xdf\x20\xf4\x7d\xb9\x8a\x8f\xc3\xc0\x98\x6a\x5a\x63\x09\x72\x06\x00\x90\x09\xa3\x09\xb7\x94\xb1\x49\x26\x39\xf1\x35\x77\x38\x77\x5f\xea\xb9\xb4\xea\x2b\xda\x8c\xb1\x49\xb6\x51\xf4\xd4\xad\x4b\x61\x9a\x79\x6b\x1c\x6a\xb4\x73\x63\x9b\xb9\x30\x4d\x63\x74\x16\x61\xfa\xbe\xfc\xb4\x6b\xb1\x7c\x1f\xc0\x57\x9c\x9e\x86\x21\x63\x05\x63\xb4\x6b\x11\x1e\x05\xd7\x1a\x2d\x28\x4d\x68\x2b\x2e\x10\xfa\x10\xf6\x4e\x59\x47\xb9\x54\xbc\x46\x41\xe0\xc8\x2a\xbd\x99\xc1\x57\x5e\x77\xe8\xe0\xcf\xbf\x22\x87\xf2\x77\x7f\x2e\x20\xbf\xeb\xfb\x69\xcc\xf2\xc1\xe8\x95\x09\x60\xc3\x30\x03\xb4\xd6\xd8\x82\x0d\x8c\xf5\x3d\x58\xae\x37\x08\xd3\xbf\x67\x30\xad\xe0\x87\x05\xa4\x88\xdf\xb0\x42\x8b\x5a\xa0\x83\x61\x88\xa4\xfa\x7e\x5a\x95\x1f\x78\x83\xc3\xf0\xed\xf8\x55\x37\x09\xf6\x3d\xa0\x96\x9e\x01\x9b\xcf\x61\xbc\xb3\xc8\x61\x69\x3a\x4d\xa0\x1c\x70\x9f\xa8\x13\x04\x95\xb1\x20\xbc\x55\xe9\x0d\x58\xf3\xec\xc0\x54\x90\xa8\x1f\x46\xb2\x0b\xb6\x88\x96\x80\xa2\x98\xd1\xe1\x7e\x4b\x29\xce\x5b\x53\x5a\x4d\xdf\x7f\xc7\x22\xad\x47\xf4\x6a\xdf\x76\xaa\x96\x68\x61\xed\xff\x3b\xe0\x1a\x1e\x7f\x7d\x80\xc7\xfb\x87\xfb\xe5\x27\x70\xc4\x09\x1b\xd4\x04\x2d\xb7\xbc\x41\x42\xeb\x52\xa5\x8f\x82\xf7\xf9\x27\xc1\xcf\x41\xec\x99\x32\x7a\xad\x82\x8d\x4d\x84\xd1\x1a\xee\xfc\x5f\x36\x71\xe1\x8d\xb1\x30\x3e\x78\x52\x55\xa7\x05\xe4\x6b\xb8\x3b\x42\x2f\xc6\xb6\xca\xf7\x4f\x49\xa9\x45\xea\xac\x86\x75\x19\xb3\x96\x4b\x53\x77\x8d\x76\x65\x7e\x37\xa2\x16\x49\xeb\x1f\x4f\x68\x11\x78\xdb\xd6\x0a\x1d\x3c\x87\x93\x30\x5a\x2a\x52\x46\x3b\x30\x1a\xe8\x09\xe1\x4b\x87\x76\x77\x95\x45\xc0\xc8\xc7\xd8\xa0\x2f\x98\x8a\x13\x4f\x7f\x0d\x7b\x4a\x31\xf1\x22\xa6\x64\x93\x91\x71\xa2\xf5\xa0\x1a\x45\x7b\x5a\xa1\xf4\x75\x30\x1d\x12\x02\x8b\xae\x35\xda\xe1\x55\x66\x01\x26\x8f\x91\xa1\xc2\x37\x29\xad\xf8\x06\xcb\x98\x79\x11\xd3\x9d\xf1\xf2\x2e\xc7\xb4\x4c\x55\x39\x24\xe0\x5a\xfe\x2f\x86\x1e\x30\x8f\x18\x33\x78\x15\xd1\x8f\x31\xf1\x22\x31\x78\x85\x8e\x0b\x4b\x22\x0e\xc6\x3b\x85\xbe\xd5\xdf\x0c\x61\x50\x55\x05\xda\x10\x4c\xab\xf2\xbd\xdb\x6f\x0f\x3f\xba\xfb\x11\x39\x58\x20\xc0\xa5\x74\x87\x1b\x05\xc8\x84\x7b\x88\x0d\x87\x12\x44\xe8\x41\x3f\xc1\xfc\x85\x7e\x3a\x03\xcf\x2f\xdc\x86\x6f\xf3\x75\x39\xb6\x73\x79\x4e\x68\x01\x64\x3b\x3c\x1a\x87\xb4\x80\x6a\x37\xca\xf8\xc5\x28\x7d\x22\x02\x38\xfc\x63\x94\x4e\xf5\xf3\x1b\xe8\xc0\xe1\x2a\xe5\x13\xa0\xdc\x8d\xe3\x78\xb6\x62\x5f\x96\x72\x4a\x6a\x01\x09\xed\x92\x96\xff\x96\xe9\x47\x2b\xd1\xbe\xdd\x1d\x46\xfa\xfe\x30\xde\x3c\x16\x63\xdf\x94\x5d\x4d\x0e\xb8\x10\xc6\x4a\xbf\x5b\xc9\x8c\xe5\x09\xe1\x71\x5d\xdc\x90\x7b\x9e\x2b\x97\xca\x8e\xc3\x1f\xde\xfe\xa4\xae\x4b\x4d\x7d\x1a\xfc\x5c\xf9\xa3\x94\x79\x76\x94\x37\x9b\x81\x54\xb6\x38\xd5\xeb\x55\xfe\x6c\x4d\xd7\x1e\xab\x6c\xf8\x67\x3c\x50\xb7\xf1\x1e\xb0\xde\xbd\x4e\xd1\x39\x6e\xfe\x12\xfd\x10\x72\x91\x7e\x71\xa5\x52\x6f\x52\xa9\x96\xf1\x93\xc3\xd7\xc7\x05\xe6\xe9\x1b\x24\xb4\x9b\x3f\xfb\x5f\x9a\xdb\x33\x92\x20\x72\x41\xdb\x31\xba\x4c\xb6\x9b\xab\x63\x49\x5b\x58\x80\xa0\xed\xd1\x4a\xf8\x37\x00\x00\xff\xff\xba\xba\x66\x7b\x28\x09\x00\x00")

// FileSelectorGoTpl is "selector.go.tpl"
var FileSelectorGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x56\x4d\x6f\xe3\x36\x10\x3d\x47\xbf\x62\x56\xb0\x17\x52\x56\x2b\x5f\x8a\x1e\x5c\xf8\xb0\x48\x77\x81\xf4\x63\xbb\xdd\xa4\xbd\x04\x41\x41\x4b\xa3\x98\x29\x45\x2a\x24\x65\xd7\xd0\xea\xbf\x17\x24\x45\x49\xb6\x95\x38\x05\x9a\x4b\x2c\x72\x3e\xc8\xf7\x66\xde\xb0\x22\xd9\xdf\xe4\x01\xa1\x69\xd2\x2f\xee\x67\xdb\x06\x01\x2d\x2b\x21\x35\x44\xc1\x45\x98\x13\x4d\xd6\x44\xe1\x42\x3d\xb1\x45\x2e\xe9\x16\x65\x18\x5c\x84\x45\xa9\xc3\x00\x00\xa0\x69\x40\x12\xfe\x80\x30\xfb\x2b\x81\x59\xe7\xb8\x5c\x41\x7a\xbb\xaf\x30\xbd\xb6\xdf\x0a\xde\xb7\xad\xb5\x0e\x9b\xa6\xb3\x69\xdb\xde\x1f\x79\x0e\x7e\xff\x81\xea\x4d\xbd\x4e\x33\x51\x2e\x2a\xa1\x90\xa3\x5c\x08\x59\x2e\x32\x51\x96\x82\x87\x41\x1c\x04\x99\xe0\x4a\x03\x4a\xf9\xab\x7a\x80\x15\x84\x99\xe0\x5b\x94\x9a\xf2\x07\x98\xab\x25\x64\x82\xd5\x25\x87\x79\x0e\x3b\xaa\x37\xb0\x25\xac\x46\x98\x6f\x21\xd2\xfb\x0a\x61\x7e\x1b\x83\x16\x30\x57\x61\x10\x2c\x16\xa0\x90\x61\xa6\x85\xec\x7e\xa8\xce\x5b\x41\x21\x24\xdc\xfc\xfe\x0b\x3c\xd5\x28\x29\x2a\x20\x3c\xb7\x6b\x15\x91\xca\x64\x32\x7b\x52\xec\x54\x60\xa3\x0e\x61\xb4\xac\x33\x0d\xcd\x11\x32\x34\x81\x59\x31\x80\xf2\x89\x22\xcb\x07\x4c\x9a\x06\x68\x01\x5c\x68\x98\x15\xe9\xb5\xfa\x8a\x05\x4a\xe4\x19\xf6\x06\x37\x36\x7c\xd3\xcc\x8a\xf4\x33\x29\xb1\x6d\x61\x2d\x04\xeb\xd1\x63\x6a\x30\xfd\x49\x50\x3e\x36\x1c\xfd\xbe\xc9\x08\xe7\x28\x27\x40\x3f\xfc\xca\x44\xcd\xb5\xcd\x00\x8b\x05\xd4\x0a\xdd\xcd\xd5\x13\x83\xab\xdf\xfe\xf8\x7c\x1b\x5d\xc6\x1d\x4c\x41\x6b\x41\xbc\xea\x30\x23\x12\x41\x6f\x10\x38\x29\x51\x81\x28\x3a\x58\x30\xf7\xa8\x06\x45\xcd\x33\x88\x14\x5c\x7a\xc0\x62\xef\x1c\xc5\x70\x77\xaf\xb4\x34\xd8\x36\xc1\xc5\x96\x48\xe3\xa4\xfa\xc5\xff\x17\x50\x5a\x80\x4a\x4f\x51\x75\xb4\x39\x0c\x98\x82\x15\x90\xaa\x42\x9e\x47\xe6\x2b\xb1\xa5\x5b\xa4\xee\xbc\x6d\x1b\xc6\xd6\xb8\x7d\x11\xce\x0b\x89\xba\x96\xdc\x86\xeb\xc0\x32\x04\x39\xa8\x1e\x05\xe5\x20\x2a\x4d\x05\xb7\x68\x19\xe8\x4c\xb9\xed\xa7\x60\xb2\x6e\x16\x24\xd7\x09\xa9\x59\xf0\x48\x3d\xda\x98\x07\x5b\x2f\xe3\xd5\x23\xa2\x0e\x21\x39\xae\x9e\x37\x2b\xe0\x94\x8d\x70\x71\x99\x7a\x60\xec\x67\x02\xa3\xbc\x83\xa9\xf9\x73\x60\x2d\x8f\xa1\x4b\x0e\x8c\xbe\x62\x71\x4b\xd6\x0c\xbd\x99\x3d\xa1\x5d\x99\x32\x3d\x0c\x69\x6d\xbf\x48\x5a\x12\xb9\xff\x19\xf7\x53\x29\xda\x69\xa2\x3c\x35\xf6\x0a\x7d\x21\x9b\xca\xa7\x0a\xb4\xac\x11\x76\x1b\xe4\x40\x8e\x6b\x1e\xd4\x46\xd4\x2c\x87\x35\x02\xc9\x73\xcc\x8d\x9a\xbc\xc8\x9c\x0d\x1a\xc5\xae\xa1\x1c\x3c\x5d\x6a\x95\xda\x56\xeb\x92\x7f\xa2\x52\x69\x77\x02\x95\x11\x53\x23\xdc\xab\x8c\xc9\x41\x8c\x42\xdb\xeb\x3a\x6e\x3a\x4b\xab\x37\x53\x69\x87\x70\x51\x4e\x89\x59\x05\xd7\x49\x89\x91\x44\x53\x2d\x4e\xcb\xd3\x3f\x8d\x40\xc6\x10\x5d\x4e\xc4\x4f\x8c\xcc\x9a\x68\xee\xdc\x6a\x47\x75\xb6\x01\x1f\xcf\xab\xdc\xfb\xf1\x00\xf0\x9b\xcb\x15\xcc\xd2\x1f\xdd\x87\xea\xa5\x85\x28\xb4\xcc\x75\x56\x5d\xae\x70\xd9\xb3\xd5\x43\x63\x30\x38\x31\x8c\xcc\xd1\xe3\x31\x93\xbe\x7c\x73\x2c\x48\xcd\xf4\x49\x20\x4e\x59\x02\x45\xa9\xd3\x8f\xe6\x22\x45\x14\xd6\x5c\xd5\x95\x99\x3f\x98\xf7\x17\x99\xab\x30\xf1\x1f\xbe\x5a\xda\x9e\x93\x57\xd1\x71\x8e\x89\xff\x4a\xc2\xc7\x7f\xb4\x0b\x7c\x44\x01\xd5\x58\xda\x25\x03\xb0\x4a\x4f\x59\x76\x91\x63\xdf\xd3\xc6\xf2\xa4\x8b\xc7\xd8\xa0\x94\xa3\xfe\xe8\x76\xde\x9a\x34\xe9\xe1\x0d\x13\x63\x6f\x6a\xf5\x70\xe0\x3f\xcb\xb7\x99\xae\x53\x1c\x76\x70\x9e\xae\x8f\x90\x9d\xbd\x12\xda\xe9\x04\xd1\xb3\xd8\xce\xce\x55\xb8\x51\xd3\x68\xc0\x49\xec\x60\xca\xa9\x37\x20\x8c\x81\xa1\xc1\x9d\xe8\x03\x63\x51\xdc\xef\x51\xa0\x9d\x65\x3c\xf1\x4e\x2a\x1c\x60\xe3\xe9\x75\x6e\x78\x0d\x42\x6d\xd2\x7e\xfb\x76\x66\x84\xd1\xc2\x56\xc2\x1d\xbd\x1f\x0a\x60\x84\xd5\x95\x7b\x35\x59\x70\xae\x44\x8e\x30\x2b\xba\x04\x43\x35\xd8\x30\xef\xde\x1d\xe9\xa7\x79\x6e\x4c\x9c\xe5\xec\xec\x38\x8d\x34\x35\x32\x87\x61\xe4\xde\x21\x83\x7f\xa7\x3e\x5b\xc2\x0c\x76\xdd\xdd\x52\xfb\xac\x8b\xc7\xa3\xdb\x68\x0c\xe5\xfa\xfb\xef\x96\x07\xb3\x43\x8a\x5d\xea\x44\xd3\x3a\x1f\xda\xdf\xdd\xaf\xf7\x1a\x9f\x77\x30\xcf\x3e\xbc\xe6\xda\x94\xd6\x40\xf1\x89\xe8\xbc\x20\x3c\xee\xbd\x9a\x40\xe8\xe7\x49\x98\x00\x4d\xfc\x35\x46\x3f\x42\x7b\xf6\xa4\x3b\x52\x18\xbf\xc8\xc9\x99\xca\x1a\xcd\xf9\x81\xb1\x47\x27\x1d\x47\x7c\xfd\x00\x8f\xa7\x94\x99\x76\x30\x0a\x62\x7b\x64\xdc\x16\xe9\xc8\xd3\xc9\xd1\x0a\x1e\x9d\x1a\x45\x13\xfa\xee\xaf\xb7\xbc\x8f\xc7\xf5\x39\x29\x4e\xcf\x09\xd4\x00\xc1\x64\xc9\x78\xe5\x92\x62\xe7\x95\x6a\x64\xd1\x3f\xf5\x3f\x30\xd6\x99\x76\x53\xde\xf6\x9a\x9f\xed\x3b\xa2\x40\x55\x98\xd1\x82\x66\x84\xb1\x7d\xff\x82\x9d\x14\x9f\xa1\xe9\x27\x86\xfb\x2b\xde\xaa\x4d\xf3\x5c\xa3\xc3\x9b\xa9\xde\x7e\xfb\x16\x9a\x06\x79\x6e\x3c\xed\x3f\x63\xe6\x9f\x10\xff\x06\x00\x00\xff\xff\x0a\x92\x10\x58\xc8\x0d\x00\x00")

// FileWhereGoTpl is "where.go.tpl"
var FileWhereGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x52\x4d\x8b\xdb\x30\x10\x3d\x5b\xbf\x62\x30\x39\xd8\xe0\xca\xf7\xc2\x5e\xb6\xb8\xd4\x97\xb4\x94\x40\x0e\x21\x14\xad\x3d\xb6\xc5\x5a\x23\x23\x2b\x49\x8b\x98\xff\x5e\x22\xc7\xed\x42\xbb\x10\xd8\xbd\xd9\x33\xef\x43\xef\x49\x93\x6a\x9e\x55\x8f\x10\x82\xfc\xb6\x7c\x32\x0b\xa1\xcd\x64\x9d\x87\x4c\x00\x00\x84\x00\x4e\x51\x8f\xb0\xf9\x51\xc0\xe6\xb6\xfa\xf8\x00\x72\xf7\x6b\x42\x59\xc7\xff\x19\x3e\x30\x47\x74\x1a\xc2\x0d\xc3\x9c\xae\x7c\xa4\x16\x98\x45\x92\xf6\xda\x0f\xa7\x27\xd9\x58\x53\x4e\x76\x46\x42\x57\x5a\x67\xca\xc6\x1a\x63\x29\x15\xb9\x10\x21\xbc\x30\xeb\xfe\xfa\x7c\xd6\x38\xb6\x33\xb3\x28\x4b\xd8\x0f\xe8\x30\x84\x4d\x27\xb7\xca\x20\x33\xa8\xb6\x9d\x41\x41\x63\xa9\xd5\x5e\x5b\x02\x4b\xf0\x72\xef\x2d\xf8\x01\x61\xff\xa5\xfa\x5e\xc1\xec\x95\x47\x83\xe4\x45\x77\xa2\xe6\x1f\xb1\xcc\x4e\xb0\x9c\x47\x7e\x9d\x0a\x38\xab\x71\x91\x8a\xa7\xa8\x7e\xfa\x05\x95\xaf\x98\x48\x87\x20\x12\x87\xfe\xe4\x68\x1d\x6f\xf1\x12\x37\x99\x9d\x8a\xd8\xc9\xc2\xdf\xa9\xa7\x11\x99\xd3\x65\xd6\xc9\x4f\x76\x3c\x19\x8a\x83\xb3\x1a\x73\xc1\xe2\x7f\xf9\x6a\xba\x25\x24\xa8\xb7\xef\x93\xb2\xa6\xec\xac\xc6\x19\xa4\x94\x77\xa5\x53\xae\x9f\xaf\x97\x61\xd4\x33\x66\x87\xa3\x26\x8f\xae\x53\x0d\x06\x2e\x60\xc4\x45\x2c\xcf\x45\xd2\x59\x07\xfa\x0a\x5c\x6e\x31\x7a\x04\x91\x44\xfe\x41\x1f\xe1\x21\x8e\x0e\xfa\x28\x12\x7e\xad\xb3\x9a\xb2\xbb\x1a\xbb\x6a\x4a\x29\x5f\x6b\xed\x11\xfd\x05\x71\xad\x0e\x1e\xab\xdd\xbe\xaa\xde\xa9\xbe\x9b\x76\x36\xda\x4b\x01\x83\xee\x87\xb7\xbc\x91\x55\xec\xae\xd0\x7f\x1c\xaf\xb1\x43\x40\x6a\x99\xc5\xef\x00\x00\x00\xff\xff\x18\x22\x51\xf0\xc6\x03\x00\x00")



func init() {
  if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}



var err error



  




  
  var f webdav.File
  

  
  
  var rb *bytes.Reader
  var r *gzip.Reader
  
  

  
  
  
  rb = bytes.NewReader(FileExecGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "exec.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(FileOrmGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "orm.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(FileOrmCommonGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "orm_common.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(FileParseGo)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "parse.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
  
  
  
  rb = bytes.NewReader(FileSelectorGoTpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "selector.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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


