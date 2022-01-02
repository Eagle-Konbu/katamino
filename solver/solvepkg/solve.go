package solvepkg

import (
	"time"
)

type Point struct {
	x int
	y int
}
type Piece struct {
	color    string
	position Point
}

func Solve(width int, height int) ([][]string, float64) {
	var solTmp []string
	// var solutions [][]string
	solutions := make([][]string, 4)

	startTime := time.Now()

	for i := 0; i < 30; i++ {
		solTmp = append(solTmp, "#AA0000")
	}
	for i := 0; i < 30; i++ {
		solTmp = append(solTmp, "#00AA00")
	}

	for i := 0; i < 4; i++ {
		tmp := make([]string, len(solTmp))
		copy(tmp, solTmp)
		solutions[i] = tmp
	}

	calcTime := float64(time.Since(startTime).Milliseconds()) / 1000

	return solutions, calcTime
}
