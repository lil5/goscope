package goscope_templates

func NavbarComponent() string {
	html := `
	<div class="flex m-2 p-2">
		<h3 class="font-l">{{.APPLICATION_NAME}}</h3>
		<a href="/goscope/"><h3 class="font-l">Requests</h3></a>
		<a href="/goscope/logs"><h3 class="font-l">Logs</h3></a>
	</div>
	`
	return MinifyHtml(html)
}
