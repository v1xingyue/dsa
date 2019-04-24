package queue

import (
	"testing"
)

func TestLinkQueue(t *testing.T) {
	lq := NewListQueue(16)
	for i := 1; i < 32; i++ {
		t.Log(lq.Insert(vtype(i)))
	}

	for {
		if v, res := lq.Shift(); res == true {
			t.Log(v)
		} else {
			break
		}
	}

}
