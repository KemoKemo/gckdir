// Code generated by go-bindata.
// sources:
// lib/templates/index.html
// DO NOT EDIT!

package lib

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _libTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x57\x5d\x6f\xdb\xb8\x12\x7d\x2f\xd0\xff\xc0\xba\xb8\x6f\x92\x6c\x39\xb6\x9b\xca\xb2\x80\xa2\x69\x71\x0b\x5c\xe0\x16\xdd\x6c\xb1\xfb\x48\x8b\x23\x93\x58\x8a\xd4\x92\x63\xc7\x59\xad\xfe\xfb\x42\x9f\xb6\x25\x26\xd8\xe4\x21\x22\xcf\xcc\xe1\x99\xc3\xcf\xc4\xef\x1e\xfe\xff\xf9\xf1\xf7\xef\x5f\x08\xc7\x5c\x26\x6f\xdf\xc4\xf5\x5f\x22\xa9\x3a\xec\x66\xa0\x66\xc9\xdb\x37\x84\xc4\x1c\x28\x6b\xbe\x08\x89\x73\x40\x4a\x52\x4e\x8d\x05\xdc\xcd\x8e\x98\xf9\xf7\xb3\x1b\x8c\x23\x16\x3e\xfc\x79\x14\xa7\xdd\xec\x37\xff\xd7\x4f\xfe\x67\x9d\x17\x14\xc5\x5e\xc2\x8c\xa4\x5a\x21\x28\xdc\xcd\xbe\x7d\xd9\x01\x3b\xc0\x6d\xaa\xa2\x39\xec\x66\x27\x01\x4f\x85\x36\x78\x15\xfd\x24\x18\xf2\x1d\x83\x93\x48\xc1\x6f\x1a\x1e\x11\x4a\xa0\xa0\xd2\xb7\x29\x95\xb0\x0b\x07\x26\x8b\xcf\x12\x08\x3e\x17\xb0\x9b\x21\x9c\x71\x9e\x5a\x3b\x80\xef\x7c\xbf\xfd\xda\x6b\xf6\x5c\x66\x5a\xa1\x9f\xd1\x5c\xc8\xe7\x88\x1a\x41\xa5\x67\xa9\xb2\xbe\x05\x23\x32\x2f\xd7\x4a\xdb\x82\xa6\x50\x05\x3c\xf4\x02\xbe\xf4\x02\x7e\xe7\xf1\xd0\xe3\x4b\x8f\xdf\x95\x39\x35\x07\xa1\x7c\xd4\x45\xb4\x5c\x14\xe7\x6d\xd7\xde\x6b\x44\x9d\x47\xe1\xa2\x38\xdf\xe4\x05\x7c\xe5\x05\x7c\xed\x05\x7c\x33\x70\x78\x7c\xe5\xf1\xb5\xc7\x37\x37\x42\x84\xe2\x60\x04\x6e\x9b\xbe\x27\x10\x07\x8e\xd1\x7a\xb1\xd8\x4a\xa1\xc0\xe7\x6d\x3b\x0c\xc2\x6d\xaa\xa5\x36\x7d\x74\xc5\xc3\x4e\x51\x14\x6c\x3e\x40\x4e\x16\x6d\xbe\x15\x7f\x41\xb4\x84\xbc\xe2\xcb\xb2\xa0\x8c\x09\x75\x88\x82\xe5\x1a\x72\x12\xac\x21\xef\x38\xde\xaf\x3e\xd6\xbf\xdb\x3d\x4d\xff\x38\x18\x7d\x54\x2c\x42\x43\x95\x2d\xa8\x01\x85\xdb\xbd\x36\x0c\x8c\x2f\x21\xc3\xc8\x6a\x29\x18\x59\x17\x67\xf2\x7e\xb3\xd9\x54\x01\x95\x60\x70\x60\x0e\xd7\x13\x23\x1a\x6f\x5a\x86\x28\x2c\xce\xa4\x25\x70\xd0\x1b\xca\xc4\xd1\x46\xab\xda\xb7\x86\xd5\xb7\xc7\x34\x05\x6b\xcb\x4e\xe4\x5d\xfa\x61\x73\xc7\xae\x44\xfa\x1d\xc0\xb2\x6c\xc1\xee\x7b\x9e\xbe\x73\x03\x1f\xd3\x5e\xa0\xcf\xa8\x3a\x80\xe9\x99\xe8\xc7\xd5\x6a\xb5\x74\x30\x65\x4b\x06\x0c\x46\x4c\xb0\x4f\x53\x16\x56\x01\xd2\xbd\x04\xbf\xc5\x80\x95\x93\x9a\xde\x33\xc6\xba\xa8\xb2\x59\xa2\x51\xb8\x58\xfc\x67\x9b\xd3\xb3\x7f\xd3\x1c\xbb\x53\xb5\x29\x13\x35\x0e\x8f\xea\xe5\x58\xfb\xbc\xb8\x92\x28\x69\x61\x21\xea\x3f\x3a\x01\x09\xd6\xeb\x3b\x41\x13\x74\x26\x26\xc8\xbc\x17\x21\x3e\x86\x12\x64\x3d\x3a\x85\xf8\x04\xca\xb4\xc6\x17\xc6\x9a\x40\x7c\x0c\xb9\xc6\x1a\xa0\xe9\x58\xcd\x51\xe4\x1e\x6b\x02\xf1\x31\xe4\x1a\x6b\x80\x86\xb1\xa6\x53\xd1\x2d\xb1\x89\xb9\xed\xb2\x72\x79\xdb\x23\x4e\x6b\x5b\xd0\xe5\xec\x08\xe9\xdd\x9b\x8e\x33\x46\x9c\xb6\xba\xd9\x1c\xe3\xf4\xce\x4d\xc7\x19\x23\x4e\x4b\xdd\x6c\x97\x71\x1c\x86\xb6\x3b\x6d\xbc\xa9\xae\x5d\xf2\x5e\xc6\xf8\x14\xbb\x54\xfd\x0a\xe6\xc8\xbb\x54\xf1\x0a\xc6\x5f\xdd\xec\x0e\xd1\x0e\xad\x0e\x89\x0e\x65\x0e\x41\x37\x3a\xfa\x83\xf6\xbe\x38\x8f\xee\x83\xd5\xf2\x7e\xfd\x21\x5c\xdd\x6d\x4f\x60\x50\xa4\x54\xfa\x54\x8a\x83\x8a\x50\x17\xfd\x61\x51\x5f\x55\xa3\x0a\xda\xbb\xd0\xf7\xfb\xeb\x71\xde\x5c\x9e\x7d\x0b\x05\x4a\x48\x7e\xd6\xb7\xa1\x48\x29\x0a\xad\xc8\x0f\xb0\x47\x89\xf1\xbc\x85\x9a\xc7\xc1\x7c\x78\x1d\xc4\x4d\xd5\x5d\x32\x0f\x93\x07\x61\x20\x45\x6d\x9e\x89\x93\x83\x87\x43\xec\x32\xf9\xa6\x32\x6d\xf2\x26\x20\x9e\xf3\x65\x8f\x14\xc9\x23\x17\x96\x98\x26\x85\x3c\x51\x4b\x0e\xa0\xc0\x50\x04\x46\x32\xa3\x73\x82\x1c\x48\xa6\xa5\xd4\x4f\x42\x1d\x88\xd5\x47\x93\x02\xa1\x8a\x11\xa4\xe6\x00\x18\xc4\xf3\xa2\xe7\x3a\xca\xee\x8b\x90\x58\x8a\xe4\x97\x26\x36\x22\x65\x19\x7c\xa7\xc8\xff\x27\x2c\x06\x6d\x5f\xdd\xac\xaa\x78\x2e\xc5\x4d\xc2\x63\xc3\x78\x9b\xd0\xf6\x4d\x12\xe2\xf9\x30\xd8\xa8\x84\xfa\x49\x43\x85\xb2\x23\xe1\xea\x98\xef\xc1\x10\x9d\x11\xd6\x99\x26\xc0\x36\x75\x64\x42\x82\x7d\xa5\x8c\x87\x4b\x7c\x23\x6d\x30\xbd\x76\x34\xb8\x42\x1d\x15\x7d\xad\xb9\x1d\x59\x4d\xff\xb4\xa0\x78\x6f\xae\x66\xec\x51\x23\x95\x97\xc9\x1c\xa6\x8c\x89\x13\x49\x25\xb5\x76\x37\x6b\x2e\x5f\xd2\x5e\xc1\x65\x19\xb4\xb1\x41\xb3\x18\x9e\xdb\x06\xf9\x9b\x18\xfd\xf4\x09\xd1\x54\xd5\x6c\xd0\xd6\xbd\xe3\x8c\x56\x87\x24\xae\xdf\x31\xa4\x7e\xc7\xec\x66\xeb\x59\xf2\x22\x4d\x41\xad\xfd\x4a\x85\xac\x65\xd7\x29\x49\xbd\x98\x1b\x86\xbe\x02\x26\x4e\xa3\x12\x1e\x00\xa9\x90\xf6\x5a\x7d\xb3\xe7\x7a\xfd\x6d\xe3\xf6\x5c\xb8\xa8\x8c\xf1\xea\x5d\xdc\x76\x98\xdb\x12\x90\x27\x3f\x40\x52\x14\xa7\x66\x4d\xc5\x73\xe4\x93\x80\xff\x52\xcb\x7f\x52\x79\x04\x27\x3a\xec\x37\x17\x44\x6d\xbd\x59\x6e\xa0\x78\x7e\xa5\xa1\xc6\xae\x15\xc6\x78\xb5\x41\xdb\x9f\xb2\x34\xf5\x11\x4d\x7a\x53\xeb\x45\x5d\x55\x31\x9a\xde\x82\xb2\xfc\x97\xf3\x55\xd3\xb3\x76\x7a\x2e\x15\xd7\x93\x81\xcc\x1d\x37\x14\xfe\x5a\xd0\x2b\x93\xfc\x52\x4a\xeb\x8b\x23\xe2\xc6\x9b\xb6\x78\x50\xac\xaa\x6e\xdc\xbb\x3e\xc2\xe6\xed\x01\xdc\x9e\x72\x1d\x10\xcf\xbb\x7f\x94\xfe\x09\x00\x00\xff\xff\x37\xe7\xf6\x0f\x3a\x0d\x00\x00")

func libTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_libTemplatesIndexHtml,
		"lib/templates/index.html",
	)
}

func libTemplatesIndexHtml() (*asset, error) {
	bytes, err := libTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/templates/index.html", size: 3386, mode: os.FileMode(438), modTime: time.Unix(1526313538, 0)}
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
	"lib/templates/index.html": libTemplatesIndexHtml,
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
	"lib": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{libTemplatesIndexHtml, map[string]*bintree{}},
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

