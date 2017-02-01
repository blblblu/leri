package parsing

type block interface {
	mergeIfPossible(block block) bool
	serializeMd(lang string, sep string) string
}
