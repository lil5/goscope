package goscope_js

func LogJs() string {
	const script = `
document.addEventListener("DOMContentLoaded", function () {
    hljs.initHighlightingOnLoad();
});
`
	return MinifyJs(script)
}
