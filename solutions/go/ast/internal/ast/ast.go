package ast

type Node interface {
	ast()
}

type Number struct {
	Value string
}

func (Number) ast() {}

type Op int

const (
	Add Op = iota + 1
	Sub
	Mul
	Div
)

type Operator struct {
	L, R Node
	Op   Op
}

func (Operator) ast() {}
