package main

// Match holds the data for each bracket match.
type Match struct {
	Value   string
	runes   []rune
	start   int
	end     int
	prevr   rune
	nextr   rune
	re      string
	group   int
	isAtEnd bool
}
