package parsing

import (
	"fmt"
	"strings"
)

type sourceBlock struct {
	chunks []string
}

func (b *sourceBlock) addChunk(chunk string) {
	b.chunks = append(b.chunks, chunk)
}

func (b *sourceBlock) serialize() string {
	sourceCode := strings.Join(b.chunks, "\n")
	return fmt.Sprintf("```\n%s\n```", sourceCode)
}
