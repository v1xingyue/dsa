package gragh

import (
	"container/list"
	"log"
)

type Gragh struct {
	count int
	edges map[int](*list.List)
}

func NewGragh(count int) *Gragh {
	g := new(Gragh)
	g.count = count
	g.edges = make(map[int]*list.List)
	for i := 0; i < count; i++ {
		g.edges[i] = list.New()
	}
	return g
}

func (self *Gragh) addEdge(from, to int) bool {
	if from != to {
		self.edges[from].PushBack(to)
		return true
	}
	return false
}

func (self *Gragh) dfs(from, to int) {

}

// 广度优先搜索
func (self *Gragh) bfs(from, to int) {
	queue := new(list.List)
	//初始化把当前节点加入队列
	queue.PushBack(from)
	visited := make(map[int]bool)
	pre := make(map[int]int)
	found := false
	for queue.Len() != 0 {
		//弹出节点
		q := queue.Back()
		queue.Remove(q)
		num := q.Value.(int)
		log.Println("gragh node number out : ", num)

		//遍历对应的关系节点
		nodeList := self.edges[num]
		for nodeList.Len() != 0 {
			node := nodeList.Back()
			nodeList.Remove(node)
			i := node.Value.(int)
			log.Println("gragh node number : ", i)
			pre[i] = num
			log.Println("	=====   gragh pre node number : ", i, num)
			if visited[i] {
				continue
			}
			if i == to {
				log.Println("found ... ---- ", to, i)
				found = true
				break
			}
			visited[i] = true
			//把对应的子节点加入队列
			queue.PushBackList(self.edges[i])
		}
		if found {
			break
		}
	}
	if found {
		printPath(pre, from, to)
	}
}

//打印深度优先遍历的路径
func printPath(path map[int]int, from, to int) {
	log.Println(" \n --------------------  Print bfs Path  -------------- \n ")

	for k, v := range path {
		log.Println(" path route : ", k, v)
	}

	log.Println("path ---- <-  ", path[to])
	log.Println("path ---- <-  ", path[5])
}
