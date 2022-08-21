package graph

type Relative int

const (
	Parents Relative = iota
	Children
	Spouse
	Cousins
)
