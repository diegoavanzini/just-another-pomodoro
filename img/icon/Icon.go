// Code generated for package icon by go-bindata DO NOT EDIT. (@generated)
// sources:
// img/icon/Icon.png
package icon

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _imgIconIconPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x78\x67\x50\x13\x5c\xd7\x6d\xe8\x12\x9a\x34\x41\x4a\x68\x4a\xef\x08\x4a\xef\x20\x0a\xd2\x3b\xa1\x88\x09\x84\x92\x20\x84\xde\x51\x9a\x05\x6c\x20\x1d\x2c\x28\xd2\x82\x20\x2d\x94\x80\x34\x15\x84\xa8\x60\x84\x10\x82\x4a\x09\x84\xde\x02\x09\x81\x3b\xcf\x33\xdf\x37\xef\xdc\x3b\xef\xdc\x3d\xb3\x67\xce\x3e\x7b\xad\x75\xce\xdf\xb5\x72\x1d\x6e\x58\x73\x01\x45\x80\x00\x00\x80\xcb\xe6\xaa\x85\x13\x00\x00\xa0\xfe\xd3\x67\x58\x01\x00\x40\xce\x9a\x97\x21\x00\xc0\x70\x17\x66\x6e\xee\x60\x63\x6e\x2e\x19\x11\x89\x80\xc2\xc2\x20\x00\x80\xfc\xe3\xe4\xc7\x86\x57\x07\x4c\xc4\x86\xfc\x9c\xef\xbf\x14\xb4\xd4\xcf\xb9\x72\xab\xe6\xc6\x39\x0b\x4b\xb1\xea\x7b\x96\x4d\x0e\x33\xc0\x06\xbd\x22\x1f\x2a\x8b\x9d\xc8\x13\x01\x21\xeb\xfe\xd6\xf9\xf9\x91\x9b\x53\xfd\x43\x9e\xee\x6c\xee\xa6\x67\x8f\x01\xb7\x23\x91\x4e\xc8\xc7\x7c\xa7\xcf\x05\xbc\x06\xe3\x22\x66\xf3\x63\x29\xa0\xd8\x4d\xc0\x0b\x5f\x68\xed\xb4\x06\x43\x4d\x09\xe4\x42\xa0\xa3\xc6\x53\xb0\xa0\xbe\x1f\xaf\x09\x0f\xf5\xf7\x48\xda\xa0\xad\x61\x85\x97\xed\x40\x54\x12\x61\x60\x04\xd7\x75\x66\x62\xc7\x88\xb2\x4c\x72\x8d\x94\x65\x4a\x92\xde\x16\xfb\xc1\x75\x3f\xee\x4e\x1c\xf4\xc7\x77\xdd\x4d\xfe\x6c\x0f\x47\xa7\x9c\x89\xdc\x11\x2e\xb4\xe4\xb6\x0c\x0e\xf2\x9c\x78\x45\x48\xec\x5c\x1c\xf4\xd3\x55\xb9\xe5\x5c\xf3\xc1\x9b\x2f\x1e\x2e\xea\x40\x2d\xcd\x2f\x56\xc8\x72\x2b\xd5\x5a\x1d\x99\x0d\x43\x32\x77\xcf\x05\xb5\xfa\x40\xbe\x26\xee\x08\xb1\x46\xcd\x5c\xbb\x47\xbe\x64\x7d\x76\x80\xef\xcf\x55\x93\xbb\x0f\x14\xaf\xc9\x0f\x98\xdd\x75\x7b\x28\xc0\x6f\x33\x93\x9f\x70\x98\x2d\x9e\x9a\xe1\xa3\xd9\x79\xcd\x7f\xc1\xc3\x84\x45\xf1\xc2\x08\xcf\x09\x1b\xa6\x4b\x7b\xd7\xf9\x29\xdc\x49\x1d\xc8\xec\x1b\x46\x0b\x14\x20\xcd\xb2\x88\x68\xe3\x4f\x8c\xda\xd6\x87\xd9\xe7\x53\x59\x1b\x3f\x14\x62\x9c\xd8\xa0\xc6\x36\xa8\xd6\x0f\x5b\xbc\xd8\x76\x2e\x1f\x9f\xf6\x4b\xf1\x3c\x1d\xb3\x0c\x59\x15\xc8\xcf\x8f\xc4\x3e\xdb\x3b\x7f\x5a\x4e\xa7\x68\x51\xdf\x28\xab\xc7\x71\x20\x51\x9d\x3d\xda\x94\xdb\x54\xc0\xeb\x77\x75\x8b\xa2\xf3\xf3\x67\x85\x95\x1c\x0e\xd0\x23\x5d\x5d\x33\x89\x34\x8a\x6d\x0b\x46\x82\xbb\x3f\x32\x15\x15\xf7\x5b\x0f\x00\x00\xb0\x47\x5c\xf5\x8c\x02\x00\x54\x65\xfe\x69\x86\xb8\x6a\xe3\x18\x00\x00\xc0\x86\xb4\xb1\xb3\x64\x5b\x60\xe1\x38\xcb\xaa\x55\x67\x5b\x51\x0d\x00\x00\xce\x23\x2d\x3d\x90\xe6\x88\xf0\x70\x08\x1c\x09\x30\x8f\x84\x04\x20\x21\xb7\x24\x63\x61\xc8\x60\x49\x6b\x1b\x3b\x07\xf7\x74\x6e\x21\x00\xe0\x5c\x8b\x8d\x85\xa9\x4b\xdc\xcc\x3a\x21\x11\xe9\x32\xf5\x62\x64\x7b\xf3\x42\x3a\x40\x56\xd2\xc6\x76\x11\x22\x19\xdc\xe4\xd8\x24\xd9\x2f\xcf\x61\xc1\x92\x61\x66\xeb\x04\xba\x51\x2d\x14\x06\xc3\x52\x6b\x7b\x0a\x14\xc1\xbd\x11\xae\xa1\x28\x0f\x93\xf7\xad\x2f\x46\xa5\xdd\xdc\xeb\x9d\x9c\x04\x54\x12\x2c\xe1\xd1\xae\xb6\x8d\x0e\x82\xaf\xaa\x2c\x32\x0c\x98\xcc\xfb\x00\xe9\xdd\xe3\xf4\xc0\xec\x51\x73\xc6\x20\xcd\x20\xe8\xf8\xd8\xb1\x51\xe8\xd7\xc8\x62\x31\xe8\xf8\x5a\x54\xca\x34\xad\xac\x67\x62\x4e\x52\xfa\xff\x5b\x3b\xe1\x15\x45\x27\x07\xd5\xef\xec\x99\x6e\xf7\x79\xf7\x05\xf5\xe9\x74\xf7\x22\x90\x5e\x39\x19\x20\xb0\x29\x86\xcf\x23\xf9\xbb\xbd\x68\x1a\x25\xdd\xb5\xcf\x38\x5d\x2d\x9d\x10\x97\x6c\x8d\x44\xe7\x64\x80\x8a\x4d\x37\xcf\xfc\xe2\xda\xd2\x99\x17\x9f\x37\x9c\x67\xf2\xe7\x68\xe2\x41\xa1\x2c\xdf\x3e\x96\x27\x6d\xf0\x24\x2a\x44\x98\x03\x36\xb7\x40\x5b\x3c\x55\x62\x87\x62\x35\x8e\xeb\x6a\x40\x34\x2d\x23\x85\x7d\x0b\x91\x4d\xdc\x3f\xbb\x15\x13\xca\xe1\x0f\x26\xef\x9f\xfa\x76\xf7\x5a\x95\x47\xbe\xac\xca\x96\x23\x81\xfc\x25\xb0\xda\xcc\x95\x07\x9c\xea\xa2\x28\x84\x9f\x70\x6f\x90\xc1\x16\xb1\xcf\xca\x1f\xe4\x4b\x5c\xd6\xa3\x07\x5a\xfb\x8c\x3d\x4f\xc3\x05\x31\xf0\x64\x92\x7e\xc0\x19\x42\x14\xec\xca\xed\x60\x21\x37\x73\x30\x2b\xe2\x7d\xf2\xfe\xb2\x5b\x0a\xd5\xef\x20\xcf\xde\xe4\xb9\x30\xbb\xbc\x1a\x71\x6a\x98\x13\xff\x25\xfa\xc7\x3e\x9c\x11\xce\x76\x8e\x85\x9e\xe5\x0b\x0b\x91\x98\x77\x96\x3c\x12\xc0\xa8\xb0\x81\xfa\xc9\xde\x8d\xfc\xa6\xda\x09\x1f\x70\xab\xad\x7d\x4a\x52\x27\x93\x77\x7f\xbe\x39\x8e\xb9\xe6\x6c\xcd\x7c\xd4\x52\x11\x26\xb3\x00\xd0\xed\x8b\x8d\xe6\xab\xb1\xca\x2f\xf0\xa2\x9b\x8d\xd9\xb5\x44\xc4\xb3\x75\x30\xa0\x10\xed\x03\xc4\xba\xb3\x1e\x3f\x87\x11\x2a\xf3\xd0\xee\x5e\x4b\x3d\x06\x81\x3d\x0f\x20\xcd\x6a\x3e\x43\x7d\x73\xe8\x4e\xdf\xab\xb8\x64\x3b\xee\xa4\xca\x85\x9f\xb2\x89\x95\x3a\x6d\x45\x99\xb8\xcb\xee\x6c\xbe\xa6\xf9\xad\x55\xf9\x21\x0d\x9e\x03\xe1\x6c\xea\x20\x14\x02\x9c\x46\x0c\x02\x45\xb8\x8b\x62\x92\x99\x51\x9b\x46\xd9\x5f\x1a\x64\x1e\xb1\x96\x31\xbd\x66\x12\x62\x08\x4c\x2f\x8e\x4b\xf6\x8e\xbe\xa9\x40\x82\x73\xa0\xd9\x0e\x81\x9f\x22\x3e\x83\x54\x18\xe1\xa1\xee\xa7\x89\xb0\xd5\xea\x22\xde\xa4\xed\xb7\x87\xeb\x12\xe9\x95\x6d\x07\x57\x3e\xbc\x2d\x43\x33\x04\xa6\xd9\xf4\x01\xaa\x9c\x43\xb1\x3f\xf2\xf0\xab\x5a\xdd\xbd\x61\xe5\x2e\x40\xb4\xf7\x40\x8a\x36\xb4\xe9\xbe\x77\x12\x8b\x4f\x1e\xae\x7d\x34\x7f\xf3\x2c\x51\x64\xe7\xba\xfd\x36\x5e\xd0\x84\xc1\xd7\x6f\xf7\xfc\xdc\xdf\xfa\xbf\x21\x2b\x93\x3f\x29\x25\xcc\x9c\x6d\x9d\x42\x97\xff\xee\x17\x90\xee\x3a\x61\x5d\xeb\xf0\x38\x80\x4f\x2e\xae\x7d\x78\x29\xb6\x83\xa3\x28\xf6\x1e\xf1\x58\x4d\x1d\x88\xda\x9c\x87\x9d\xac\xcf\xfe\x9c\xa1\x06\x63\xbb\x5a\x43\xf3\x7b\x8c\xca\xbd\x81\x68\xf2\x40\x0a\x17\x54\xa0\x18\xbf\xd7\x47\xc6\xc7\x67\x25\x23\xea\x26\xba\xd7\x0b\x4d\x61\x21\x42\xf3\x22\xf3\x88\x22\x38\xc6\xc4\x30\x07\x47\x69\xba\x71\x81\xe2\x60\x4c\x7d\x17\xfa\xaa\x3d\x5d\x5b\x75\xb6\x05\xa2\xd3\xa6\xeb\x5a\x79\x64\xe4\xef\x23\x2a\x25\xc1\x95\xc4\xac\x1c\x2b\xbf\x60\x9d\x86\x59\xbd\xe0\xf4\x10\xa4\xd0\x9a\xc0\xda\x34\x9d\x99\x58\x98\x8d\xa3\x40\x45\x69\xb3\xec\x35\xc6\x2b\x6f\x55\xfa\xf6\xe3\x92\x95\x6e\x6c\xd3\x4f\xce\x2c\xe8\x76\xa7\x86\x6e\x53\x55\xa7\xd8\x5f\x8b\x18\x3a\x85\xec\x87\xf8\xf7\x72\x50\xc2\x56\x60\xc3\x9f\x2b\x57\x13\x12\x4e\x14\xcc\x7a\xcf\x1e\x5e\x39\x23\xc5\x76\x58\x7d\x03\xc4\x62\xf8\x7d\xf6\x7e\xfc\xc2\xa4\xdd\x93\xb4\x11\xfa\x19\xaa\x84\x7d\x47\x8a\x9a\x1d\x98\x67\xd6\x2a\xa2\xe2\x27\xed\xe5\x7e\xb3\xb5\xc0\x5e\x4e\xed\x54\x3e\xc0\xd9\xa7\x23\x28\xd6\xa1\x64\x66\xc2\x0a\x3b\xba\x4d\xd5\x69\xcb\x3a\x16\x19\x93\xa9\x15\x62\xb9\x8d\x81\x7f\xb7\x79\x55\x65\x1d\xf2\xcb\x1d\xa0\xd6\x15\x31\x51\xd8\x1e\xc3\x52\xa1\x4a\x52\x6b\xcb\xea\x96\x78\x23\xf4\xc2\xf9\x2c\x25\x48\x7c\xc9\xe6\x8a\x23\x56\xe6\xe7\x25\xff\xfc\xd1\x83\x95\x9b\xfe\x3b\xb0\x8c\x39\xce\x6e\x0b\x58\x08\xe3\x1b\xbc\x32\x29\x95\x4f\xfc\x7c\xa5\x1a\x30\x41\x38\x74\x7f\xa0\xcf\x41\x68\x63\x9a\xe3\xd9\x1b\xd9\x56\xde\xbd\x76\x6f\xda\xec\x6e\xfa\x26\x4b\xbc\x71\x05\xf5\xf4\x3a\xb3\x23\x56\xc6\x60\x6d\xaa\x8f\xf9\xf9\xba\x2b\xe3\x16\x4d\x27\xe3\xc4\x8c\x9c\x7a\x07\x5c\xba\x29\x8b\xd4\xda\xcb\xc9\x00\x35\x93\xf7\x4b\x7d\xda\xdc\x4f\x4d\x61\xab\x43\x72\x2c\x74\xf5\xad\xd8\x6b\xce\x44\x60\xff\xf2\xc4\x47\x32\xbe\xee\xe9\x00\x81\x11\xa4\x1e\xe8\xb9\x9e\x08\xb8\xa7\xe2\xed\xff\x43\xf6\x57\x07\xaf\x4e\x58\xda\x97\x6b\x1e\xed\x8d\xfa\xda\xda\x8d\xc6\x30\xa3\xc3\xf8\xa0\x48\x35\xf2\x9f\x12\x7a\xb1\x2f\xf1\x63\x25\xa2\x0e\x72\x64\x68\x23\xc0\x7d\x75\x61\x1a\x25\xfa\xd9\x2f\x3a\x67\x6c\x85\x50\xc3\xf1\x67\xf9\x64\x62\x26\xc1\xea\x04\x27\x55\x4d\xba\xb5\xb2\xa2\xc0\x09\xf2\xb9\x1e\xa3\xc3\x29\x0c\xf0\xf4\xc4\x62\x8c\xcb\x62\x7d\x69\x66\xab\xdb\x16\x1e\x58\xca\x5a\xfa\x3d\xa0\xbc\xd2\x9d\x27\x7c\x4a\x4f\xf2\x21\xad\xed\x4c\xc9\x1c\x5d\x68\x68\xaa\xe2\x9c\xf1\xe0\x70\x14\x7f\xd1\x0c\x4a\x46\xd4\x0d\xe4\xdf\xbf\xa5\x7a\x6c\xa4\x10\xc7\xd5\x26\x17\x52\xd2\x41\x7e\x41\x1f\x7c\xa7\xda\x8e\x58\xb2\xe5\xc9\xcc\x7c\xf4\x94\x09\x06\x73\x14\xde\x9b\xf1\xf2\xe8\x11\xd9\xe8\x15\xad\xad\xad\x0f\xd6\x85\x47\xbb\x0f\x6e\xd7\x49\xc4\x06\x06\x3c\xd1\x6f\x4c\x3b\xb7\x11\x7c\xcb\xd2\x8b\x71\xa8\x54\xf4\xf0\x84\x25\xd7\xf4\xb8\x4c\x16\x9e\x1c\x6c\x5c\x41\xff\x02\xef\xba\x05\x4d\x28\xd2\x17\xb2\x7a\x98\xb1\x41\x78\x75\x81\x24\x28\xd8\x14\xcc\xab\xec\x75\xcb\xf2\xaf\x46\x0b\xba\x4f\xee\x72\xf4\x9f\x03\x25\xf4\x16\xdc\x2c\x4e\x42\xbc\xa3\xb7\xff\xe6\xf7\x3c\x2b\xdf\xd5\xe9\x74\x41\xf1\x8c\x16\xb5\x8b\x78\x11\x85\x93\xa1\xcb\x58\x06\x6c\xa4\x40\xba\x82\xea\x81\x83\x4f\x58\xc8\x61\x0e\x57\x42\xb2\x61\xbd\x0f\xbb\xa5\xa5\x9b\xb2\x58\x4b\xa1\xc0\x2f\x64\xb2\x75\x45\x4f\x37\x47\x7e\xf4\x7d\x1e\x7c\xca\x6c\x0b\xa6\x98\xc8\xfa\x7b\x17\x15\xfe\x21\x0e\x6a\x29\xbe\x5c\xa0\x48\xc8\xf3\x6e\x1b\x33\xf7\x05\x21\x37\xc4\x02\x4c\x51\x0d\x99\x38\xdf\x3b\xb0\x1f\xf5\x1e\x06\xf0\x6c\xb6\xa1\xed\xfd\x47\xcc\x27\x05\x84\xdc\x6c\x1a\xb2\x4a\xd1\xac\x92\xb7\xfb\x5d\x6f\x66\xef\xfe\x39\x67\xe3\x2b\x83\x47\x13\xda\xf2\x8f\xfa\x3f\xdc\xca\xcd\x66\xe7\xbf\x31\xfd\x2c\x36\x08\x4b\x26\xff\x2e\x70\xbf\xa9\xb3\x40\xfc\xce\x20\xa2\x69\x38\xb9\xff\x70\x00\xa7\x95\x4e\x80\x6d\x27\x95\xf3\xd9\x47\x47\x8a\xe1\xf5\x87\xde\x2b\x7c\xfb\xeb\xae\xa5\x4f\xc8\xc8\x48\x68\x99\x33\xcf\x97\x39\x07\x2e\xb1\xd3\xa8\x9c\x62\xb6\x7f\x2b\x77\x96\xc7\xe7\x57\xf0\xed\x1b\x4f\x87\x86\x2d\x78\xbc\xff\x1a\x86\x19\x0c\xe0\xf3\x83\x9b\x3f\xbd\x6b\x17\xfe\x13\x2a\x1d\x15\xcb\xbf\xbc\x11\xb6\x54\x4f\x57\x41\xd2\x6a\x8b\x07\xed\xa3\x85\xcd\x45\x0e\x6c\x9e\xf6\xba\xb6\x74\x69\x91\x20\x10\x8d\x1f\x3e\xb1\x54\x1e\xb3\xeb\x9a\x72\xb6\x8b\xc2\xb1\x3b\xc0\xf0\xa0\x60\x6d\x9c\x2f\xbf\x3d\xa5\xbd\xd4\xb5\xf0\xf1\xdc\x03\x35\x65\x64\x3c\x56\x49\xd1\xff\x97\x7a\xa6\xcf\x85\xb3\x8b\x8b\x15\x7f\x11\x58\x5d\x5e\xe4\x32\x35\x33\x24\xf4\x2b\x21\x9a\x97\x7d\xec\xa5\xa5\xac\x3d\x77\x64\xa3\x95\x5f\x1b\x67\xef\x69\xf2\xd4\x92\xf4\x8d\x43\x3c\xd1\x40\x40\xb0\xb3\x7e\xa5\x37\x88\xc7\x27\xc6\xd3\x85\x2b\x8b\x97\xbd\x6e\xcc\x96\x47\x22\x09\x00\x83\xa9\x86\xa0\xdd\x5d\xe4\x17\xa4\x6f\xd4\xd6\x36\xd7\xd4\x44\xac\xad\x97\xb1\x27\x78\x16\x73\x3f\xc8\xf9\x21\x36\xe6\x3a\x6f\x47\x59\x0d\xf9\xf4\xc9\x21\x5c\x95\x8f\x73\x6a\x0a\x64\x6f\x35\x7f\x9f\x95\xc5\xf7\x15\x62\xe6\xf5\x1d\x29\x3e\x65\xee\x45\xa5\xad\x2a\xc4\x24\xa7\x99\xb4\xe9\x5f\x87\x5e\x46\x1e\x6a\x16\x97\x9e\xb0\x74\xff\x9f\x7f\xce\x09\x90\x6a\xee\xac\xab\x4d\x57\x9e\xa1\xfa\xee\xf4\xc2\xc1\xb1\xe2\x5c\x8f\x25\x6f\xcb\xfb\xa5\x13\xa3\xd9\x08\xb7\x85\xd8\xa5\xf3\x99\x74\xa7\x94\x9b\xa4\x35\x4c\x4c\x85\x1c\xb1\x96\x16\xa6\x42\xbc\x2c\x60\x0b\x01\x24\xef\x7f\x86\x3c\x1b\x07\xde\x2b\x92\xc6\x12\x56\xf9\x3a\x4a\xd0\x83\x35\x76\x05\x85\x47\x65\x9c\x60\xb4\x89\x19\xe3\xe0\x88\x79\x94\xd6\x25\xed\xfe\xc7\x32\x11\xf3\x80\x35\x22\x8f\x7e\xac\x99\xab\x60\xf7\xa9\x91\x01\xad\xfe\x8b\x63\x86\x95\xb9\x3f\xfe\x55\xcd\x99\x46\x0c\x48\xb5\x93\x99\x7a\xd0\xda\xea\x30\x77\x7e\xea\x67\x89\x35\x1f\x31\x89\x1f\x14\x67\x12\xf4\x46\x4c\xe2\xfb\x16\x6a\x3f\x30\x5d\x90\x1f\xe5\x0d\x7f\x5f\x8f\xf3\x6d\xa9\x6f\xf0\x56\xa9\xab\x1e\x02\x2a\x56\x4e\x34\xfb\x26\x7c\x6e\x6e\x6e\x19\xa0\x95\x6f\x62\x88\x91\x0d\xae\x75\xf5\x1e\xba\x94\x49\x24\x8f\xe5\x57\xc3\xf0\x92\xe2\xdf\x6d\xd6\xf9\xd8\x64\xf8\x35\xa0\x6d\x31\x23\xce\x97\x43\x0d\xc3\xa9\x05\x25\x79\xe4\xc0\x87\x75\x15\x75\xcf\xa7\x50\xf4\x0f\xf7\x9c\x5b\xb2\x7c\x02\x02\x82\xcd\x53\x76\x7a\x8f\x0f\x47\xb6\x16\x9f\xab\xc7\x37\xe7\x3d\x94\x27\xe5\x75\xae\xeb\x1c\x1e\x4d\x48\xe9\x8b\xfe\xf1\x71\x32\xa0\xd3\x4a\x4b\xa3\x31\x27\xad\x79\xf9\xf2\x24\x47\x15\x05\xf6\x8d\x4e\xb6\x8a\x84\x6b\x10\x29\x60\x82\xcb\xf5\xab\x69\xde\x88\xcc\xf8\x5d\x97\x73\x7f\x88\xef\x5e\x49\x47\x7c\x7a\x2c\xd3\x94\xac\x9f\xb4\x2f\x18\x36\xcd\x58\x7a\x79\x40\xc8\xaf\xa8\x77\xe7\x6b\x69\x40\x1f\x80\x84\xe3\x19\x37\x14\x62\xe5\xbe\xcb\xa1\x76\x09\x1c\x55\xae\x77\x74\xf4\x56\x97\x3b\x85\xea\x5c\xd3\x98\x75\xb4\xeb\xba\xf1\xb2\x8a\x0d\x1c\xae\x92\xb4\xc3\xde\x4d\x53\x35\x38\xba\x30\x39\xf9\xa1\x27\xdc\x75\xaa\x3d\xb5\xb0\x3c\x9c\x3a\x36\x92\xb7\x8f\x52\x33\xd2\xe0\x85\x4e\x64\xe1\xfc\x70\x48\xfc\xcb\x2a\x36\xcb\x40\x99\xa4\x63\x8d\xee\x6d\xd6\xd9\xd4\x54\xff\x4d\x57\x4f\x76\x5f\xb5\xd7\x6f\x4b\x4b\xc3\x71\xbe\xa1\xa7\xac\xfe\x09\x7a\xca\x3e\xc1\x27\x57\xef\xfe\xa8\x30\x8c\x18\xd7\xd5\x53\x43\x3e\x04\x5d\xe4\x16\x53\xff\xad\x2a\x9f\xf6\xc2\x9c\x89\xe7\x68\x3b\x24\xc4\xff\x23\x73\x94\x3b\xbd\x8f\x81\x76\x0a\x9f\x94\x6f\x31\x96\x06\x26\xa4\x71\x83\x2e\xdd\x27\x70\x49\x70\xe3\xc0\x41\x1f\x41\xca\x3c\x46\x67\x5a\xa0\xc5\x13\x1f\x5e\x99\xb3\x75\x90\x55\xc5\x65\x1c\xb1\xf8\x86\x17\xfd\x3e\x4c\x7e\xd1\x05\xe7\xa5\x28\xdb\x23\x2f\x8e\x66\x83\x2a\xd2\xe5\x2a\xd2\x0a\xcf\x25\xe2\x79\xc4\x81\x3d\x89\xe4\xb8\xdc\x8c\x8d\xe7\x4b\x50\x8f\xae\x73\xf0\x6e\xde\xe1\xfc\x91\xcf\xda\xb5\xf4\xda\x69\x62\x32\x6e\xae\xcd\xc6\xa6\x76\xca\xd2\x94\x82\xff\x9b\x3f\x50\xfd\xfc\xe3\xf7\xea\xdf\x84\xfc\x6c\x50\xc6\x40\xf6\x3a\x6e\x55\x1a\x98\xc0\xec\xd5\x23\x62\x74\x8c\x9d\x6b\xa3\xd3\xea\xba\x2d\x4c\xf3\x3d\xea\xdc\xa6\xf1\xaa\x2a\x76\x90\x8b\xc0\x84\x41\xaf\xee\x5f\x63\xd7\xa2\x64\xcb\xf5\x8a\x8d\xe2\x78\x59\x0c\x19\x5d\x5d\xfb\x7b\x93\x0f\xda\xd6\xf1\xee\x8b\xff\xac\x6b\xeb\x5a\xda\xc1\xed\x70\x45\x25\x47\xac\x8c\x5b\x54\x4b\xcf\x3a\xdf\xa5\x12\x4f\x81\xbd\x4f\x05\x45\xc1\x88\x2e\x3a\xcd\xd5\xed\x65\x95\x75\xfd\xeb\x0b\x89\xf8\x70\x62\x63\x33\x1f\x8b\xa1\xf4\x43\x61\x86\xa0\xca\xe4\xa5\x75\x11\xd1\x01\xe3\x3e\x6b\x01\x24\xec\x6d\x2e\x7f\x7f\x60\xe0\x35\x44\xa5\xfd\xbf\xef\x89\x18\xcd\xce\xc5\x94\x5e\xae\xab\x5f\xf8\x80\xac\xfb\xf9\x55\x0a\x73\x78\x33\x11\x9f\xb5\xfa\x19\xfd\x43\x41\x3e\xed\xf6\xdf\xb1\x67\x9e\x86\x89\x25\xe7\xa3\x9e\x3c\xfd\xfa\xea\x83\xf5\xf3\xd7\x21\x61\x97\x50\x68\xbd\xa8\xdd\x08\x0a\x5c\x8b\xe0\x66\xc9\xf2\xec\x19\xa9\xa8\xcb\x73\xac\x7e\x92\x6f\x15\xf7\xbd\xfa\xc0\x4d\x19\x16\xdc\x54\x3f\x3b\xbb\xd5\x30\x19\x79\xd8\x04\x09\xc7\x9f\x3b\x5c\xd5\xa6\x51\xbc\xdd\xe2\x3b\x21\x2a\x1e\xfe\x9e\x2b\x2b\x7a\x27\xdc\xa8\xc9\x97\x6f\xba\x44\x17\x97\x3e\x16\x69\x1c\xb4\xd1\xd7\xb4\x3c\x5a\x64\x93\x68\x25\x45\xc1\x21\x11\xf6\x69\xa3\x77\xa2\xd7\x81\xba\xba\x11\xcb\xe5\x6f\x0a\xfd\x3a\xbf\xa7\x9c\xcf\xc9\x00\x3d\x06\xa5\xfc\xba\x94\xb8\x1e\xb7\x52\x4d\xa2\x9d\x80\xbb\x4a\x3a\xb7\x0a\xd5\x2e\x3f\x4d\xe5\x14\x79\x1d\x65\x66\x9a\xaf\x19\x3b\x1b\x00\x6d\x7d\xff\xaa\x16\x08\x4a\xbe\x94\x30\xfa\x87\x34\xa0\x87\xc0\xac\xaf\x89\x25\x03\xd1\x77\x5d\x13\xd0\x9e\x0f\x87\xc6\x23\xdc\x46\x33\x8c\xe9\x0a\x3a\x5f\x2e\xaa\x7c\xb5\x1d\x7c\x1f\x56\x3c\xce\xcb\x62\x08\xec\xde\x62\xdc\x5b\x85\x20\xf0\x3c\xc6\x77\x89\xc7\xb2\x59\xe3\x93\x57\x5a\xe6\x21\xf1\x9d\x47\x77\xe4\x81\x09\x16\x25\xfa\x42\x8d\x9d\xc3\x73\x6d\xd9\xdc\xc5\xe4\x43\x8f\x67\x4f\x03\x1b\xf1\xeb\x73\x03\x67\x4d\xe8\xa4\x0d\xf5\x64\xb4\x27\xf4\xf3\x46\x39\xf9\xe2\xf5\xf9\x29\x77\xc2\xfb\x4d\xda\x41\x7f\x9e\xf8\xd6\x3e\x6e\x09\x7e\xd4\x08\x86\x3a\x66\xaf\xe3\xbf\x2a\xea\xd8\xd8\x40\xa7\x57\xea\x95\x3f\x7d\x4a\xdb\x9d\xb8\x3c\x31\x4e\xd9\x44\xa4\x8a\x9d\x7e\x56\xed\x11\x70\xab\x4f\x7b\x2c\x13\xbd\x71\xf7\x1e\x55\x97\x29\x89\x4b\x02\xee\xd9\x15\x33\xbd\x89\x38\xa1\x60\xcc\xdb\x11\x6f\x57\xad\x11\xc1\x4d\xf5\xfb\x64\x15\x72\xe7\xb2\x72\x5d\x95\xe4\x16\x45\x4d\xad\xbd\x7e\x78\xf8\x60\xbf\xe7\xf9\xb3\x1a\x09\xb0\x69\x7e\x9a\x7b\x64\x33\x24\xf9\xa0\x98\xce\x8f\x5f\x6e\xa9\xe7\x95\xf9\xc8\xbf\xd7\xe0\xad\x42\x3e\x5a\xbe\xa3\x75\x25\xe8\x8b\x00\x7f\x5c\xaa\xb8\xee\x4d\xcf\xcb\x73\x35\x55\x4a\x24\xd5\xd4\x1e\x5a\x23\xfd\x29\xfa\xfd\xb7\xe6\x94\xed\xfb\xd8\x0d\xe2\xda\x74\xa1\xab\x27\x8c\x51\xe2\xde\x44\xe2\x02\xef\x68\x6d\xdd\xae\xda\x1a\x19\xa7\x44\x02\xa2\xef\xfb\x41\x20\x27\xe5\xd8\x09\x43\x21\xd1\x45\x4d\xc8\x3d\xc8\xfa\xec\x8a\xdf\xf4\x58\xf1\xab\x7f\x44\x92\x0c\x29\x65\xa9\x72\x3b\x64\xd4\x88\xdb\xd1\x03\x09\x84\xa7\xb6\xf6\xc1\xe8\xf2\xee\x84\x20\x92\x85\x00\x41\x65\x72\xa8\xd4\x75\x4e\x9f\xd8\x1c\xcf\x25\xbe\xf7\xbc\x31\x34\x54\xe8\xea\x87\xc0\x94\xc1\x3b\x2d\x4d\xdd\xb1\xe0\x9d\x84\x35\x8d\xd3\x0d\x75\x14\x38\x14\xfe\x6b\x28\x24\x9e\xbe\xa8\x76\x30\xba\x2a\x71\xea\x80\x15\x44\x0a\x75\x83\xe2\xfa\x66\x2f\xb6\x91\x7f\xd8\xab\xd4\xc5\x22\x0a\xf0\x07\xeb\x07\x5f\x42\x71\xaa\x2c\x60\x99\xbf\x1f\xc3\x7e\x2a\x17\x19\x8a\xea\xc5\x7c\x68\xa7\xbf\xb1\x83\xbe\x5b\x6a\xf0\x76\x8d\x19\x7d\x8e\xc7\xab\x1a\x82\xeb\xa6\x6a\xac\x78\xfe\x92\x0f\xf0\x2a\xe4\x79\x53\x9a\x76\x7c\x74\x2b\x64\x71\xa1\x66\x36\xaf\x57\x53\x4b\x85\x8c\x95\x4a\x84\x50\x26\x9d\x6b\xbf\xb9\xbe\xca\xee\x7f\x70\x7e\x49\xe3\xbb\x2a\x78\x2d\x1c\x13\x98\x7d\x9e\x5b\x34\x4f\x7b\x9a\x50\x5b\x0b\xd4\x09\x6b\x6a\x0f\x77\x2b\xfc\x72\xdd\x94\xe7\xbd\xc4\x09\xd6\xa0\xe6\x26\xe5\xf7\x7d\x42\xd7\x8b\x8e\x5f\xca\xcf\x95\xe5\xa8\xc7\x51\x89\xf8\xf2\xa8\xbc\x2f\xd7\x4c\x79\x5e\xe4\xf2\x07\xf4\x33\x26\x46\x35\x33\x4d\xfe\xfc\x18\xde\x5e\xcc\xa6\xbd\x80\x5f\xfb\x49\x22\x74\xd5\xd5\x0b\xec\x9d\xfd\x3c\x91\x6d\x7c\xcc\xe0\xe4\xb4\x16\x8e\x49\x35\x3e\x5e\x0b\xda\xda\x19\x55\xec\x6a\x4b\xda\xc9\xed\x59\xd7\x8b\x75\xf3\x54\x20\x49\xa9\x43\x39\x95\xbd\x6e\x61\xd8\xe3\xe7\x01\xc7\x3b\xcf\xb7\x8e\x77\x15\xbb\x7e\x1c\xfd\xe1\x48\xc4\x4f\xaf\x4f\xe5\x98\x56\x98\x10\xfb\xb4\xc5\xa9\xbf\x75\x6e\x32\x3f\x2b\x2c\x39\xef\xb8\x5e\x83\xca\x4e\xdc\xcc\x94\xab\x74\x99\x9e\xd6\xfb\x16\x56\x06\x25\x26\xf1\x17\x8c\x5a\x27\x8a\x7d\x7a\x17\x82\x94\x5b\x58\x2a\x81\x55\xad\x65\x1d\x6f\x05\x14\x69\x59\x4b\xb5\x8f\x6e\x28\x37\xc5\x3c\x2f\x72\xfb\x01\x9e\x89\xe6\xeb\x3d\x94\xde\x5b\x09\x44\xd0\xb2\x58\xe3\xd6\x34\x0e\xd7\x11\xae\x31\x31\xe0\x88\x74\x7a\xb3\x20\x0b\xe1\xb6\xa7\x27\x54\xf0\x47\x0a\x7d\x2d\x9c\x38\xc9\x6c\x6f\x10\x10\x12\xf5\x1e\x31\x6f\x42\x53\xb9\xe8\x88\x15\x9f\xea\xba\xfc\xac\xa0\x58\x17\xbb\xa6\xa3\x54\x99\x00\x4c\xd9\x57\x2a\x28\xc7\xb0\x72\x36\xb5\x13\xe6\x7a\x14\x95\x58\x08\xa6\x2b\x2b\x50\xd9\xef\x9b\x29\xb8\xb9\x03\x15\x15\x28\x6f\xff\xe9\xb1\x7d\x01\x38\xd1\x27\xee\x45\xcf\xfa\xc1\x9a\xa3\xa7\x02\x89\xdb\xe6\x46\xd6\xe5\xb9\x26\x62\xd9\x07\xe5\xc3\xdd\xca\xdf\x3f\xbd\x5d\x63\xd6\x1f\x17\xc3\x30\x56\xc4\x0f\x36\x8b\xd5\x55\xfa\x0b\x0d\xed\xfe\x9f\xce\xd2\xf0\x98\x2d\xc6\xe3\x05\x39\xc5\xf6\x83\x12\xed\xd1\x13\x93\xa2\xcb\x44\x97\xf8\x56\x53\x3a\xe9\xe4\x70\xd9\xd8\x2e\xb5\x13\x8c\xea\x66\x7c\x2a\x47\x56\xe6\x18\x6c\x0b\x3f\x40\xab\x24\x1b\xb5\x20\x9d\x91\x59\x2e\x74\x92\x54\xc3\x67\x3a\xfb\x51\x1d\x4a\x71\x21\x3d\xe1\x24\x9e\x16\xc8\x6f\x44\xbd\xba\x58\x5d\xc5\x46\x9a\x72\x0f\xd8\x0b\xa6\xe1\x51\x3e\xec\x38\x3f\x54\x24\xc7\xe0\x83\xf3\x4b\xbf\x25\xbb\xf6\x4c\x1c\x22\x2c\xe9\x83\x6c\x6a\x98\x35\xe2\x83\x86\x3a\x37\x58\x73\x8e\x89\x5f\x7c\xfc\x49\x83\xef\xee\x94\x7b\x09\x8b\x9b\x76\x42\x9a\xdd\x88\xcd\xd3\x35\xc4\x29\x2a\xb5\x7f\xca\xdd\xfa\x20\x36\xe5\xf8\xb6\x47\xcb\x30\x66\x66\xd2\x86\x5f\xee\xe3\x9f\xa1\x93\x86\x22\x34\xf2\xdd\xb8\xeb\x14\x69\x3d\xa9\xe0\x99\x87\xdb\xf5\xfb\x74\xaf\xea\x6e\x03\xbd\x5e\x4a\x06\xcb\xc8\x4f\x8f\xc4\x4d\xc3\xd1\x12\x23\x43\x64\xa5\x61\x50\x1e\x24\x79\xf7\x61\xaa\x4a\x8d\x1a\x66\xf2\x9a\x0a\x30\x41\xdf\x6f\x77\xd9\xdb\xcd\xf9\x29\xfd\x69\x38\x01\x4d\x2a\x2e\x3c\x0a\xc6\xe6\x87\xc0\x75\x46\x35\x55\x7c\x84\xbb\x3d\x61\xf7\xef\x29\x93\xe9\x14\x3a\x6d\x20\x01\x59\x5c\xfb\xe7\x26\x0b\xd8\x14\x31\x42\xfe\x66\x4d\x2f\x4d\x4c\xa6\x27\x8c\x5e\xb3\xcf\xaf\xd8\x4c\x8d\xc4\xe6\xc7\x6e\x55\x15\x87\xf6\xbe\x1e\x6c\xde\xd9\x65\x3a\xad\xa7\x6d\xb4\xd4\x6b\x87\xda\x92\x5f\x18\xfa\x26\x39\xd6\xbd\x9b\xd9\x7c\xc9\x61\x40\x72\x8c\x56\x71\xf2\x84\xa5\x06\xd7\x95\xd6\x83\x43\x7f\x72\xf1\x4c\x66\xb0\x44\x2c\x97\x1f\xb4\x11\xba\x70\xb6\x36\x3d\x51\x72\x61\xd3\x91\x61\xe5\x7b\x2a\xde\x63\x62\x71\xde\xed\x25\xbd\xeb\xbe\x3d\x63\xfc\xc1\xc9\xb4\x92\x20\x26\xea\xbe\x16\xf9\x15\xbe\xdc\x7e\xd3\xc5\xe8\xb0\xcf\x00\x24\xc6\x54\x46\xb2\xb6\xc7\x82\x59\x61\xed\xcf\x04\x76\x9b\x69\xdd\x1d\xce\x46\xd4\xab\x84\xe7\x68\xb2\xaa\x63\xc3\x4b\xc6\xe1\xbc\xcc\x12\x91\x82\xcc\x0c\x55\xc7\xb7\x9d\xe7\x4f\xa8\xee\xc2\x3a\xb6\x52\xee\xd8\x8b\x8a\xe1\x78\x68\xe1\x07\xf4\x6f\xd6\x82\xcc\xc6\x6f\x17\x57\xd7\xf0\x41\x15\x34\xcc\xe2\xb9\x9a\x3a\xe7\x78\x79\xd2\x9f\x22\x6b\x90\x86\x7d\xc6\xe2\x52\x09\xbb\x04\x75\x2f\xfd\xee\xd9\x4a\x59\xef\xb6\x3e\xfd\xcd\xbe\x64\x0a\xec\x85\x0d\x0b\x38\xe3\xe5\x1b\x56\x50\xca\xe1\xaf\x53\x1e\x11\x66\xe7\xb7\x25\xb9\x3a\x9a\x9a\x55\xd6\x3c\xcb\xee\x16\xb0\x10\xcd\xdf\x8d\xf5\x97\x06\x3c\x85\xd4\x7e\xde\x8e\x28\xa2\x7e\x1b\x4a\x7d\x65\x55\xfc\xf6\x48\x53\x73\xac\x68\x6f\xd2\x85\xff\x64\xe1\x16\x09\xfb\xfb\x97\x92\x14\x83\x01\x88\x7f\xe1\x8d\xac\xd6\x75\x72\x3a\x13\xf4\x24\xba\x92\xa7\xb6\xce\xae\x63\xf4\xcb\x58\x6a\x75\xc3\xa3\x04\x01\x24\x2f\x0b\xfd\x01\x12\x07\xbe\xc0\x2a\x4f\x92\x06\xa2\x73\xe4\x49\xd2\x2f\xab\xee\xe7\x64\xa8\xfe\x5f\xc3\xff\x5c\xf1\x5a\xe5\xbf\x1b\x16\x21\x0b\x4c\x48\xc6\x1b\x28\x0a\x4d\xfd\x17\xa8\xca\x78\x2c\x2a\xb3\x96\x66\x55\xf9\x17\xae\x67\xd4\x48\x15\x9b\x22\x49\x87\xa0\xff\x17\xd3\x61\x3e\xcf\x71\xc5\xa3\xf0\x5c\xc4\x59\x76\xe8\x55\x43\x90\xca\xbc\xe4\xff\x30\x81\xe8\x1c\x8b\xdc\x3e\xeb\xe1\xf2\xa4\x4a\x13\x99\x82\x3d\x5d\xaa\x58\x90\xb1\xb5\xf6\xff\x8a\xca\x93\xa4\xf9\xb4\x19\x25\xce\xb3\x89\x9c\x97\xe3\x9b\x7d\x14\x72\xab\xec\x3f\x8b\xff\xe7\x03\xd7\x4d\x9a\x9b\xfd\xbf\x08\x14\x98\x73\x7e\xa3\x92\xde\x07\x1f\xc7\x09\xeb\xf2\x35\x76\xf3\x29\x2a\xfa\x6f\xdf\x2c\xc8\xcd\x50\x75\x5c\x9c\x09\x0d\x18\x3c\x23\x7c\x99\x9b\x1b\x74\xd7\x28\x39\x8f\xc5\x50\xf7\xfc\xea\xaa\x8b\x55\x08\xfa\xba\x45\x17\x35\xa4\x63\xdb\x4a\x58\x93\x59\x2f\x3a\x63\xd2\xa5\x83\x24\x5d\x64\x2c\xd1\xdc\x22\x18\x39\x45\xe2\xdb\xce\xe2\xd4\xf8\xba\xff\x58\xe6\x77\xf7\x6e\x1e\xbd\xdb\x5d\xed\xd3\x44\x76\x38\xbe\xaa\xd9\x69\xf1\xb5\x08\x58\xda\x2d\x52\x53\xe0\x9c\x99\x23\x6f\x59\xe2\x46\x2e\xe6\xba\xed\x5d\x90\x61\x55\x15\x60\xa3\x87\x32\x98\x2a\x03\x44\x5b\x38\x7c\x4a\xdd\x62\xea\x38\x8e\x14\xbe\x74\x46\x2f\xe6\xf8\x40\xa7\x5c\x7f\xe0\xe2\xf3\xbe\x22\x0d\x6c\x57\xfb\x51\x71\xc9\x36\x10\x6d\xf1\xe8\x7d\xe0\xef\x2c\x51\x7f\x12\xf6\x07\xe2\xf9\xb3\x5b\x8d\x3f\x2a\x53\x44\xac\x1e\x3a\x75\x2a\x90\xf8\x34\x80\x40\x49\x35\x1e\x56\xee\xbb\x2d\x13\x46\xbb\x07\x2b\x2b\xea\x61\xc2\xb6\x45\x55\xcd\xce\xd8\x8b\xd2\xb7\x82\x4a\xdf\xc4\xa5\xf4\x76\x76\xc1\xc9\x61\x39\x77\x9e\x09\xf0\xef\x65\x72\x44\xf7\x34\x56\xdd\xaf\x83\x40\xfa\x27\xdf\x3d\x2a\x0d\x85\xf7\x0e\x1d\xdd\x2e\xd2\x8f\xb9\x9b\x42\x9f\x6f\x1e\xf5\x1f\xe6\x04\xa3\x94\x48\x7f\xae\x7d\xfe\xbc\x24\x92\x78\x1c\x77\x38\x47\x99\xa1\x80\xc4\xf5\x4b\x8c\xcf\x5d\x26\x72\xf2\xcb\xd5\xda\x0a\x22\x6f\x97\xbf\x73\xb6\x73\xf7\xc2\xa4\x9c\x95\x20\x9e\x50\x56\x66\x04\xc7\x8d\x25\x64\xed\x34\x09\x91\x21\xc3\xd6\x97\xb8\x8a\x72\x65\xe3\xe3\xa1\x71\xe4\x1a\xc7\x37\xa3\xca\x0e\xe0\xae\x8f\xae\xee\xc5\xd5\xd9\x85\x2a\xce\x47\xb5\x56\xcf\x9b\x44\x8d\x66\x39\x11\x3d\xc7\x91\x41\x66\xdc\x6f\x2c\x1f\x15\x42\x55\x19\x96\x26\x2e\x27\xaf\xa1\x7d\xdd\xb1\x17\x6d\xb3\x38\x35\xfe\xfe\xca\xe6\xb9\x6c\x17\x23\x37\x34\x62\x4e\xec\x8d\x20\x44\x2f\x71\xf7\xca\x93\xf8\x64\x71\xe0\xa0\xc6\xea\xc4\x7d\xe0\x38\x4d\x4e\xd6\x6c\xa2\x02\x7a\x97\x50\xef\xf9\x2f\x67\xe1\xde\x7d\xbb\xa8\xe4\x79\x93\x5d\x0b\xaf\xa6\xb2\xe8\x42\xd8\x47\xbf\xe4\xd6\xac\xcf\x8a\x24\x3e\xa7\xa3\xa3\x37\xb2\x98\x3f\xf7\xe6\xeb\xa7\x92\x9b\x5f\x5f\xd0\xf2\x4d\xf9\x6d\xb6\xeb\x37\xb5\x0a\x44\xdf\x95\x5c\xe0\x97\xfd\x25\x31\xb7\xca\x9c\x3f\xc2\xd9\x77\x5e\x4f\xa0\xdb\xd7\xce\xe5\x45\xd5\x75\x07\x08\xa4\x18\x8e\x3a\x4e\x3a\x1c\x8f\x35\x4b\x1a\x1e\x2c\xd6\xdf\xb6\xb4\x54\x0f\xca\x6b\xc4\x09\x22\xd9\xef\x8d\xd9\x16\xe9\x7c\xbd\x84\x4b\xdc\xdb\x42\xc4\xd5\x39\xf3\xf1\x7f\x25\xd0\x67\x7d\xa2\xf2\x4d\x70\x64\x20\xfa\x0a\x89\xb4\xf5\xf3\x7d\xa0\x94\x9c\x5c\x37\xed\x49\x51\x30\xa8\xac\x58\x63\xc0\xc0\xe0\x95\x17\xd8\x54\x28\xdb\xb6\x22\x0c\xde\x5b\x9e\x9a\x12\x17\xb2\xf3\x34\x9a\x0c\x3d\x80\x50\x91\xbb\x57\xd7\xbe\x79\x81\xd3\x63\xc5\xb9\xd2\x1f\x0a\x58\x67\xb7\xe7\x89\xc1\xd0\xe8\xd9\x3f\x6b\x2d\x01\x7b\xc1\xc1\x21\x19\x98\x53\xf0\xd7\xb8\x29\xd2\x77\xb8\xae\xee\x92\x6a\x3b\xde\x8c\x1b\xba\xe0\x13\x52\x5b\xdc\x74\x73\xae\x9b\x5f\x4d\xeb\xc2\xbd\x8c\x8a\x8a\x8a\x97\x37\x39\x5f\x9a\xde\x19\xd3\x8e\x80\x37\x76\xf2\x2e\x7f\x89\x09\xc8\x98\x6e\xbf\xfd\xd7\x56\xcb\xb7\x15\x1d\x25\xf6\xc5\x7d\xde\x54\xc2\xa0\xc3\xcd\x5a\xe6\x93\xdd\x81\x82\xb3\xa4\xcc\x00\x81\x56\x5c\x14\x9c\x35\x7f\x42\xb2\x8e\xc0\x81\xf5\xe0\x37\xfe\x31\x62\x6c\x86\xfa\xde\x6c\x32\x1e\xd0\xf1\x4b\xb8\x45\x21\xbf\xb5\x59\x7e\x58\xff\xb3\xb2\xa4\x35\x86\xd2\x7a\x65\x71\x63\xb9\x7f\x7d\x9b\x91\x7f\x1a\xb7\x95\xba\x80\x40\xd1\xb4\xf7\xee\x0e\x63\xa9\x89\xb4\x58\x80\x84\xd1\xec\x98\xac\x52\xc7\xc5\x7f\x8c\x1c\xce\x6f\x63\x5e\x3b\x93\x9f\x6f\x8c\xb0\x86\xd6\x0b\x86\xe5\x3e\xa0\xe7\xe5\x67\x67\xb1\x12\x1e\x44\xdd\x33\x3f\x02\x1e\xda\x9c\xa9\x43\x8b\x68\x6a\x2a\x8b\xf3\x3e\x96\xac\x69\x3c\x88\x13\xeb\x30\xe6\x8f\xfb\x41\xef\xde\xf5\x5f\xb2\xcd\xcf\xcb\xc9\x19\x80\x8e\xef\x95\x35\xdd\xf0\xf5\xb8\xc3\x0d\xd5\x34\xe3\xbd\x5d\x78\xa7\x33\x52\x96\xa4\xab\xfb\xa8\xf8\x9d\x17\x78\x50\xa1\xb8\x18\x24\x57\x19\x14\x21\x4b\x82\xc1\xd4\x63\x1b\xb9\xf3\x72\x9c\x97\xc6\xb5\x3b\x23\x5d\xb1\xfd\x06\xae\xae\x3e\xf5\x5e\xe0\x0b\x0e\x43\x43\xa0\xb7\x42\x68\xce\xe3\xc6\xae\xe8\x59\x2e\xca\x3d\xa0\xb0\x8e\x6d\xec\xbc\x25\xa3\xa3\x6f\x8a\x59\x0b\xe1\x6f\xdf\xa5\x80\x71\xa3\xc9\x27\x4f\x99\x46\x0b\xdf\xb9\x09\xc1\x50\x90\x51\x83\xfe\xa2\xd0\xbe\xa3\xb7\xb8\x9d\x1d\xff\xcf\x7c\x99\xe0\x41\xbf\xeb\x85\x71\x9f\x46\x39\xbe\x78\xeb\x4b\xec\x8d\x97\x7f\x7d\x9d\x09\xec\xca\x49\x95\xaa\x54\x97\x0a\x9e\xb8\x74\x25\xf1\x61\x4a\x37\x35\xc1\xa3\x93\xef\xde\x3d\xe4\x57\xd5\xc8\x37\x4b\x09\xbe\xbe\xa5\x52\xb1\x44\x6d\xea\x67\xed\xc4\x80\x17\x9d\xef\xab\xba\xb6\xa5\x89\x6e\x3e\x6b\x2f\xba\x67\xcd\x03\x02\xb0\x95\x37\x38\x7d\x1f\xa6\x74\xff\x27\x35\x71\x19\xb8\x1e\x36\x3a\xb6\xa7\x1e\xdb\xb8\xf2\x64\x63\x8e\x3e\x26\x15\x5d\x65\xc5\x63\x87\xcd\x9f\x69\x19\x41\xf6\x17\x14\x98\xda\xe4\x32\xaa\xbd\x87\x2c\xd9\x3a\x27\x2e\x2f\x1c\x6a\x0b\xcb\xee\x64\xa4\x74\xff\x97\x78\x28\x31\xda\xea\xaa\xc5\x75\x7b\xfb\x80\xad\x79\xaf\xe6\x96\xaf\xbb\x0b\xff\xc6\x3d\xeb\xc9\x26\x43\x7f\xbf\x15\xa0\x72\xfe\x6b\x59\x54\x9e\x32\x3d\xe4\xb8\x82\x79\x69\x27\x85\x01\x00\x00\x00\x1b\xcb\x1b\x16\x8d\x66\xfe\x19\xff\x27\x00\x00\xff\xff\x41\x72\x14\x5d\xc9\x18\x00\x00")

func imgIconIconPngBytes() ([]byte, error) {
	return bindataRead(
		_imgIconIconPng,
		"img/icon/Icon.png",
	)
}

func imgIconIconPng() (*asset, error) {
	bytes, err := imgIconIconPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "img/icon/Icon.png", size: 6345, mode: os.FileMode(438), modTime: time.Unix(1589213211, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"img/icon/Icon.png": imgIconIconPng,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"img": &bintree{nil, map[string]*bintree{
		"icon": &bintree{nil, map[string]*bintree{
			"Icon.png": &bintree{imgIconIconPng, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
