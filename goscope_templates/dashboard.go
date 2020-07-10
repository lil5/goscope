package goscope_templates

import (
	"bitbucket.org/prowarehouse-nl/goscope/goscope_css"
	"bitbucket.org/prowarehouse-nl/goscope/goscope_js"
	"fmt"
)

func DashboardView() string {
	const template = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.APPLICATION_NAME}} - GoScope</title>
    <link href="https://cdn.jsdelivr.net/gh/tretapey/raisincss@1.1.0/raisin.min.css" rel="stylesheet"/>
	<link href="https://fonts.googleapis.com/css2?family=Manrope&display=swap" rel="stylesheet"> 
	<link href="https://fonts.googleapis.com/css2?family=Fira+Code&display=swap" rel="stylesheet"> 
    <style>%s</style>
    <style>%s</style>
</head>
<body>
%s
<div class="m-1 p-1 text-center">
	<table id="request-table" style="line-height: 1.6em; margin: 0 auto;">
	</table>
	<div>
		<button id="requests-prev-page" class="paginate-button"><span class="font-4xl">&#8592;</span> prev</button>
		&nbsp;
		<button id="requests-next-page" class="paginate-button">next <span class="font-4xl">&#8594;</span></button>
	</div>
    %s
</div>
<script>%s</script>
<script>%s</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</body>
</html>
`
	return fmt.Sprintf(template, goscope_css.HighlightTheme(), goscope_css.WatcherStyles(), NavbarComponent("REQUESTS"), GopherImage, goscope_js.JsUtils(), goscope_js.DashboardJs())
}
