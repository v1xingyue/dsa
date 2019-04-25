package binarytree

import (
	"dsa/queue"
)

/*
几种遍历方式，每种遍历方式，为方便代码编写，传递一个回调，用来对数据，做后续处理
*/
type LoopHandler func(x vtype)

//前序遍历
func PreLoopTree(tree *node, fn LoopHandler) {
	if tree == nil {
		return
	}
	if fn != nil {
		fn(tree.data)
	}
	PreLoopTree(tree.left, fn)
	PreLoopTree(tree.right, fn)
}

//后续遍历
func PostLoopTree(tree *node, fn LoopHandler) {
	if tree == nil {
		return
	}
	PostLoopTree(tree.left, fn)
	PostLoopTree(tree.right, fn)
	if fn != nil {
		fn(tree.data)
	}
}

//中序遍历
func MidLoopTree(tree *node, fn LoopHandler) {
	if tree == nil {
		return
	}
	MidLoopTree(tree.left, fn)
	if fn != nil {
		fn(tree.data)
	}
	MidLoopTree(tree.right, fn)
}

//按层遍历
func LevelLoopTree(tree *node, fn LoopHandler) {
	lq := queue.NewListQueue(256)
	if tree != nil {
		lq.Insert(tree)
	}
	for !lq.Empty() {
		root, _ := lq.Shift()
		_root := root.(*node)
		fn(_root.data)
		if _root.left != nil {
			lq.Insert(_root.left)
		}
		if _root.right != nil {
			lq.Insert(_root.right)
		}
	}
}
