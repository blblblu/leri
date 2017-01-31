package parsing

import "strings"

type documentationBlock struct {
	chunks []string
}

func (b *documentationBlock) addChunk(chunk string) {
	b.chunks = append(b.chunks, chunk)
}

func (b *documentationBlock) serialize() string {
	return strings.Join(b.chunks, "\n")
}
