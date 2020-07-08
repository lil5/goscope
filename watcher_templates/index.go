package watcher_templates

import (
	"bitbucket.org/prowarehouse-nl/gohttpwatcher/watcher_css"
	"bitbucket.org/prowarehouse-nl/gohttpwatcher/watcher_js"
	"fmt"
)

var IndexTemplate = fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go HTTP Watcher</title>
    <style>%s</style>
	<style>%s</style>
	<style>%s</style>
</head>
<body>
<div class="m-3 p-3 text-center">
    <h1 class="font-xl m-2 text-center">{{.APPLICATION_NAME}} - Go HTTP Watcher</h1>
    <div class="m-3 p-3">
        <table id="request-table" class="p-6 md:w-2/3 lg:w-2/3" style="line-height: 1.6em; margin: 0 auto;">
            <thead>Incoming Requests</thead>
            <tr>
                <th class="custom-td">Verb</th>
                <th class="custom-td text-center">Path</th>
                <th class="custom-td">Happened</th>
                <th class="custom-td"></th>
            </tr>
        </table>
        <div>
            <button id="requests-prev-page" class="paginate-button"><span class="font-4xl">&#8592;</span> prev</button>
            &nbsp;
            <button id="requests-next-page" class="paginate-button">next <span class="font-4xl">&#8594;</span></button>
        </div>
    </div>
</div>
<script>%s</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</body>
</html>`,
	watcher_css.RaisinCss, watcher_css.HighlightTheme, watcher_css.WatcherStyles, watcher_js.DashboardJs)
