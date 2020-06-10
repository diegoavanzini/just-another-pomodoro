// Code generated for package pause by go-bindata DO NOT EDIT. (@generated)
// sources:
// img/pause/pause.png
package pause

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

var _imgPausePausePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x98\xc6\xc0\xc0\xa8\xc3\xc1\xc6\xc0\xc0\xa0\xa8\xee\xe5\xca\xc0\xc0\xd8\x9a\xe9\xec\x1c\xe0\xe9\xec\xac\x50\x50\x94\x9f\x96\x99\x93\xca\xc0\xa0\x31\xb1\x76\xa2\xad\xc7\x61\x07\x99\xa3\xf1\x97\xa7\x75\x85\xc8\xbe\x51\x78\x34\x61\xa1\x55\xa4\xc0\x7e\x81\x98\x2d\x5c\xae\x2b\x02\xb2\x5a\xc3\x64\x7d\xf4\xfe\x0a\xcc\x32\x51\xd1\xd4\xeb\x6b\xca\x3e\x78\xb0\xff\xe8\xaa\xe6\xee\xcd\xeb\xf8\xd6\x35\x29\xfc\x60\x7c\xf4\xe4\xc5\x92\x17\x4a\xaa\xff\xf4\x35\xb7\x74\x7c\x7c\xb0\x43\xfe\xfd\x3b\xdb\x77\xf7\x18\x9d\xf6\x9e\x09\x31\x36\x60\x5a\x31\x2b\x75\x89\x70\xa1\xe1\xec\x68\x41\xe7\xf5\x1c\x8b\xa4\x19\x7d\x12\xa2\x92\x96\xd9\xcd\xf6\x52\x7d\x68\xb9\xef\xdd\xc3\xc7\xaf\x63\x7b\xd2\x2b\xe7\xff\x2e\x2e\xd1\xac\x9c\xcc\x68\x37\xb1\xa2\xed\x65\x9b\xeb\x1e\x3f\x9b\x9d\x99\x2f\xe7\x7c\xef\x54\x5e\xa3\xa2\x79\xe4\xf9\x91\xc3\x6c\xb1\x13\x7e\x4c\x7c\xbd\xe9\x74\xfe\x9c\x6e\x3e\x6e\x1b\xe3\xa4\x94\x4b\xc5\x47\x04\x93\x36\xde\x38\x56\x24\xbd\x79\xd2\x42\xa1\x7e\x61\x77\xd1\xd0\xc5\x35\x8a\xaf\x37\x1d\xaa\xec\xde\x1c\xb5\x6e\xd3\x33\xdb\x9f\xdc\x87\x2b\x73\x2b\x3b\x5f\x9b\xb0\x4f\x38\xc0\x24\x36\xb1\x89\x43\x32\x34\x67\x01\x0f\xb7\x61\x94\x69\x90\xcf\x92\x4c\xfb\x77\x0f\xdf\x97\xff\xd1\x2c\x3a\xf9\xe4\x16\x83\xf2\x71\x1b\xad\x00\x45\xcb\xba\x0f\xf6\x3f\xee\xae\x8d\x94\x90\x2e\x9d\x7c\xa8\x50\xe1\x47\xd2\x6d\xf9\x66\x87\x0c\xf1\x8c\x2b\x39\xbf\x7f\xdf\xd8\x26\xad\xde\xf0\x5f\x72\xc7\xa5\xa8\x7f\x5e\xb2\x33\x7e\xa9\xec\xbd\x76\x69\x91\xbb\xc3\x79\xb3\x82\xc2\xbb\x97\x9e\xd8\x3e\x28\x62\x7f\x51\x7d\x49\xd5\xb9\x48\xc5\xde\x77\x61\x00\xc3\xbd\xf3\x2f\x66\x17\x39\xbc\xd7\xbd\xb0\xf3\xd5\xb7\x6b\x77\x8e\xbf\x64\x9e\xb0\x70\xdd\x24\xf3\x03\x07\x15\x0d\x42\x96\xbc\x31\xed\x2f\x2f\xdf\xf2\xfd\xdb\xfb\x99\x69\xf5\xf6\xbb\x19\x8a\x8e\xcc\xea\x6c\xf0\x65\x60\x60\xe0\x2c\xf0\x88\x2c\x66\x60\xe0\x16\x06\x61\x46\x86\x59\x73\x24\x18\x18\x18\xd8\x4b\x3c\x7d\x5d\xd9\x9f\xb0\xf2\xb1\x0a\xcb\x37\x0b\xb5\xfe\x66\x60\x60\x90\x2c\x71\x8d\x28\x71\xce\xcf\xcd\x4d\xcd\x2b\x61\x70\x2e\x4a\x4d\x2c\x49\x4d\x51\x28\xcf\x2c\xc9\x50\x70\xf7\xf4\x0d\x08\x6f\xe4\x13\x67\x60\x60\x52\xf7\x74\x71\x0c\xa9\xb8\xf5\xf6\xd2\x41\x4e\x06\x03\x1e\xe6\x8d\x4b\x7e\xbf\x7f\xa3\x77\xa1\xde\x73\x82\x9a\xc1\x33\x2b\xcf\xac\x76\xb9\x6d\xe5\xa7\xdf\xfb\x2f\x77\x88\xe0\xe0\x4e\x6c\xdc\x20\x20\x76\x88\x29\x51\x51\xa7\x85\xf5\x90\x53\x04\x07\x77\x4b\xeb\x06\x01\x31\x0e\xee\x44\x45\x1d\x01\xb1\x43\x4e\x11\x0a\x3a\x2d\xad\x1b\xe8\xa4\x30\x71\x27\xc3\xa0\x70\x07\x9a\xc2\xc1\xe4\x16\xb0\x42\xea\x58\xf1\xf1\x2e\x2b\xaf\x08\x7b\xc8\x47\xd1\xed\x9b\x19\x18\x18\x18\x3c\x5d\xfd\x5c\xd6\x39\x25\x34\x01\x02\x00\x00\xff\xff\xe3\x65\x6d\x00\x3e\x04\x00\x00")

func imgPausePausePngBytes() ([]byte, error) {
	return bindataRead(
		_imgPausePausePng,
		"img/pause/pause.png",
	)
}

func imgPausePausePng() (*asset, error) {
	bytes, err := imgPausePausePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "img/pause/pause.png", size: 1086, mode: os.FileMode(438), modTime: time.Unix(1589441078, 0)}
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
	"img/pause/pause.png": imgPausePausePng,
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
		"pause": &bintree{nil, map[string]*bintree{
			"pause.png": &bintree{imgPausePausePng, map[string]*bintree{}},
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
