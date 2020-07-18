package goscope_templates

import "fmt"

func NavbarComponent(selected string) string {
	navlinkKeys := []string {
		"REQUESTS", "LOGS",
	}
	navLinks := []string{
		"<a class=\"%s\" href=\"/goscope/\"><h3 style=\"margin: 1.2em;\" class=\"font-l\">Requests</h3></a>",
		"<a class=\"%s\" href=\"/goscope/logs\"><h3 style=\"margin: 1.2em;\" class=\"font-l\">Logs</h3></a>",
	}
	html := `
	<div class="flex m-2 p-2">
		<img width="100px" src="https://pro-warehouse-res.cloudinary.com/image/upload/v1595015193/git-repositories/goscope/application/q1btajmprnufw3ljr8re.png">
		<h3 class="font-l" style="margin: 1.2em;">{{.APPLICATION_NAME}}</h3>
		%s
	</div>
	`
	var generatedLinks string
	for i, s := range navlinkKeys {
		if s == selected {
			generatedLinks += fmt.Sprintf(navLinks[i], "active-navbar-link")
		} else {
			generatedLinks += fmt.Sprintf(navLinks[i], "navbar-link")
		}
	}
	return MinifyHtml(fmt.Sprintf(html, generatedLinks))
}
