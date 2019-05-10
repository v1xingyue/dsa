package trietree

import (
	//"log"
	"strconv"
	"strings"
)

type IpTree struct {
	number   int
	children map[int]*IpTree
}

func (self *IpTree) findIp(ip string) bool {
	ip = strings.Trim(ip, " \r\n\t")
	nums := strings.Split(ip, ".")
	if len(nums) != 4 {
		return false
	}
	parent := self
	for _, v := range nums {
		num, _ := strconv.Atoi(v)
		if parent == nil {
			return false
		}
		v, ok := parent.children[num]
		if ok == false {
			return false
		}
		parent = v
	}
	return true
}

func (self *IpTree) addIp(ip string) bool {
	ip = strings.Trim(ip, " \r\n\t")
	parent := self
	nums := strings.Split(ip, ".")
	if len(nums) != 4 {
		return false
	}
	for _, v := range nums {
		num, _ := strconv.Atoi(v)
		if parent.children == nil {
			parent.children = make(map[int]*IpTree)
		}
		if _, ok := parent.children[num]; ok == false {
			parent.children[num] = &IpTree{number: num}
		}
		parent = parent.children[num]
	}
	return true
}
