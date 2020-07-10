package goscope_templates

import (
	"bitbucket.org/prowarehouse-nl/goscope/goscope_css"
	"bitbucket.org/prowarehouse-nl/goscope/goscope_js"
	"fmt"
)

func RequestView() string {
	const template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Request at {{.REQUEST_TIME}} - {{.APPLICATION_NAME}}</title>
    <link href="https://cdn.jsdelivr.net/gh/tretapey/raisincss@1.1.0/raisin.min.css" rel="stylesheet"/>
	<link href="https://fonts.googleapis.com/css2?family=Manrope&display=swap" rel="stylesheet"> 
	<link href="https://fonts.googleapis.com/css2?family=Fira+Code&display=swap" rel="stylesheet"> 
    <style>%s</style>
    <style>%s</style>
</head>
<body>
<div class="m-3 p-3 text-center">
	  <button class="tab-button p-4 font-xl" onclick="openTab('request-tab')">Request</button>
	  <button class="tab-button p-4 font-xl" onclick="openTab('response-tab')">Response</button>
</div> 
<div class="m-3 p-3 text-center" style="line-height: 2em;">
    <div id="request-tab" class="view-tab md:w-2/3 lg:w-2/3 text-left" style="margin:0 auto;">
		<h1 class="font-xl m-2">Request at {{.REQUEST_TIME}} - {{.APPLICATION_NAME}} - GoScope</h1>
        <p>Client IP Address:
        <pre><code class="plaintext">{{.REQUEST_CLIENT_IP}}</code></pre>
        </p>
        <p>Host:
        <pre><code class="plaintext">{{.REQUEST_HOST}}</code></pre>
        </p>
        <p>Verb:
        <pre><code class="plaintext">{{.REQUEST_METHOD}}</code></pre>
        </p>
        <p>Path:
        <pre><code class="plaintext">{{.REQUEST_PATH}}</code></pre>
        </p>
        <p>Url:
        <pre><code class="plaintext">{{.REQUEST_URL}}</code></pre>
        </p>
        <p>Referrer:
        <pre><code class="plaintext">{{.REQUEST_REFERRER}}</code></pre>
        </p>
        <p>Time:
        <pre><code class="plaintext">{{.REQUEST_TIME}}</code></pre>
        </p>
        <p>Uid:
        <pre><code class="plaintext">{{.REQUEST_UID}}</code></pre>
        </p>
        <p>User Agent:
        <pre><code class="plaintext">{{.REQUEST_USER_AGENT}}</code></pre>
        </p>
        <p>Headers:
        <pre><code class="language-json">{{.REQUEST_HEADERS}}</code></pre>
        </p>
        <p>Body:
        <pre><code class="language-json">{{.REQUEST_BODY}}</code></pre>
        </p>
    </div>
	<div id="response-tab" class="view-tab md:w-2/3 lg:w-2/3 text-left" style="margin:0 auto;">
		<h1 class="font-xl m-2">Response at {{.RESPONSE_TIME}}</h1>
		<p>Client IP Address:
		<pre><code class="plaintext">{{.RESPONSE_CLIENT_IP}}</code></pre>
		</p>
		<p>Status:
		<pre><code class="plaintext">{{.RESPONSE_STATUS}}</code></pre>
		</p>
		<p>Path:
		<pre><code class="plaintext">{{.RESPONSE_PATH}}</code></pre>
		</p>
		<p>Time:
		<pre><code class="plaintext">{{.RESPONSE_TIME}}</code></pre>
		</p>
		<p>Uid:
		<pre><code class="plaintext">{{.RESPONSE_UID}}</code></pre>
		</p>
		<p>Size:
		<pre><code class="plaintext">{{.RESPONSE_SIZE}}</code></pre>
		</p>
		<p>Headers:
		<pre><code class="language-json">{{.RESPONSE_HEADERS}}</code></pre>
		<p>Body:
		<pre><code class="language-json">{{.RESPONSE_BODY}}</code></pre>
		</p>
	</div>
</div>
<script>%s</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
</body>
</html>`
	return fmt.Sprintf(template, goscope_css.HighlightTheme(), goscope_css.WatcherStyles(), goscope_js.RequestJs())
}
