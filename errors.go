package main

import "errors"

var (
	// ErrDirectory is returned when the file type trying
	// to be opened is a directory.
	ErrDirectory = errors.New("File type cannot be a directory")

	// ErrEmpty is returned when the file is empty.
	ErrEmpty = errors.New("File is empty")

	// ErrMultiBrackets is returned when the example input given
	// contains multiple starting brackets without an ending bracket.
	ErrMultiBrackets = errors.New("Multiple starting brackets without ending bracket")

	// ErrMissingStartingBracket is returned when the example input
	// given has an ending bracket, but is missing a starting
	// bracket.
	ErrMissingStartingBracket = errors.New("Missing starting bracket")

	// ErrMissingEndingBracket is returned when the example input
	// given has a starting bracket, but is missing an ending
	// bracket.
	ErrMissingEndingBracket = errors.New("Missing ending bracket")

	// ErrMatchNotFound is returned when a given match could not
	// be found in a source record of the file.
	ErrMatchNotFound = errors.New("Could not find match in record")
)
