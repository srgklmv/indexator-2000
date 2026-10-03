package parser

import (
	"context"
)

type Parser interface {
	ParseFile(ctx context.Context, filename string) File
}
