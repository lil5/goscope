package goscope_templates

import (
	"bitbucket.org/prowarehouse-nl/goscope/goscope_css"
	"bitbucket.org/prowarehouse-nl/goscope/goscope_js"
	"fmt"
)

var LogsTemplate = fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.APPLICATION_NAME}} Logs - GoScope</title>
    <style>%s</style>
    <style>%s</style>
    <style>%s</style>
</head>
<body>
<div class="m-1 p-1 text-center">
	<div class="m-3 p-3">
		<table id="log-table" class="p-6 md:w-2/3 lg:w-2/3" style="line-height: 1.6em; margin: 0 auto;">
		</table>
		<div>
			<button id="log-prev-page" class="paginate-button"><span class="font-4xl">&#8592;</span> prev</button>
			&nbsp;
			<button id="log-next-page" class="paginate-button">next <span class="font-4xl">&#8594;</span></button>
		</div>
	</div>
    <img src="%s">
</div>
<script>%s</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</body>
</html>`,
	goscope_css.RaisinCss, goscope_css.HighlightTheme, goscope_css.WatcherStyles, GopherImage, goscope_js.LogsJs)
