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
	result := make([]string, 0)

	for _, file := range files {
		f, _ := os.Open(file)
		content, _ := ioutil.ReadAll(f)
		f.Close()

		lines := strings.Split(string(content), "\n")
		var (
			fileNum, lineNum              string
			fileNames, entireLine, invert bool
		)

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
				case "-x":
					entireLine = true
				case "-v":
					invert = true
				}
			}

			if len(files) > 1 {
				fileNum = file + ":"
			}

			if strings.Contains(line, pattern) || invert {
				if fileNames {
					var found bool
					for r := range result {
						if result[r] == file {
							found = true
							break
						}
					}
					if !found {
						result = append(result, file)
					}
				} else if entireLine && !invert {
					if line == pattern {
						result = append(result, fileNum+lineNum+lines[i])
					}
				} else if invert {
					for ii := range lines {
						if lines[ii] == "" {
							continue
						}
						if !strings.Contains(lines[ii], pattern) {
							result = append(result, fileNum+lineNum+lines[ii])
						}
					}
					break
				} else {
					result = append(result, fileNum+lineNum+lines[i])
				}
			}
		}
	}

	return result
}
