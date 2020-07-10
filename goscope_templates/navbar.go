package goscope_templates

func NavbarComponent() string {
	html := `
	<div class="flex m-2 p-2">
		<h3 class="font-l" style="margin: 1.2em;">{{.APPLICATION_NAME}}</h3>
		<a class="navbar-link" href="/goscope/"><h3 style="margin: 1.2em;" class="font-l">Requests</h3></a>
		<a class="navbar-link" href="/goscope/logs"><h3 style="margin: 1.2em;" class="font-l">Logs</h3></a>
	</div>
	`
	return MinifyHtml(html)
}
