package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push('a')
	s.Push('b')
	s.Push('c')
	t.Log(string(s.Pop()))
	t.Log(string(s.Pop()))
	t.Log(string(s.Pop()))
	t.Log(s.IsEmpty())
}
