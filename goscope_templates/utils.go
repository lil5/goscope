package goscope_templates

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"log"
)

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
