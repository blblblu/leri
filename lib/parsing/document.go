package parsing

import "strings"

type document struct {
	blocks []block
}

func (d *document) serialize() string {
	mdBlocks := []string{}
	for _, b := range d.blocks {
		mdBlocks = append(mdBlocks, b.serialize())
	}
	return strings.Join(mdBlocks, "\n")
}
