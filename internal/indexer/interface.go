package indexer

import "context"

type Indexer interface {
	Index(ctx context.Context, headings []string) []string
}
