package stack

import (
	"testing"

	"github.com/halimath/expect"
	"github.com/halimath/expect/is"
)

func TestStack(t *testing.T) {
	var s Stack[int]

	expect.That(t, is.EqualTo(s.Empty(), true))

	s.Push(1)
	expect.That(t, is.EqualTo(s.Empty(), false))

	s.Push(2)
	s.Push(3)
	expect.That(t, is.EqualTo(len(s), 3))

	v := s.Peek()
	expect.That(t, is.EqualTo(v, 3))
	expect.That(t, is.EqualTo(len(s), 3))

	v = s.Pop()
	expect.That(t, is.EqualTo(v, 3))
	expect.That(t, is.EqualTo(len(s), 2))

	v = s.Shift()
	expect.That(t, is.EqualTo(v, 1))
	expect.That(t, is.EqualTo(len(s), 1))

	v = s.Pop()
	expect.That(t, is.EqualTo(v, 2))
	expect.That(t, is.EqualTo(len(s), 0))

	s.Push(1)
	v = s.Shift()
	expect.That(t, is.EqualTo(v, 1))
	expect.That(t, is.EqualTo(len(s), 0))
}
