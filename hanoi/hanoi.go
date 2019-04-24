package hanoi

import (
	"log"
)

func Hanoi(num int, fpos int, tpos int) {
	if num == 1 {
		log.Printf(" Move disk : %d to pos: %d \n", fpos, tpos)
	} else {
		mpos := 6 - fpos - tpos
		Hanoi(num-1, fpos, mpos)
		Hanoi(1, fpos, tpos)
		Hanoi(num-1, mpos, tpos)
	}
}
