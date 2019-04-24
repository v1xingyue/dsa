package stack

//链表实现的一个栈
//默认最大元素个数为2048
type Node struct {
	val  rune
	next *Node
}

type Stack struct {
	head *Node
	size int
	max  int
}

func NewStack() *Stack {
	s := new(Stack)
	s.size = 0
	s.max = 2048
	return s
}

// 大小控制，默认可以装2048个元素,可以传递大小参数
func NewStackWithSize(max int) *Stack {
	s := new(Stack)
	s.size = 0
	s.max = max
	return s
}

//push 操作注意，头指针的next 指针变动
//遍历结果，从head 指针往next 遍历即可完成操作
func (self *Stack) Push(c rune) bool {
	if self.size >= self.max {
		return false
	}
	self.size++
	self.head = &Node{val: c, next: self.head}
	return true
}

func (self *Stack) Pop() rune {
	if self.head == nil {
		return 0
	}
	b := self.head.val
	self.head = self.head.next
	self.size--
	return b
}

func (self *Stack) Count() int {
	return self.size
}

func (self *Stack) IsEmpty() bool {
	return self.size == 0
}
