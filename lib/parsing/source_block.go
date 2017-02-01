package parsing

import (
	"fmt"
	"strings"
)

type sourceBlock struct {
	chunks []string
}

func (b *sourceBlock) mergeIfPossible(block block) bool {
	switch block.(type) {
	case *sourceBlock:
		b.chunks = append(b.chunks, block.(*sourceBlock).chunks...)
		return true
	}
	return false
}

func (b *sourceBlock) serializeMd(lang string, sep string) string {
	sourceCode := strings.Join(b.chunks, sep)
	return fmt.Sprintf("```%s\n%s\n```", lang, sourceCode)
}

func newSourceBlock(chunk string) sourceBlock {
	return sourceBlock{
		chunks: []string{chunk},
	}
}
