package main

import (
	"flag"

	"github.com/srgklmv/indexator-2000/internal/parser"
)

var (
	file = flag.String("file", "", "path to file for indexation")
)

func main() {
	_ = parser.New()
}
