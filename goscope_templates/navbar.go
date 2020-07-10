package goscope_templates

import "fmt"

func NavbarComponent(selected string) string {
	navLinks := map[string]string{
		"REQUESTS": "<a class=\"%s\" href=\"/goscope/\"><h3 style=\"margin: 1.2em;\" class=\"font-l\">Requests</h3></a>",
		"LOGS":     "<a class=\"%s\" href=\"/goscope/logs\"><h3 style=\"margin: 1.2em;\" class=\"font-l\">Logs</h3></a>",
	}
	html := `
	<div class="flex m-2 p-2">
		<h3 class="font-l" style="margin: 1.2em;">{{.APPLICATION_NAME}}</h3>
		%s
	</div>
	`
	var generatedLinks string
	for i, s := range navLinks {
		if i == selected {
			generatedLinks += fmt.Sprintf(s, "active-navbar-link")
		} else {
			generatedLinks += fmt.Sprintf(s, "navbar-link")
		}
	}
	return MinifyHtml(fmt.Sprintf(html, generatedLinks))
}
