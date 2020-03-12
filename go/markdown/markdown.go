// Package markdown implements parsing a given string with Markdown syntax
// and returns the associated HTML
package markdown

import (
	"strconv"
	"strings"
)

// Render translates markdown to HTML.
func Render(markdown string) string {
	var header, list int

	html := strings.Builder{}
	html.Grow(len(markdown))

	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	//r := strings.NewReplacer("__", "<strong>", "__", "</strong>", "_", "<em>", "_", "</em>")
	//markdown = r.Replace(markdown)

	for pos := 0; pos < len(markdown); {
		char := markdown[pos]

		switch char {
		case '#':
			for ; char == '#'; char = markdown[pos] {
				header++
				pos++
			}
			html = *writeHeaderSuffix(&html, header, true)
			pos++
		case '*':
			if list == 0 {
				html.WriteString("<ul>")
			}

			html.WriteString("<li>")
			list++
			pos += 2
		case '\n':
			if list > 0 {
				html.WriteString("</li>")
			}

			if header > 0 {
				html = *writeHeaderSuffix(&html, header, false)
				header = 0
			}

			pos++
		default:
			html.WriteByte(char)
			pos++
		}

	}

	switch {
	case header > 0:
		html = *writeHeaderSuffix(&html, header, false)
	case list > 0:
		html.WriteString("</li></ul>")
	default:
		html.WriteString("</p>")
		return "<p>" + html.String()
	}

	return html.String()
}

func writeHeaderSuffix(sb *strings.Builder, header int, open bool) *strings.Builder {
	h := strconv.Itoa(header)
	sb.WriteByte('<')
	if !open {
		sb.WriteByte('/')
	}
	sb.WriteByte('h')
	sb.WriteString(h)
	sb.WriteByte('>')

	return sb
}
