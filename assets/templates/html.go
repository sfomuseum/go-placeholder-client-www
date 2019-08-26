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

var _templatesHtmlInc_footHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\x52\xa8\xad\xe5\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x56\x48\xcd\x4b\x01\xc9\x01\x02\x00\x00\xff\xff\x35\xa4\x0d\x91\x32\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/inc_foot.html", size: 50, mode: os.FileMode(420), modTime: time.Unix(1566839780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\x52\xa8\xad\xe5\xb2\xc9\x28\xc9\xcd\xb1\xe3\x52\x50\x50\x50\xb0\x01\x09\xdb\x71\x71\xda\x94\x64\x96\xe4\xa4\xda\xd9\xe8\x43\x68\x88\x9c\x3e\x44\x12\xcc\x4e\xca\x4f\xa9\xb4\xe3\xaa\xae\x56\x48\xcd\x4b\x01\x99\x01\x08\x00\x00\xff\xff\xc6\xc9\x76\x22\x5a\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 90, mode: os.FileMode(420), modTime: time.Unix(1566839804, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlSearchHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\xcd\x8e\xe3\x36\x0c\x3e\x27\x4f\x41\x08\x73\x6c\xac\xc3\xde\x06\x8a\x81\xa2\x1d\x2c\x16\x58\x14\xed\xb6\x7d\x00\xc5\xa2\x6d\xb5\xb2\xe4\xa1\xe8\x99\x06\xc6\xbc\x7b\x21\xd9\x72\x9c\xfd\xc3\x5e\x0c\x89\xe4\x27\x7e\xfc\x48\xc9\xf3\x0c\x06\x5b\xeb\x11\x44\x44\x4d\x4d\x2f\xe0\xed\xed\x38\xcf\xc0\x38\x8c\x4e\x33\x82\xe8\x51\x1b\x24\x01\x55\xf2\x28\x63\x5f\xa0\x71\x3a\xc6\xb3\x68\x82\x67\x6d\x3d\x92\xa8\x8f\x47\x00\x80\xbd\x93\xc2\xab\xa8\x8f\x07\xd5\x06\x1a\x8a\x2d\xad\x4f\xd6\x3b\xeb\x51\xc0\x80\xdc\x07\x73\x16\xef\x9f\xfe\x4a\x81\x9f\xe3\x73\x6c\x47\x61\x1a\x93\xf7\xa0\xac\x1f\x27\x06\xbe\x8e\x78\x16\x8c\xff\xb1\xb8\x0b\x4c\x54\x28\x38\x01\x5e\x0f\x5b\x80\x35\x65\x35\x3a\xdd\x60\x1f\x9c\x41\x3a\x8b\x27\xcf\x48\xa0\x17\x63\x06\x00\x07\x58\xaa\x87\x36\x90\x80\x17\xed\x26\x3c\x8b\x79\x86\xea\x8f\x09\xe9\x0a\x6f\x6f\x02\x64\xad\x64\x26\x51\xc8\x4a\x63\x5f\xca\xfa\x32\x31\x07\xbf\xd2\x8b\xd3\x65\xb0\x37\x82\x17\xf6\x70\x61\x7f\x1a\xc9\x0e\x9a\xae\xa2\xfe\x33\xa7\x52\x72\x01\x25\x91\x64\x2a\xa2\x3e\xee\x8e\x4d\xcb\xd4\x06\xdb\xde\x38\x7c\x4b\xe3\xe3\x41\xf5\xef\xea\x25\x8a\x30\x4e\x8e\x63\xaa\x03\xd4\x73\xbd\x2f\x41\xc9\xe7\x5a\xc9\xfe\x5d\x46\xcc\xb3\x6d\x01\xaa\x27\xa2\x40\xe9\xe8\xc3\xfe\x5c\xed\x90\x18\xf2\xf7\xf4\xaa\xc9\x5b\xdf\x09\xa0\xe0\x70\x75\x6d\xfd\x1a\xeb\xbf\xbd\xbe\xb8\x2c\x60\x13\x86\xd1\x21\x23\x3c\xe7\x7c\x71\x6a\x1a\x8c\xb1\x9d\x9c\xbb\x56\xb0\xe4\x21\x1c\x03\x31\x1a\xb0\xf1\x51\xc9\xb1\x9c\xd2\x04\x83\x99\x69\x61\xa3\x64\x36\x25\x61\x16\x8d\xe7\x19\xd0\x45\xcc\x44\x57\xaa\xa9\xb9\x83\x1e\x45\xbd\xc6\x1c\x0f\x8a\x33\x95\xb5\x86\x65\x93\xbf\xa7\xc8\x64\x47\x34\x1b\x6d\xa6\x3c\x52\xdc\xd7\x1f\x7e\x55\x92\xfb\xb2\xfb\x4d\x0f\xb8\xdf\xff\x9e\x26\x24\xf5\x74\x6f\xfc\x84\x9d\x0d\x3e\x5b\x0e\x87\xd5\xf6\x4b\x98\x3c\xd3\x75\x1f\xf6\x51\xb3\xe5\xc9\xdc\x41\x3f\x06\xdf\xdd\x8c\x09\xbd\x74\x3c\xf3\x49\xcb\x79\x06\xd2\xbe\x43\x78\xb0\x3f\xc1\x03\xc1\xe3\x19\xaa\x4f\x6b\x47\x53\xed\x2b\xfb\xad\xff\xd9\xb5\x0c\xfa\xb2\x3e\xcd\x33\x3c\x50\xf5\xc1\xa4\x91\x5d\xd2\x9a\x5a\x69\xe8\x09\xdb\xb3\xe8\x99\xc7\xf8\x28\x65\x1c\xd1\x4d\xfe\x5f\xa4\xea\xb5\x0f\x31\xf8\xd6\x52\xe4\x2a\x50\x27\xad\x91\x77\x27\xec\x36\x4a\xea\x5a\x49\x36\xeb\xa9\xf7\x1c\x4e\xe9\x1a\x95\xf0\x24\x63\x06\x7c\x2b\x78\x2c\xba\x16\xc4\x26\xf4\x77\x61\xcd\xa2\x72\x06\xad\x32\xfd\x93\x64\xea\x92\x4e\x3f\xfb\x06\x23\x07\x8a\x49\x37\x41\xb9\x49\xe9\x2d\xcb\x09\xba\xc2\x29\x0d\x92\x37\x25\xcb\xda\xbf\x1f\xcc\xd3\x7c\x99\xa6\x44\xae\x79\x9a\xaf\xa7\xf9\x6a\x12\xb7\xce\x47\x91\xe0\x3d\x86\x01\x99\xae\x55\x19\x9c\xef\xa3\xcb\x24\x7d\x09\x2f\x9e\x7d\x91\xbb\x31\xdb\xa8\xa5\xdb\x95\xaf\xc7\xfa\x1e\x6c\xe6\xcf\x9f\xa1\x9b\x2b\x3f\x4a\xab\xe7\xee\x17\xd1\x86\xc0\xdb\x2f\xe2\x16\xff\x7f\x00\x00\x00\xff\xff\xe0\x21\xb1\xfc\x5d\x06\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/search.html", size: 1629, mode: os.FileMode(420), modTime: time.Unix(1566848118, 0)}
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

