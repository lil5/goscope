package watcher_templates

const RequestTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Request at {{.TIME}} - {{.APPLICATION_NAME}}</title>
    <link rel="stylesheet" href="/css/raisin.css">
    <link rel="stylesheet" href="/css/watcher.css">
    <link rel="stylesheet" href="/css/highlight-theme.css">
</head>
<body>

<div class="m-3 p-3 text-center" style="line-height: 2em;">
    <h1 class="font-xl m-2">Request at {{.TIME}} - {{.APPLICATION_NAME}} - Go Watcher</h1>
    <div class="md:w-2/3 lg:w-2/3 text-left" style="margin:0 auto;">
        <p>Client IP Address:
        <pre><code class="plaintext">{{.CLIENT_IP}}</code></pre>
        </p>
        <p>Host:
        <pre><code class="plaintext">{{.HOST}}</code></pre>
        </p>
        <p>Verb:
        <pre><code class="plaintext">{{.METHOD}}</code></pre>
        </p>
        <p>Path:
        <pre><code class="plaintext">{{.PATH}}</code></pre>
        </p>
        <p>Url:
        <pre><code class="plaintext">{{.URL}}</code></pre>
        </p>
        <p>Referrer:
        <pre><code class="plaintext">{{.REFERRER}}</code></pre>
        </p>
        <p>Time:
        <pre><code class="plaintext">{{.TIME}}</code></pre>
        </p>
        <p>Uid:
        <pre><code class="plaintext">{{.UID}}</code></pre>
        </p>
        <p>User Agent:
        <pre><code class="plaintext">{{.USER_AGENT}}</code></pre>
        </p>
        <p>Headers:
        <pre><code class="language-json">{{.HEADERS}}</code></pre>
        </p>
        <p>Body:
        <pre><code class="language-json">{{.BODY}}</code></pre>
        </p>
    </div>
</div>
<script src="/js/request.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
</body>
</html>
`