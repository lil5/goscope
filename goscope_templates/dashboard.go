package goscope_templates

import (
	"github.com/averageflow/goscope/goscope_css"
	"github.com/averageflow/goscope/goscope_js"
	"fmt"
)

func DashboardView() string {
	const template = `
<!DOCTYPE html>
<html lang="en">
<head>
    %s
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
	return fmt.Sprintf(template, CommonHead, goscope_css.HighlightTheme(), goscope_css.WatcherStyles(), NavbarComponent("REQUESTS"), FooterText, goscope_js.JsUtils(), goscope_js.DashboardJs())
}
