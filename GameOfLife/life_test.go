package gamelife

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	w := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	PrintLife(w)
	gameOfLife(w)
	PrintLife(w)

	w = [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	PrintLife(w)
}

func TestLifeLive(t *testing.T) {
	w := GameLifeInit(54, 96, 50)
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		gameOfLife(w)
		PrintLife(w)
		time.Sleep(2 * time.Second)
	}
}
