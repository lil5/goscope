package goscope_templates

import (
	"github.com/averageflow/goscope/goscope_css"
	"github.com/averageflow/goscope/goscope_js"
	"fmt"
)

func LogsView() string {
	const template = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.APPLICATION_NAME}} Logs - GoScope</title>
    <link href="https://cdn.jsdelivr.net/gh/tretapey/raisincss@1.1.0/raisin.min.css" rel="stylesheet"/>
	<link href="https://fonts.googleapis.com/css2?family=Manrope&display=swap" rel="stylesheet"> 
	<link href="https://fonts.googleapis.com/css2?family=Fira+Code&display=swap" rel="stylesheet"> 
    <style>%s</style>
    <style>%s</style>
	<link rel="apple-touch-icon" sizes="180x180" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595014987/git-repositories/goscope/application/lrhheuzhlvgtsxyayru3.png">
	<link rel="icon" type="image/png" sizes="32x32" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015010/git-repositories/goscope/application/blaaguzlyw9s5x4jbhgq.png">
	<link rel="icon" type="image/png" sizes="16x16" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015010/git-repositories/goscope/application/bchsism3hubielueb6xk.png">
	<link rel="mask-icon" href="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015042/git-repositories/goscope/application/wswxoxpwirad8udbkwz8.svg" color="#5bbad5">
	<meta name="msapplication-TileColor" content="#ffffff">
	<meta name="theme-color" content="#ffffff">
</head>
<body>
%s
<div class="m-1 p-1 text-center">
	<table id="log-table" style="line-height: 1.6em; margin: 0 auto;">
	</table>
	<div>
		<button id="logs-prev-page" class="paginate-button"><span class="font-4xl">&#8592;</span> prev</button>
		&nbsp;
		<button id="logs-next-page" class="paginate-button">next <span class="font-4xl">&#8594;</span></button>
	</div>
    %s
</div>
<script>%s</script>
<script>%s</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</body>
</html>`
	return fmt.Sprintf(template, goscope_css.HighlightTheme(), goscope_css.WatcherStyles(), NavbarComponent("LOGS"), FooterText, goscope_js.JsUtils(), goscope_js.LogsJs())
}
