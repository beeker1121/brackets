package main

import (
	"bufio"
	"bytes"
	"regexp"
)

// parse parses the example and file data to determine match data.
func parse(example *string, filedata []byte) ([]*Match, error) {
	// Get first record from file.
	bfd := bytes.NewBuffer(filedata)
	scanner := bufio.NewScanner(bfd)
	scanner.Scan()
	record := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Get matches from example input.
	matches, err := getMatches(example)
	if err != nil {
		return nil, err
	}

	// Find the previous and next runes for each match.
	if err = findPrevNextRune(record, matches); err != nil {
		return nil, err
	}

	// Get the regexp pattern for each match.
	if err = getRegexp(record, matches); err != nil {
		return nil, err
	}

	return matches, nil
}

// getMatches parses the example brackets.
func getMatches(example *string) ([]*Match, error) {
	var matches []*Match
	var start int
	var startFound, endFound, inEscapeMode bool

	// Convert example string into runes.
	runes := []rune(*example)

	// Create variable to hold unescaped example runes.
	newRunes := make([]rune, 0)

	// Loop through the example runes.
	var idx int
	for _, r := range runes {
		// Handle escaping.
		if inEscapeMode {
			switch r {
			case rune('\\'):
				newRunes = append(newRunes, r)
				inEscapeMode = false
				idx++
				continue
			case rune('{'):
				newRunes = append(newRunes, r)
				inEscapeMode = false
				idx++
				continue
			case rune('}'):
				newRunes = append(newRunes, r)
				inEscapeMode = false
				idx++
				continue
			default:
				return nil, ErrUnknownEscapeSeq
			}
		} else if !inEscapeMode && r == rune('\\') {
			inEscapeMode = true
			continue
		}

		// If we found a starting bracket.
		if r == rune('{') {
			// Check if we have already found a starting bracket.
			if startFound {
				return nil, ErrMultiBrackets
			}
			startFound = true

			start = idx
		}

		// If we found an ending bracket.
		if r == rune('}') {
			// Check if we have a starting bracket.
			if !startFound {
				return nil, ErrMissingStartingBracket
			}
			startFound = false
			endFound = true

			// Create a new Match.
			match := &Match{
				Value: string(newRunes[start+1 : idx]),
				runes: newRunes[start+1 : idx],
				start: start,
				end:   idx,
			}

			// Append to matches.
			matches = append(matches, match)
		}

		newRunes = append(newRunes, r)
		idx++
	}

	// If an ending bracket was not found.
	if !endFound {
		return nil, ErrMissingEndingBracket
	}

	// Set given example to new, unescaped version.
	*example = string(newRunes)

	return matches, nil
}

// getMatchIndex finds the index of the given match value within
// the given record, using runes.
func getMatchIndex(record, value string) int {
	// Convert record and value into runes.
	recordr := []rune(record)
	valuer := []rune(value)

	// Loop through the record runes.
	for i := 0; i < len(recordr)-len(valuer)+1; i++ {
		match := true
		for j := 0; j < len(valuer); j++ {
			if recordr[i+j] != valuer[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}

// findPrevNextRune finds the previous and next rune in the
// given record for each match.
func findPrevNextRune(record string, matches []*Match) error {
	// Convert record string into runes.
	runes := []rune(record)

	// Loop through the matches.
	for _, m := range matches {
		// Get the index of this match in the record.
		idx := getMatchIndex(record, m.Value)
		if idx == -1 {
			return ErrMatchNotFound
		}

		if idx > 0 {
			m.prevr = runes[idx-1]
		}

		// If this match is at the end of the record.
		if idx+len(m.Value) == len(runes) {
			m.isAtEnd = true
		} else {
			m.nextr = runes[idx+len(m.Value)]
		}
	}

	return nil
}

// escape escapes regexp tokens.
func escape(token string) string {
	switch token {
	case ".":
		return "\\."
	case "^":
		return "\\^"
	case "$":
		return "\\$"
	case "[":
		return "\\["
	case "]":
		return "\\]"
	case "{":
		return "\\{"
	case "}":
		return "\\}"
	case "*":
		return "\\*"
	case "?":
		return "\\?"
	case "+":
		return "\\+"
	default:
		return token
	}
}

// getRegexp forms the regexp pattern for each match based
// on the given record.
func getRegexp(record string, matches []*Match) error {
	// Loop through the matches.
	for _, m := range matches {
		var idx int
		var re string
		var found bool

		// Start loop to form regexp pattern.
		for !found {
			var prev, next string

			// If previous rune is not null.
			if m.prevr != 0 {
				prev = escape(string(m.prevr))
			}

			// If this match is at the end of the record.
			if m.isAtEnd {
				next = "$"
			} else {
				next = escape(string(m.nextr))
			}

			// If first iteration.
			if idx == 0 {
				re = prev + "(.*?)" + next
			} else {
				re = prev + "(.*?)" + re
			}

			// Compile this regular expression and try to find exact
			// match in submatches.
			r := regexp.MustCompile(re)
			submatches := r.FindStringSubmatch(record)

			// Loop through submatches.
			var subIdx int
			for _, sm := range submatches {
				if sm == m.Value {
					m.re = re
					m.group = subIdx
					found = true
				}
				subIdx++
			}
			idx++
		}
	}

	return nil
}
