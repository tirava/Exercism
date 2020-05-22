// Package grep searches files for lines matching a regular expression pattern.
package grep

import (
	"io/ioutil"
	"os"
	"strings"
)

// Search returns found strings.
func Search(pattern string, flags, files []string) []string {
	var result []string

	for _, file := range files {
		f, _ := os.Open(file)
		content, _ := ioutil.ReadAll(f)
		f.Close()
		//fmt.Println(string(content))

		lines := strings.Split(string(content), "\n")
		for i := range lines {
			if strings.Contains(lines[i], pattern) {
				result = append(result, lines[i])
			}
		}
	}

	return result
}

//-n Print the line numbers of each matching line.
//-l Print only the names of files that contain at least one matching line.
//-i Match line using a case-insensitive comparison.
//-v Invert the program -- collect all lines that fail to match the pattern.
//-x Only match entire lines, instead of lines that contain a match.
