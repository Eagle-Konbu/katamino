package solvepkg

import (
	"time"
)

type Point struct {
	x int
	y int
}
type Piece struct {
	color     string
	positions []Point
}

func Fill(piece Piece, pointIdx int, board [][]bool) bool {
	width, height := len(board[0]), len(board)
	targets := make([]Point, 5)
	for i, p := range piece.positions {
		newTarget := Point{x: pointIdx + p.x, y: pointIdx + width*p.y}
		if newTarget.x >= width && newTarget.y >= height {
			return false
		}
		targets[i] = newTarget
	}

	for _, p := range targets {
		board[p.x][p.y] = true
	}
	return true
}

func Solve(width int, height int) ([][]string, float64) {
	var solTmp []string
	solutions := make([][]string, 4)

	puzzle := make([][]bool, height)
	for i := range puzzle {
		puzzle[i] = make([]bool, width)
	}
	pieces := []Piece{
		Piece{color: "fdf100", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 1, y: 0}, Point{x: 2, y: 1}, Point{x: 2, y: 2}}},
		Piece{color: "29005d", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 1, y: -2}, Point{x: 1, y: -3}}},
		Piece{color: "66e25a", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 2, y: -1}}},
		Piece{color: "BB0000", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: 1}, Point{x: 1, y: -1}}},
		Piece{color: "996e5b", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: 1}, Point{x: 1, y: -1}}},
		Piece{color: "234c83", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}, Point{x: 2, y: -2}}},
		Piece{color: "808080", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}, Point{x: 2, y: -1}}},
		Piece{color: "000080", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 0, y: -3}, Point{x: 0, y: -4}}},
		Piece{color: "dad400", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 0, y: -3}, Point{x: 1, y: -3}}},
		Piece{color: "62b7ff", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 1, y: -2}, Point{x: 2, y: -2}}},
		Piece{color: "ffc0cb", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 1, y: 0}, Point{x: 1, y: -1}}},
		Piece{color: "004900", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}}},
	}

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

func Rotate90(base Point) {

}
