package parsing

type block interface {
	addChunk(chunk string)
	serialize() string
}
