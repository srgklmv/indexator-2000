package main

import (
	"context"
	"flag"

	"github.com/srgklmv/indexator-2000/internal/indexer"
	"github.com/srgklmv/indexator-2000/internal/parser"
)

var (
	_ = flag.String("file", "", "path to file for indexation")
)

const (
	defaultFilename = "README.md"
)

// todo: move logic to app module
func main() {
	ctx := context.Background()

	p := parser.New()
	i := indexer.New()

	data, err := p.ParseFile(ctx, defaultFilename)
	if err != nil {
		panic(err)
	}

	index := i.Index(ctx, data.Headings)

	err = p.AddIndex(ctx, defaultFilename, index)
	if err != nil {
		panic(err)
	}
}
