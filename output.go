package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
)

// output writes the formatted data to the given Writer.
func output(w io.Writer, example string, filedata []byte, matches []*Match) error {
	// Convert example string into runes.
	exrunes := []rune(example)

	// Loop through the file records.
	bfd := bytes.NewBuffer(filedata)
	scanner := bufio.NewScanner(bfd)
	for scanner.Scan() {
		// Store the current line.
		line := scanner.Text()

		// Format this line.
		var result []rune
		for i, m := range matches {
			// Add data before first match.
			if i == 0 && m.start != 0 {
				result = append(result, exrunes[0:m.start]...)
			}

			// Handle appending data between this match and the prior match.
			if i > 0 && m.start > matches[i-1].end+1 {
				result = append(result, exrunes[matches[i-1].end+1:m.start]...)
			}

			// Create regexp for this match and run against line to get data.
			r := regexp.MustCompile(m.re)
			submatch := r.FindStringSubmatch(line)

			// Append found data to result.
			result = append(result, []rune(submatch[m.group])...)

			// If this is the last match.
			if i == len(matches)-1 && len(example) > m.end+1 {
				result = append(result, exrunes[m.end+1:]...)
			}
		}

		fmt.Fprintf(w, "%s\n", string(result))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
