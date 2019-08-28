// Code generated by go-bindata.
// sources:
// templates/html/inc_foot.html
// templates/html/inc_head.html
// templates/html/search.html
// DO NOT EDIT!

package templates

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

var _templatesHtmlInc_footHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\x52\xa8\xad\xe5\xe2\xb4\xd1\x4f\xc9\x2c\xb3\xe3\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x56\x48\xcd\x4b\x01\xa9\x01\x04\x00\x00\xff\xff\x1a\x92\x03\x62\x3a\x00\x00\x00")

func templatesHtmlInc_footHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInc_footHtml,
		"templates/html/inc_foot.html",
	)
}

func templatesHtmlInc_footHtml() (*asset, error) {
	bytes, err := templatesHtmlInc_footHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inc_foot.html", size: 58, mode: os.FileMode(420), modTime: time.Unix(1567012776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xc1\x8e\x9c\x30\x0c\x86\xcf\xf0\x14\x96\x2b\xf5\x96\xc9\x76\x0f\xbd\x34\x70\xeb\x03\x54\x7d\x02\x13\x3c\x4b\xb6\x21\xa0\xc4\xa0\x22\x34\xef\x5e\x85\xcc\x4c\x77\xdb\x3d\xce\x85\x58\xe6\xf7\xe7\x3f\x3f\x62\xdf\xa1\xe7\xb3\x0b\x0c\x38\x30\xf5\x1c\x11\x2e\x97\xda\x0c\x32\xfa\xb6\x06\x00\x30\xb9\xdd\xd6\x95\x11\x27\x9e\x5b\xa3\xcb\x59\x57\x26\xd9\xe8\x66\x01\xd9\x66\x6e\x50\xf8\xb7\xe8\x57\x5a\xa9\x74\x11\x52\xb4\x0d\xfe\x6d\xe8\xd9\x93\xe5\x61\xf2\x3d\xc7\x93\xf5\x8e\x83\x9c\x46\x9a\xd3\xe9\x35\x61\x6b\x74\x11\x3d\x84\x1a\x39\x2d\x5e\xfe\x01\x57\x8f\x20\xbb\xe0\xe4\x3f\xbf\xde\x85\x5f\x6f\x99\x36\x25\x84\xc8\xbe\xc1\x24\x9b\xe7\x34\x30\x0b\xc2\x10\xf9\xdc\xa0\x4d\xe9\x23\xee\x31\xa2\xaf\x71\xeb\x92\xf7\x51\x77\x53\xbf\xe5\x1d\xbd\x5b\xc1\x7a\x4a\xa9\x41\x3b\x05\x21\x17\x38\x62\x5b\xd7\xd5\xa1\x0a\x74\x7f\x1b\x68\xed\x28\x42\x39\x94\x77\x2f\x83\x40\xf7\x52\x8a\x63\xa0\x32\xf4\x5e\xab\xba\x48\xa1\xc7\xf6\xf3\xa7\x2f\x5f\x9f\xbe\x19\x4d\x45\x75\x9e\xe2\x78\x13\xe6\x5a\xb9\xe0\x5d\x60\x6c\xeb\xaa\x2c\x75\x61\x5e\xe4\x9d\x22\x3b\x8b\x93\x87\x31\xaa\x34\xaa\x67\xbc\x86\x92\x98\xa2\x1d\x10\x02\x8d\xd7\x88\x10\x5c\x7f\xab\xde\xa4\xd1\xe0\xf7\x20\x1c\x81\x4a\xf3\x18\x00\x99\xa0\x00\xe0\x3c\x45\x84\x95\xfc\xc2\x0d\xee\x3b\x9c\x7e\x2c\x1c\x37\xb8\x5c\x10\x28\x3a\x52\x9e\xba\x9c\xf9\xcf\xb2\xed\x6e\xb3\x5b\x44\xa6\x70\xf3\xd9\x49\x80\x4e\x82\x9a\x16\xc9\xb7\x51\x69\xb1\x96\x53\x82\x71\x53\xcf\xf9\x91\x46\xf5\x74\xf7\xbd\x74\xa3\x13\x6c\x0b\xd1\xe8\x42\xca\x60\xa3\xf3\x7d\x73\x75\xfd\x00\x3a\xd0\xda\xd6\x55\xbd\xef\xc0\xa1\xcf\xff\xcf\x9f\x00\x00\x00\xff\xff\x08\x93\x43\xf3\x56\x03\x00\x00")

func templatesHtmlInc_headHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInc_headHtml,
		"templates/html/inc_head.html",
	)
}

func templatesHtmlInc_headHtml() (*asset, error) {
	bytes, err := templatesHtmlInc_headHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 854, mode: os.FileMode(420), modTime: time.Unix(1567016608, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlSearchHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdd\x8e\x9b\x3c\x10\xbd\x86\xa7\x18\xa1\xbd\xfc\x80\x8b\xbd\x5b\x39\x48\xdf\xb6\x55\xb5\xd2\xaa\x6a\x57\xea\x03\x38\x78\x00\xb7\x8e\x4d\xc6\x26\xbb\x11\xca\xbb\x57\x36\x38\x81\x4d\xfa\x73\x83\x60\xe6\x9c\x99\x33\x3f\xcc\x38\x82\xc0\x46\x6a\x84\xcc\x22\xa7\xba\xcb\xe0\x74\x4a\xc7\x11\x1c\xee\x7a\xc5\x1d\x42\xd6\x21\x17\x48\x19\x14\xde\x03\x00\xe0\xdd\xb2\x81\xe2\xdb\x80\x74\x8c\x46\x26\xe4\x01\x6a\xc5\xad\xdd\x64\x35\x27\x91\x55\x69\x92\x26\xac\xbb\x5f\x1a\xf3\x39\x56\x35\x51\x09\xed\xa0\x9c\x85\xc6\x10\xb0\x7d\x35\x8e\x97\x98\xac\xdc\x57\xac\xec\xee\xab\x34\x4d\xde\x87\xce\xb7\x46\x1c\x7d\xfc\x20\x26\x3c\xc7\x51\x36\x00\xc5\x27\x22\x43\x5e\xd1\xc2\xb5\x64\x73\x85\xe4\x20\x3c\xf3\x57\x4e\x5a\xea\x36\x03\x32\x0a\x67\x97\x0f\x9a\xb0\xbe\xfa\xae\xf9\x56\x21\x38\x03\xb5\xd9\xf5\x0a\x1d\xc2\x3e\x08\xb3\x43\x5d\xa3\xb5\xcd\xa0\xd4\xb1\x80\x29\x1d\x61\x6f\xc8\xa1\x00\x69\x1f\x58\xd9\x87\x10\xb5\x11\x18\xea\x89\x8a\x58\x19\x4c\xb3\xa4\x52\xc8\xc3\x3b\xfd\x80\xca\xe2\x2d\xe9\x52\x6c\xb2\x1d\xef\xb3\xea\x9a\xb5\x2c\xcd\x79\xc5\x39\xa1\xed\x8d\xb6\xf2\x80\x53\x29\xc1\xba\x82\xc0\x04\xec\xcc\xc1\xcf\x21\x4d\xa6\x40\x8e\xfc\x6b\xc2\x5c\x57\x3d\x7d\x64\xa5\xeb\xce\x9f\x5f\xf8\x0e\x57\x86\xaf\x8a\xd7\xe8\x8e\xfd\xda\xfa\x82\xad\x34\x3a\x98\x92\x24\x1a\x3f\x98\x41\x3b\x3a\xae\x80\xcf\xdc\x49\x37\x88\x35\xfb\xd9\xe8\xf6\x62\x0d\x01\xa6\x3e\x4d\xba\xe6\x0e\x11\xd7\x2d\xc2\x9d\xfc\x0f\xee\x08\x1e\x36\x50\xbc\xcc\xfb\xe3\xbb\x16\xeb\x88\xc5\x4e\xbb\x95\x85\xf6\x4d\xef\xf9\x38\xc2\x1d\x15\x4f\x02\x4e\xa7\x0c\x04\x77\x3c\x7f\xed\x8c\x35\xba\x91\x64\x5d\xee\x81\xd7\x88\xad\x19\xb4\x90\xba\xcd\xb7\xe6\x2d\xfa\x3f\xa3\xd9\xa1\xa3\x63\xf1\x38\x3b\x1f\xcd\x9b\x67\xcc\xe5\x88\x8a\x71\xe8\x08\x9b\x4d\xd6\x39\xd7\xdb\x87\xb2\xb4\x3d\xaa\x41\xff\x44\x2a\x16\x19\x0b\x43\x6d\x29\x45\xb9\x4c\x5a\x2d\x3e\x58\xc9\x2b\x56\x3a\x11\xc3\xae\x2b\xcb\x35\xdf\x61\xc4\xfb\x21\x05\xc6\x6f\xd1\x7d\x9c\x5a\xa4\x9c\xc7\xf8\x67\x5e\x3d\x8d\x30\xb0\xe6\xfe\xff\xf0\xfd\x6f\xfd\x00\xfe\xd7\x35\x5a\x67\xc8\xfa\x81\x64\x14\x56\xc0\xdf\x8f\x90\xa1\x8d\xaa\xfc\x6e\x6b\x11\xd3\xc4\xed\xf8\xc7\x44\xf5\x75\x9e\x88\x9c\x13\xd5\xb7\xf3\xdc\xce\xa2\xe6\xed\x8b\x5d\x38\x8f\x32\xae\xe5\x5f\xe8\x71\x4f\xaf\xf9\xd1\xb3\xaa\xf3\xc6\x12\x4f\x12\xfd\xbf\x59\x86\x3f\xf1\xe6\x49\x58\xe0\xa2\x23\x5d\xa0\xd6\xa1\xe2\x49\xbe\x5c\xec\xc6\x18\x77\xbe\xd8\x17\xe0\xaf\x00\x00\x00\xff\xff\xab\x9e\x20\x4e\xec\x05\x00\x00")

func templatesHtmlSearchHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlSearchHtml,
		"templates/html/search.html",
	)
}

func templatesHtmlSearchHtml() (*asset, error) {
	bytes, err := templatesHtmlSearchHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/search.html", size: 1516, mode: os.FileMode(420), modTime: time.Unix(1567014811, 0)}
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
	"templates/html/inc_foot.html": templatesHtmlInc_footHtml,
	"templates/html/inc_head.html": templatesHtmlInc_headHtml,
	"templates/html/search.html": templatesHtmlSearchHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"inc_foot.html": &bintree{templatesHtmlInc_footHtml, map[string]*bintree{}},
			"inc_head.html": &bintree{templatesHtmlInc_headHtml, map[string]*bintree{}},
			"search.html": &bintree{templatesHtmlSearchHtml, map[string]*bintree{}},
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

