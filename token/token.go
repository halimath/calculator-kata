package token

type Type int

const (
	Number Type = iota + 1
	Add
	Sub
	Mul
	Div
	LParen
	RParen
)

type Token struct {
	Type  Type
	Value string

	// TODO: Add location infos
}

func IsOperator(t Token) bool {
	return t.Type == Add || t.Type == Sub || t.Type == Mul || t.Type == Div
}
