package goscope_css

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"log"
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
