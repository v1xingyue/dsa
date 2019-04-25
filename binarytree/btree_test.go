package binarytree

import (
	"dsa/queue"
	"log"
	"testing"
)

func simplePrint(v vtype) {
	log.Println(v)
}

func TestPreOrder(t *testing.T) {
	root := Node(3)
	l := root.addLeft(12)
	l.addLeft(15)
	l.addRight(18)

	r := root.addRight(32)
	r.addLeft(24)
	r.addRight(18)

	PreLoopTree(root, simplePrint)
}

func TestPostOrder(t *testing.T) {
	root := Node(3)
	l := root.addLeft(12)
	l.addLeft(15)
	l.addRight(18)

	r := root.addRight(32)
	r.addLeft(24)
	r.addRight(18)

	PostLoopTree(root, simplePrint)
}

func TestMidOrder(t *testing.T) {
	root := Node(3)
	l := root.addLeft(12)
	l.addLeft(15)
	l.addRight(18)

	r := root.addRight(32)
	r.addLeft(24)
	r.addRight(18)

	MidLoopTree(root, simplePrint)
}

func TestLevelOrder(t *testing.T) {
	root := Node(3)
	l := root.addLeft(12)
	l.addLeft(15)
	l.addRight(18)

	r := root.addRight(32)
	r.addLeft(24)
	r.addRight(18)

	LevelLoopTree(root, simplePrint)

}

func TestPushNilNode(t *testing.T) {
	lq := queue.NewListQueue(32)
	root := Node(3)
	t.Log(lq, "\n", root.left, "\n", root.right, "\n", root)
	lq.Insert(root)
	lq.Insert(root.left)
	lq.Insert(root.right)
	t.Log(lq, "\n", root.left, "\n", root.right, "\n", root)

	a, _ := lq.Shift()
	b, _ := lq.Shift()
	c, _ := lq.Shift()

	t.Log("a : ", a)
	t.Log("b : ", b, b == nil)
	t.Log("c : ", c, c == nil)
}

func TestInterfaceNil(t *testing.T) {
	// 包含两个指针， pt,pv  分别指向 值指针和类型指针
	// 所以就算赋值为nil ， 那么该 值依然不等于nil
	// 具体可以参看下边的代码
	var l interface{}
	var n *int
	n = nil
	l = n
	t.Log(l == nil)
}
