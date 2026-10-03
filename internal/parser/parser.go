package parser

import "regexp"

var (
	levelOneHeadingRegexp = regexp.MustCompile("^#[^#]*$")
)

const (
	cell rune = '#'
)

type parser struct{}

func New() Parser {
	return &parser{}
}
