package goscope_templates

import (
	"bitbucket.org/prowarehouse-nl/goscope/goscope_css"
	"bitbucket.org/prowarehouse-nl/goscope/goscope_js"
	"fmt"
)

func LogView() string {
	const template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Request at {{.TIME}} - {{.APPLICATION_NAME}}</title>
    <link href="https://cdn.jsdelivr.net/gh/tretapey/raisincss@1.1.0/raisin.min.css" rel="stylesheet"/>
    <style>%s</style>
    <style>%s</style>
</head>
<body>
<div class="m-3 p-3 text-center" style="line-height: 2em;">
	<div class="md:w-2/3 lg:w-2/3 text-left" style="margin:0 auto;">
		<h1 class="font-xl m-2">Logged at {{.TIME}} - {{.APPLICATION_NAME}} - GoScope</h1>
		<p>Time:
		<pre><code class="plaintext">{{.TIME}}</code></pre>
		</p>
		<p>Message:
		<pre><code class="plaintext">{{.MESSAGE}}</code></pre>
		</p>
	</div>
</div>
<script>%s</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
</body>
</html>`
	return fmt.Sprintf(template, goscope_css.HighlightTheme(), goscope_css.WatcherStyles(), goscope_js.LogJs())
}
