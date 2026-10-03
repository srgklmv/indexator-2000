package parser

import (
	"context"
)

type Parser interface {
	ParseFile(ctx context.Context, filename string) (File, error)
	AddIndex(ctx context.Context, filename string, lines []string) error
}
