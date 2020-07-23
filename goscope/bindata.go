// Code generated by go-bindata. DO NOT EDIT.
// sources:
// static/css/goscope.css
// static/css/highlight.css
// static/html/common_footer.html
// static/html/common_head.html
// static/html/common_navbar.html
// static/html/log_dashboard.html
// static/html/request_dashboard.html
// static/html/single_log.html
// static/html/single_request.html
// static/js/abstractDashboard.js
// static/js/logsDashboard.js
// static/js/requestDashboard.js
// static/js/singleLog.js
// static/js/singleRequest.js
// static/js/utils.js

package goscope

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// bindataStaticCssGoscopecss reads file data from disk. It returns an error on failure.
func bindataStaticCssGoscopecss() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/css/goscope.css"
	name := "static/css/goscope.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticCssHighlightcss reads file data from disk. It returns an error on failure.
func bindataStaticCssHighlightcss() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/css/highlight.css"
	name := "static/css/highlight.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlCommonfooterhtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlCommonfooterhtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/common_footer.html"
	name := "static/html/common_footer.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlCommonheadhtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlCommonheadhtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/common_head.html"
	name := "static/html/common_head.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlCommonnavbarhtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlCommonnavbarhtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/common_navbar.html"
	name := "static/html/common_navbar.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlLogdashboardhtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlLogdashboardhtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/log_dashboard.html"
	name := "static/html/log_dashboard.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlRequestdashboardhtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlRequestdashboardhtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/request_dashboard.html"
	name := "static/html/request_dashboard.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlSingleloghtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlSingleloghtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/single_log.html"
	name := "static/html/single_log.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticHtmlSinglerequesthtml reads file data from disk. It returns an error on failure.
func bindataStaticHtmlSinglerequesthtml() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/html/single_request.html"
	name := "static/html/single_request.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticJsAbstractDashboardjs reads file data from disk. It returns an error on failure.
func bindataStaticJsAbstractDashboardjs() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/js/abstractDashboard.js"
	name := "static/js/abstractDashboard.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticJsLogsDashboardjs reads file data from disk. It returns an error on failure.
func bindataStaticJsLogsDashboardjs() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/js/logsDashboard.js"
	name := "static/js/logsDashboard.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticJsRequestDashboardjs reads file data from disk. It returns an error on failure.
func bindataStaticJsRequestDashboardjs() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/js/requestDashboard.js"
	name := "static/js/requestDashboard.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticJsSingleLogjs reads file data from disk. It returns an error on failure.
func bindataStaticJsSingleLogjs() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/js/singleLog.js"
	name := "static/js/singleLog.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticJsSingleRequestjs reads file data from disk. It returns an error on failure.
func bindataStaticJsSingleRequestjs() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/js/singleRequest.js"
	name := "static/js/singleRequest.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// bindataStaticJsUtilsjs reads file data from disk. It returns an error on failure.
func bindataStaticJsUtilsjs() (*asset, error) {
	path := "/Users/joe/workspace/goscope/goscope/static/js/utils.js"
	name := "static/js/utils.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"static/css/goscope.css":             bindataStaticCssGoscopecss,
	"static/css/highlight.css":           bindataStaticCssHighlightcss,
	"static/html/common_footer.html":     bindataStaticHtmlCommonfooterhtml,
	"static/html/common_head.html":       bindataStaticHtmlCommonheadhtml,
	"static/html/common_navbar.html":     bindataStaticHtmlCommonnavbarhtml,
	"static/html/log_dashboard.html":     bindataStaticHtmlLogdashboardhtml,
	"static/html/request_dashboard.html": bindataStaticHtmlRequestdashboardhtml,
	"static/html/single_log.html":        bindataStaticHtmlSingleloghtml,
	"static/html/single_request.html":    bindataStaticHtmlSinglerequesthtml,
	"static/js/abstractDashboard.js":     bindataStaticJsAbstractDashboardjs,
	"static/js/logsDashboard.js":         bindataStaticJsLogsDashboardjs,
	"static/js/requestDashboard.js":      bindataStaticJsRequestDashboardjs,
	"static/js/singleLog.js":             bindataStaticJsSingleLogjs,
	"static/js/singleRequest.js":         bindataStaticJsSingleRequestjs,
	"static/js/utils.js":                 bindataStaticJsUtilsjs,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"static": {Func: nil, Children: map[string]*bintree{
		"css": {Func: nil, Children: map[string]*bintree{
			"goscope.css": {Func: bindataStaticCssGoscopecss, Children: map[string]*bintree{}},
			"highlight.css": {Func: bindataStaticCssHighlightcss, Children: map[string]*bintree{}},
		}},
		"html": {Func: nil, Children: map[string]*bintree{
			"common_footer.html": {Func: bindataStaticHtmlCommonfooterhtml, Children: map[string]*bintree{}},
			"common_head.html": {Func: bindataStaticHtmlCommonheadhtml, Children: map[string]*bintree{}},
			"common_navbar.html": {Func: bindataStaticHtmlCommonnavbarhtml, Children: map[string]*bintree{}},
			"log_dashboard.html": {Func: bindataStaticHtmlLogdashboardhtml, Children: map[string]*bintree{}},
			"request_dashboard.html": {Func: bindataStaticHtmlRequestdashboardhtml, Children: map[string]*bintree{}},
			"single_log.html": {Func: bindataStaticHtmlSingleloghtml, Children: map[string]*bintree{}},
			"single_request.html": {Func: bindataStaticHtmlSinglerequesthtml, Children: map[string]*bintree{}},
		}},
		"js": {Func: nil, Children: map[string]*bintree{
			"abstractDashboard.js": {Func: bindataStaticJsAbstractDashboardjs, Children: map[string]*bintree{}},
			"logsDashboard.js": {Func: bindataStaticJsLogsDashboardjs, Children: map[string]*bintree{}},
			"requestDashboard.js": {Func: bindataStaticJsRequestDashboardjs, Children: map[string]*bintree{}},
			"singleLog.js": {Func: bindataStaticJsSingleLogjs, Children: map[string]*bintree{}},
			"singleRequest.js": {Func: bindataStaticJsSingleRequestjs, Children: map[string]*bintree{}},
			"utils.js": {Func: bindataStaticJsUtilsjs, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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