// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/.iplan.yaml
// assets/.templates/userstory.md
// assets/gitignore
package configs

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

var _IplanYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x4a\x2c\x2d\xc9\xc8\x2f\xb2\x52\x88\xf4\x0f\x0d\x52\xf0\x73\xf4\x75\x55\xb0\xa9\xcc\x2f\x75\x48\xcd\x4d\xcc\xcc\xd1\xcb\x4d\xb5\xe3\x02\x04\x00\x00\xff\xff\x08\xfc\x42\xa4\x25\x00\x00\x00")

func IplanYamlBytes() ([]byte, error) {
	return bindataRead(
		_IplanYaml,
		".iplan.yaml",
	)
}

func IplanYaml() (*asset, error) {
	bytes, err := IplanYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".iplan.yaml", size: 37, mode: os.FileMode(420), modTime: time.Unix(1613533059, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesUserstoryMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\xd5\xe5\x2a\xc9\x2c\xc9\x49\xb5\x52\x08\x2d\x4e\x4d\x2b\xcd\x51\x08\x01\xf1\xb8\x12\x4b\x4b\x32\xf2\x8b\xac\x14\x22\xfd\x43\x83\x14\xfc\x1c\x7d\x5d\x15\x6c\x2a\xf3\x4b\x1d\x8a\xf3\x73\x53\x4b\x32\x32\xf3\xd2\xf5\x92\xf3\x73\xed\xb8\x40\xba\x01\x01\x00\x00\xff\xff\x37\x54\x32\x52\x42\x00\x00\x00")

func TemplatesUserstoryMdBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesUserstoryMd,
		".templates/userstory.md",
	)
}

func TemplatesUserstoryMd() (*asset, error) {
	bytes, err := TemplatesUserstoryMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".templates/userstory.md", size: 66, mode: os.FileMode(420), modTime: time.Unix(1613532170, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xcb\x2c\xc8\x49\xcc\xd3\xd3\xe2\x02\x04\x00\x00\xff\xff\x14\x79\x06\x32\x09\x00\x00\x00")

func gitignoreBytes() ([]byte, error) {
	return bindataRead(
		_gitignore,
		"gitignore",
	)
}

func gitignore() (*asset, error) {
	bytes, err := gitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gitignore", size: 9, mode: os.FileMode(420), modTime: time.Unix(1613533059, 0)}
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
	".iplan.yaml":             IplanYaml,
	".templates/userstory.md": TemplatesUserstoryMd,
	"gitignore":               gitignore,
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
	".iplan.yaml": &bintree{IplanYaml, map[string]*bintree{}},
	".templates": &bintree{nil, map[string]*bintree{
		"userstory.md": &bintree{TemplatesUserstoryMd, map[string]*bintree{}},
	}},
	"gitignore": &bintree{gitignore, map[string]*bintree{}},
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
