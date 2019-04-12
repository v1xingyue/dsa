package hanoi

import (
	"log"
)

func Hanoi(num int, fpos int, tpos int) {
	if num == 1 {
		log.Printf(" Move disk : %d to pos: %d \n", fpos, tpos)
	} else {
		mpos := 0
		switch fpos {
		case 1:
			if tpos == 2 {
				mpos = 3
			} else {
				mpos = 2
			}
		case 2:
			if tpos == 3 {
				mpos = 1
			} else {
				mpos = 3
			}
		case 3:
			if tpos == 1 {
				mpos = 2
			} else {
				mpos = 1
			}
		}
		Hanoi(num-1, fpos, mpos)
		Hanoi(1, fpos, tpos)
		Hanoi(num-1, mpos, tpos)
	}
}
