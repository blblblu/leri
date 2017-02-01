package parsing

import "strings"

type document struct {
	blocks []block
}

func (d *document) serializeMd(lang string, sep string) string {
	mdBlocks := []string{}
	for _, b := range d.blocks {
		mdBlocks = append(mdBlocks, b.serializeMd(lang, sep))
	}
	return strings.Join(mdBlocks, "\n\n")
}

func (d *document) addBlock(block block) {
	if d.blocks == nil {
		d.blocks = append(d.blocks, block)
		return
	}

	if ok := d.blocks[len(d.blocks)-1].mergeIfPossible(block); !ok {
		d.blocks = append(d.blocks, block)
	}
}
