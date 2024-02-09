// Package stack provides a generic implementation of a stack abstract datastructure.
package stack

// Stack implements a stack of elements of type T. The default value for Stack is ready for use, though empty.
type Stack[T any] []T

// Empty returns whether s is empty, i.e. contains no element.
func (s *Stack[T]) Empty() bool { return len(*s) == 0 }

// Push pushes v on top of s.
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Peek returns the top element of s without removing it. It panics if s is empty.
func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		panic("empty stack")
	}

	return (*s)[len(*s)-1]
}

// Pop removes the top element from s and returns it. It panics if s is empty.
func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("empty stack")
	}

	t := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return t
}

// Shift removes the bottom-most element from s and returns it. It panics if s is empty.
func (s *Stack[T]) Shift() T {
	if len(*s) == 0 {
		panic("empty stack")
	}

	t := (*s)[0]

	*s = (*s)[1:len(*s)]

	return t
}
