package hashids

type HashType int

const (
	HashTypeUser HashType = iota
	HashTypeWebsite
	HashTypeComment
)
