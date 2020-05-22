// Package grep searches files for lines matching a regular expression pattern.
package grep

import (
	"io/ioutil"
	"os"
	"strconv"
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
		var lineNum string
		var fileNames bool

		for i := range lines {
			line := lines[i]
			for j := range flags {
				switch flags[j] {
				case "-n":
					lineNum = strconv.Itoa(i+1) + ":"
				case "-i":
					pattern = strings.ToLower(pattern)
					line = strings.ToLower(lines[i])
				case "-l":
					fileNames = true
				}
			}
			if strings.Contains(line, pattern) {
				if fileNames {
					result = append(result, file)
				} else {
					result = append(result, lineNum+lines[i])
				}
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
