package goscope_js

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
	"log"
	"regexp"
)

func MinifyJs(uncompressed string) string {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minified, err := m.String("application/javascript", uncompressed)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return minified
}
