// Code generated by go-bindata.
// sources:
// views/index.html
// DO NOT EDIT!

package main

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

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xc1\x8e\xd3\x3c\x10\xbe\xe7\x29\x66\x7d\xe9\xbf\xfa\x37\x76\x61\xa9\x10\x55\x12\x09\x41\x0f\x2b\xc1\x89\x22\x21\x55\x3d\xb8\xf6\xb4\x71\xd7\xb1\x43\x3c\x5d\x51\x45\x39\xac\xc4\x05\xde\x84\x2b\x17\x1e\x68\x0f\xf0\x18\xc8\x49\x77\xb7\x2b\x01\x07\x38\x65\xc6\xf3\xcd\xe7\x7c\xdf\x8c\xb3\x13\xed\x15\xed\x6b\x84\x92\x2a\x5b\x24\x49\x36\x7c\xb3\x12\xa5\x2e\x12\x80\x8c\x0c\x59\x2c\xde\xec\x6a\x6c\x5e\xcb\xe6\x12\x09\xe6\x5e\xfb\x4c\x0c\xe7\x11\x51\x21\x49\x28\x89\xea\x14\xdf\xef\xcc\x55\xce\xde\xa5\x6f\x9f\xa7\x2f\x7c\x55\x4b\x32\x2b\x8b\x0c\x94\x77\x84\x8e\x72\x76\x31\xcb\x51\x6f\x90\xf5\x7d\xd6\xb8\x4b\x68\xd0\xe6\x2c\xd0\xde\x62\x28\x11\x89\x41\xd9\xe0\x3a\x67\xa2\xde\xad\xac\x51\x42\x85\x20\xc8\x6b\xcf\x55\x08\xac\x48\x62\xdb\x49\x9a\x2e\xcc\x1a\x2c\xc1\xc5\x0c\x9e\x2d\x7b\xaa\xa0\x1a\x53\x13\x84\x46\xe5\x2c\xfe\x49\x98\x0a\xa1\xb4\xdb\x06\xae\xac\xdf\xe9\xb5\x95\x0d\x72\xe5\x2b\x21\xb7\xf2\x83\xb0\x66\x15\x04\x86\x49\x1a\x4a\x53\x89\x27\x7c\xcc\x27\x77\x29\xaf\x8c\xe3\xdb\xc0\x8a\x4c\x0c\xa4\x7f\xcb\x1f\x7d\x9c\x84\xd2\x5c\x89\x73\xfe\x94\x3f\xbe\xcf\xff\x74\x43\xd1\xa3\xb8\xd4\x7a\x66\xb1\x42\x47\xe1\xbf\x51\x94\x3f\x3a\x7d\x00\x3e\x59\xa0\xd3\x66\xbd\x4c\xd3\x38\x31\x31\xcc\x2a\xc9\x56\x5e\xef\x07\x8f\x62\x4f\x91\xf5\xce\x0d\x07\xc7\x02\x6e\xbd\x25\xb9\x19\xbc\x25\xb9\x61\x10\x97\x20\x67\x8d\xf1\x14\x0b\xff\xae\xbf\x67\x3a\xe7\x8f\xc6\x7c\xdc\xc7\xff\x2b\x5f\xd5\xc6\x62\xf3\x0b\xfd\x47\x06\x24\x00\x00\x11\xcf\x2b\xbf\x73\x74\x90\x7f\x06\x6d\x5f\x00\xe8\xd7\x6e\x0a\xa3\x9b\xeb\xcf\x3f\xbe\x7e\xbb\xb9\xfe\xf8\xfd\xd3\x97\x79\x0f\x39\x00\x0c\x61\x15\xa6\xb0\x38\xa4\x6d\x6b\xd6\xc0\x23\xe2\x95\x09\xd4\x07\x5d\x77\xa8\xc5\x6a\x23\xdd\x06\x7f\x0f\x00\x68\x6f\x6f\x6c\x5b\x3e\x8f\x61\xd7\x8d\xce\x40\x7b\x87\x53\x68\x5b\xfe\xd2\x3b\xec\x3a\xe8\xce\x8e\x38\xd1\xe9\x3b\x8a\x87\xd9\x72\x80\x75\xa7\x51\xf1\xbd\xfc\x4c\x1c\x46\x97\x89\xe1\xf9\xfd\x0c\x00\x00\xff\xff\xb3\x13\x5e\xc2\x97\x03\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 919, mode: os.FileMode(420), modTime: time.Unix(1531877863, 0)}
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
	"views/index.html": viewsIndexHtml,
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
	"views": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{viewsIndexHtml, map[string]*bintree{}},
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

