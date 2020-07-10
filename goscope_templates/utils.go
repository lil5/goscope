package goscope_templates

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"log"
	"regexp"
)

func MinifyCss(uncompressed string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	minified, err := m.String("text/css", uncompressed)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return minified
}

func MinifyHtml(uncompressed string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	minified, err := m.String("text/html", uncompressed)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return minified
}

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
