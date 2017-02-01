package parsing

import "strings"

type documentationBlock struct {
	chunks []string
}

func (b *documentationBlock) mergeIfPossible(block block) bool {
	switch block.(type) {
	case *documentationBlock:
		b.chunks = append(b.chunks, block.(*documentationBlock).chunks...)
		return true
	}
	return false
}

func (b *documentationBlock) serializeMd(_ string, sep string) string {
	return strings.Join(b.chunks, sep)
}

func newDocumentationBlock(chunk string) documentationBlock {
	return documentationBlock{
		chunks: []string{chunk},
	}
}
