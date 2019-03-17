// 使用数组实现的queue

package queue

type XQueue struct {
    items []int
    head int
    last int
    size int
}

func (self *XQueue) Init(size int){
    self.items = make([]int,size,size)
    self.head = 0
    self.last = 0
    self.size = size
}


func (self *XQueue) Insert(value int){
    self.items[self.last] = value
    self.last = ( self.last + 1 ) %  self.size
    if self.last == self.head {
        self.head = self.head + 1
    }
}

func (self *XQueue) Shift() (int,error){
    if self.Empty(){
        return -1,ErrorQueueEmpty
    }
    v := self.items[self.head]
    self.head = (self.head + 1) % (self.size)
    return v,nil
}

func (self *XQueue) Empty() bool{
    return self.head == self.last
}

func (self *XQueue) Count() int {
    if self.last > self.head {
        return self.last - self.head
    } else {
        return self.size - (self.head - self.last) 
    }
}
