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

const CommonHead = `
<meta charset="UTF-8">
<title>{{.APPLICATION_NAME}} - GoScope</title>
<link href="https://cdn.jsdelivr.net/gh/tretapey/raisincss@1.1.0/raisin.min.css" rel="stylesheet"/>
<link href="https://fonts.googleapis.com/css2?family=Manrope&display=swap" rel="stylesheet"> 
<link href="https://fonts.googleapis.com/css2?family=Fira+Code&display=swap" rel="stylesheet"> 
<link rel="apple-touch-icon" sizes="180x180" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595014987/git-repositories/goscope/application/lrhheuzhlvgtsxyayru3.png">
<link rel="icon" type="image/png" sizes="32x32" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015010/git-repositories/goscope/application/blaaguzlyw9s5x4jbhgq.png">
<link rel="icon" type="image/png" sizes="16x16" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015010/git-repositories/goscope/application/bchsism3hubielueb6xk.png">
<link rel="mask-icon" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015042/git-repositories/goscope/application/wswxoxpwirad8udbkwz8.svg" color="#5bbad5">
<meta name="msapplication-TileColor" content="#ffffff">
<meta name="theme-color" content="#ffffff">
`