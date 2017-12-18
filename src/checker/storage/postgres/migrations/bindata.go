// Code generated by go-bindata.
// sources:
// src/checker/storage/postgres/migrations/0001_base_init.down.sql
// src/checker/storage/postgres/migrations/0001_base_init.up.sql
// DO NOT EDIT!

package migrations

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

var __0001_base_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xce\x48\x4d\xce\x4e\x2d\x8a\xcf\xc9\x4f\xb7\xe6\xe2\x02\x4b\xb8\x85\xfa\x39\x87\x78\xfa\xfb\x29\xe4\x26\x56\xc4\x67\xe6\x25\x26\x97\x64\x96\xa5\xc6\xe7\xa4\xe6\xa5\x97\x64\x68\x94\x64\xe6\xa6\x16\x97\x24\xe6\x16\x94\x54\xe9\x28\xa0\x72\x52\x2b\x4a\x34\xad\xb9\x00\x01\x00\x00\xff\xff\xd2\xf0\x20\x98\x5c\x00\x00\x00")

func _0001_base_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_base_initDownSql,
		"0001_base_init.down.sql",
	)
}

func _0001_base_initDownSql() (*asset, error) {
	bytes, err := _0001_base_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_base_init.down.sql", size: 92, mode: os.FileMode(420), modTime: time.Unix(1513388581, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0001_base_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x51\x6f\xda\x30\x14\x85\xdf\xfd\x2b\xce\x5b\xd3\x29\x30\xba\x3d\x76\x7b\x30\x60\x06\x52\x48\x2a\x62\x46\xb7\x97\xc8\x98\x5b\x88\x96\x10\x14\x9b\x96\xfd\xfb\xc9\x89\xa0\x4e\xd6\x27\xcb\xd7\xe7\x1c\x7f\xf7\xda\x6c\x30\x40\x51\xed\x0d\x8c\xad\x6a\xb5\x27\x36\x59\x09\x2e\x05\x24\x1f\x47\x02\xfa\x40\xfa\x0f\xd5\x59\x51\xed\x11\x30\x40\xd7\xa4\x2c\xed\x32\x65\x21\x17\x4b\x91\x4a\xbe\x7c\xc2\x66\x21\xe7\xcd\x16\xbf\x93\x58\x20\x4e\x24\xe2\x75\x14\x61\x2a\x66\x7c\x1d\x49\x1c\xab\xb7\xe0\x3e\x64\xc0\xb9\x2e\x20\xc5\xb3\xbc\x49\x5c\xd1\x9c\xb5\x26\x63\x5e\xce\x05\xc6\x49\x12\x09\x1e\xff\x9f\x30\xe3\x51\x2a\xd8\xfd\x23\xbb\xd2\x2d\xe2\xa9\x78\xf6\x68\xb2\x73\x5d\x64\xf9\xee\x82\x24\xee\x32\xbf\x4b\x42\x77\xbd\x8b\x70\x1d\x5b\x32\x16\xaf\x39\xbd\x5d\x03\x7f\x2e\xc4\x06\xd9\x81\x54\x61\x0f\x99\xb1\xca\x1a\xf0\x94\x01\xa9\x88\xc4\x44\x3a\x6b\x08\x5d\xa9\x82\x8c\xa6\x40\x69\x9b\xbf\x52\x66\xf3\x92\x3e\xab\xa2\xc8\x9a\x7d\x6e\xff\x36\x15\x7c\xc2\xc3\x68\x14\x62\x74\x0f\x9e\xa2\x95\x32\x60\xb6\x4a\x96\x0c\x00\x02\x2f\xb2\x29\xc0\xcd\xa0\x0c\x6e\x21\x86\x74\x75\xdc\x99\xd6\xde\x4f\xef\x58\x26\x3c\x15\xd8\xcc\x45\xec\x0f\x51\xba\x7d\x3f\x0c\x22\x9e\x7a\x3c\x4d\x54\x93\x74\xc3\xf2\xc8\xbc\x99\x5d\x8f\xe0\xbf\x93\x5f\xf5\xbb\x00\x70\xa4\x8b\x6d\xc2\x8d\x55\xe5\xc9\x3f\xa1\x8b\xad\x95\xb6\xc1\x1d\x9d\x2a\x7d\xb8\xc3\x4b\x5d\x95\x3d\x39\x06\xde\xd5\xef\xb4\x5e\x1b\xd7\x3c\x9f\xda\x03\x57\xc3\x8f\xd1\x01\xa8\x61\x8f\xb4\xa9\x7d\xdc\x13\x80\x32\x3f\x06\xdb\x61\x8f\xa6\x4b\xeb\x19\x1c\x4e\xe7\xdb\x39\xf4\xb0\x5f\xd9\x7a\x8e\xcd\x5c\xac\x44\x87\x17\xdf\xe0\x5f\x08\x1e\x4f\x5b\x68\x7c\xc7\xd6\xad\x9e\xfb\xc7\x2a\x59\x3f\x61\xfc\x0b\x0f\x21\xbe\x84\xf8\xda\xce\xaa\x6c\x96\xad\x32\xed\xcb\xde\x44\xee\xdb\xbb\x93\x13\xd5\x9a\x8e\x56\xed\x29\x73\xa2\x47\xf6\x2f\x00\x00\xff\xff\xf3\xe0\xcc\x9c\xfb\x03\x00\x00")

func _0001_base_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_base_initUpSql,
		"0001_base_init.up.sql",
	)
}

func _0001_base_initUpSql() (*asset, error) {
	bytes, err := _0001_base_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_base_init.up.sql", size: 1019, mode: os.FileMode(420), modTime: time.Unix(1513560639, 0)}
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
	"0001_base_init.down.sql": _0001_base_initDownSql,
	"0001_base_init.up.sql": _0001_base_initUpSql,
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
	"0001_base_init.down.sql": &bintree{_0001_base_initDownSql, map[string]*bintree{}},
	"0001_base_init.up.sql": &bintree{_0001_base_initUpSql, map[string]*bintree{}},
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

