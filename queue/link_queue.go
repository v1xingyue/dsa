package binarytree

import (
	_ "log"
)

/*
链式实现一个队列系统
*/
type linknode struct {
	data vtype
	next *linknode
}

type linkqueue struct {
	head *linknode
	last *linknode
	max  int
	size int
}

func NewListQueue(max int) *linkqueue {
	m := new(linkqueue)
	m.max = max
	m.size = 0
	m.head = nil
	m.last = nil
	return m
}

func (self *linkqueue) Insert(v vtype) bool {
	if self.size >= self.max {
		return false
	}
	if self.last == nil {
		self.last = &linknode{v, nil}
		self.head = self.last
	} else {
		self.last.next = &linknode{v, nil}
		self.last = self.last.next
	}
	self.size += 1
	return true
}

func (self *linkqueue) Shift() (vtype, bool) {
	if self.size != 0 {
		v := self.head.data
		self.head = self.head.next
		self.size -= 1
		return v, true
	} else {
		return -1, false
	}
}
