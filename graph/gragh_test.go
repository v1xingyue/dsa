package gragh

import (
	"container/list"
	"testing"
)

func TestGraghInit(t *testing.T) {
	g := NewGragh(32)
	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(4, 3)
	g.addEdge(3, 3)
	t.Log(g)
}

func TestQueuePass(t *testing.T) {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(32 + i)
	}
	l0 := l.Len()
	t.Log("l", l, l0)
	b := l
	x := b.Front()
	t.Log("front x : ", x)
	b.Remove(x)
	b0 := b.Len()
	t.Log("b", b, b0)
	t.Log("l", l, l0)
}

func TestGraghBfs(t *testing.T) {
	g := NewGragh(32)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(2, 3)
	g.addEdge(2, 5)
	g.addEdge(4, 5)
	g.addEdge(3, 5)
	g.addEdge(5, 10)
	g.bfs(1, 10)
}
