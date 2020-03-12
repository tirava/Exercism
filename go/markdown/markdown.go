// Package markdown implements parsing a given string with Markdown syntax
// and returns the associated HTML
package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML.
func Render(markdown string) string {
	var header, pos, list int
	var html string

	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	//r := strings.NewReplacer("__", "<strong>", "__", "</strong>", "_", "<em>", "_", "</em>")
	//markdown = r.Replace(markdown)

	for {
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
			pos++
			continue
		}
		if char == '*' {
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos += 2
			continue
		}
		if char == '\n' {
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
			pos++
			continue
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"

}
