package parser

import "regexp"

var (
	levelOneHeadingRegexp = regexp.MustCompile("^#[^#]*$")
)

const (
	cell rune = '#'
)

type parser struct{}

func New() *parser {
	return &parser{}
}

func (p *parser) IsLevelHeading(level int, line string) bool {
	depth := p.GetHeadingDepth(line)

	return level == depth
}

func (p *parser) GetHeadingDepth(line string) (depth int) {
	for symbol := range line {
		if symbol != int(cell) {
			return depth
		}
		depth++
	}

	return depth
}
