package gospec

import (
	"strings"
	"unicode"
)

// Identifier represents a Go identifier in a variety of common
// case conventions.
type Identifier struct {
	Camel   string
	Kebab   string
	Natural string
	Package string
	Pascal  string
	Snake   string
	Source  string
}

// NewIdentifier parses the supplied string into an Identifier.
// Capital letters, whitespace, and punctuation are treated as
// word boundaries.
func NewIdentifier(s string) Identifier {
	words := parse(s)
	return Identifier{
		Camel:   camel(words),
		Kebab:   kebab(words),
		Natural: natural(words),
		Package: packge(words),
		Pascal:  pascal(words),
		Snake:   snake(words),
		Source:  s,
	}
}

// parser manages state for parsing an identifier.
type parser struct {
	word  strings.Builder
	words []string
}

// shift adds the current word to the rolling set of words.
// This is a no-op if the current word is empty.
func (p *parser) shift() {
	if p.word.Len() > 0 {
		p.words = append(p.words, p.word.String())
		p.word.Reset()
	}
}

// write adds the given rune to the current word.
func (p *parser) write(r rune) {
	p.word.WriteRune(r)
}

// parse the given string into a slice of words.
func parse(s string) []string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil
	}
	p := new(parser)
	for _, r := range s {
		if isUpper(r) {
			r = unicode.ToLower(r)
			p.shift()
		}
		if isLower(r) {
			p.write(r)
			continue
		}
		p.shift()
	}
	p.shift()
	return p.words
}

// isUpper return strue if the given rune represents an
// uppercase character.
func isUpper(r rune) bool {
	return unicode.IsUpper(r) || unicode.IsTitle(r)
}

// isLower return strue if the given rune represents a
// lowercase character.
func isLower(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

// camel case variant of the identifier.
func camel(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(words[0])
	for i := 1; i < len(words); i++ {
		sb.WriteString(title(words[i]))
	}
	return sb.String()
}

// kebab case variant of the identifier.
func kebab(words []string) string {
	return strings.Join(words, "-")
}

// natural is the space-separated variant of the identifier.
func natural(words []string) string {
	return strings.Join(words, " ")
}

// packge is the all-lowercase variant of the identifier.
func packge(words []string) string {
	return strings.Join(words, "")
}

// pascal case variant of the identifier.
func pascal(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, word := range words {
		sb.WriteString(title(word))
	}
	return sb.String()
}

// snake case variant of the identifier.
func snake(words []string) string {
	return strings.Join(words, "_")
}

// title returns the title-equivalent representation of the
// given string.
func title(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if i == 0 {
			sb.WriteRune(unicode.ToTitle(r))
			continue
		}
		sb.WriteRune(r)
	}
	return sb.String()
}
