package binarytree

import (
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
