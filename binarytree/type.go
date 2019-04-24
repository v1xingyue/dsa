package binarytree

type vtype int

/*
二叉树基本类型定义，及常用方法添加，
左边加入节点，
右边加入节点，
左右两边同时加入数据
同时返回加入后新生成的节点
*/
type node struct {
	data  vtype
	left  *node
	right *node
}

func Node(val vtype) *node {
	return &node{val, nil, nil}
}

func (self *node) addLeft(l vtype) *node {
	self.left = Node(l)
	return self.left
}

func (self *node) addRight(r vtype) *node {
	self.right = Node(r)
	return self.right
}

func (self *node) addBoth(l, r vtype) (*node, *node) {
	self.left = Node(l)
	self.right = Node(r)
	return self.left, self.right
}
