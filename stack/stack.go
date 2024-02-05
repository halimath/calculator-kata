package stack

type Stack[T any] []T

func (s *Stack[T]) Empty() bool { return len(*s) == 0 }

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Peek() *T {
	if len(*s) == 0 {
		panic("empty stack")
	}

	return &(*s)[len(*s)-1]
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("empty stack")
	}

	t := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return t
}

func (s *Stack[T]) Shift() T {
	if len(*s) == 0 {
		panic("empty stack")
	}

	t := (*s)[0]

	*s = (*s)[1:len(*s)]

	return t
}
