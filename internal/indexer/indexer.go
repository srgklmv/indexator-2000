package indexer

import "context"

type indexer struct{}

func New() Indexer {
	return &indexer{}
}

func (i *indexer) Index(ctx context.Context, headings []string) []string {
	panic("not implemented")
}
