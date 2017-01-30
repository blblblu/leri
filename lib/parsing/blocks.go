package parsing

type Block interface {
	AddChunk(chunk string)
}
