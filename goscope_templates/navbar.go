package goscope_templates

func NavbarComponent() string {
	html := `
	<div class="m-2 p-2"><h3 class="font-l">{{.APPLICATION_NAME}}</h3><a href="/goscope/">Requests</a><a href="/goscope/logs">Logs</div>
	`
	return MinifyHtml(html)
}
