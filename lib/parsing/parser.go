package parsing

import (
	"regexp"
	"strings"
)

type Parser interface {
	GenerateMd(source string) string
}

type SimpleRegexParser struct {
	language   string // the language identifier that will be later be used in the markdown code block, e.g. "```go\n...\n````"
	chunkDelim string
	docRegex   string
}

func (p *SimpleRegexParser) GenerateMd(source string) string {
	document := document{}

	chunks := strings.Split(source, p.chunkDelim)

	docRe := regexp.MustCompile(p.docRegex)

	for _, chunk := range chunks {
		if match := docRe.FindStringSubmatch(chunk); match != nil {
			documentation := newDocumentationBlock(match[1])
			document.addBlock(&documentation)
			continue
		}
		source := newSourceBlock(chunk)
		document.addBlock(&source)
	}

	return document.serializeMd(p.language, p.chunkDelim)
}

func NewSimpleGoParser() SimpleRegexParser {
	return SimpleRegexParser{
		language:   "go",
		chunkDelim: "\n",
		docRegex:   `^\s*\/\/\ (.*)$\s*`,
	}
}

// NewSimpleGoDoccommentParser creates a similar Parser like NewSimpleGoParser, but only `///` comments will be used as documentation
func NewSimpleGoDoccommentParser() SimpleRegexParser {
	return SimpleRegexParser{
		language:   "go",
		chunkDelim: "\n",
		docRegex:   `^\s*\/\/\/(.*)\s*$`,
	}
}

func NewSimpleSasParser() SimpleRegexParser {
	return SimpleRegexParser{
		language:   "sas",
		chunkDelim: "\n\n",
		docRegex:   `^\s*\/\*\s*((.*|\n)*)\*\/\s*$`,
	}
}
