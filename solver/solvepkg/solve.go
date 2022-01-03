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

func Fill(piece Piece, pointIdx int, board [][]string) bool {
	width, height := len(board[0]), len(board)
	targets := make([]Point, 5)
	for i, p := range piece.positions {
		newTarget := Point{x: pointIdx%width + p.x, y: pointIdx/width + width*p.y}
		if newTarget.x >= width && newTarget.y >= height {
			return false
		}
		targets[i] = newTarget
	}

	for _, p := range targets {
		board[p.x][p.y] = piece.color
	}
	return true
}

func Search(pieces []Piece, board [][]string, solutions [][]string) {
	width, height := len(board[0]), len(board)
	var remainingPieces []Piece
	var clonedBoard [][]string
	copy(remainingPieces, pieces)
	copy(clonedBoard, board)

	idx := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == "" {
				idx = i*width + j
				break
			}
		}
	}

	for i, pTmp := range remainingPieces {
		allAnglePieces := []Piece{pTmp}
		for _, piece := range allAnglePieces {
			if Fill(piece, idx, clonedBoard) {
				remainingPieces = append(remainingPieces[:i], remainingPieces[i+1:]...)
				if len(remainingPieces) == 0 {
					solution := make([]string, 60)
					for i := 0; i < height; i++ {
						for j := 0; j < width; j++ {
							solution[i*width+j] = clonedBoard[i][j]
						}
					}
					solutions = append(solutions, solution)
				} else {
					Search(remainingPieces, clonedBoard, solutions)
				}
			}
		}
	}
}

func Solve(width int, height int) ([][]string, float64) {
	// var solTmp []string
	var solutions [][]string

	puzzle := make([][]string, height)
	for i := range puzzle {
		puzzle[i] = make([]string, width)
	}
	pieces := []Piece{
		Piece{color: "#fdf100", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 1, y: 0}, Point{x: 2, y: 1}, Point{x: 2, y: 2}}},
		Piece{color: "#29005d", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 1, y: -2}, Point{x: 1, y: -3}}},
		Piece{color: "#66e25a", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 2, y: -1}}},
		Piece{color: "#BB0000", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: 1}, Point{x: 1, y: -1}}},
		Piece{color: "#996e5b", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: 1}, Point{x: 1, y: -1}}},
		Piece{color: "#234c83", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}, Point{x: 2, y: -2}}},
		Piece{color: "#808080", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}, Point{x: 2, y: -1}}},
		Piece{color: "#000080", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 0, y: -3}, Point{x: 0, y: -4}}},
		Piece{color: "#dad400", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 0, y: -3}, Point{x: 1, y: -3}}},
		Piece{color: "#62b7ff", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 1, y: -2}, Point{x: 2, y: -2}}},
		Piece{color: "#ffc0cb", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: -1}, Point{x: 0, y: -2}, Point{x: 1, y: 0}, Point{x: 1, y: -1}}},
		Piece{color: "#004900", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}}},
	}

	startTime := time.Now()

	Search(pieces, puzzle, solutions)

	// for i := 0; i < 30; i++ {
	// 	solTmp = append(solTmp, "#AA0000")
	// }
	// for i := 0; i < 30; i++ {
	// 	solTmp = append(solTmp, "#00AA00")
	// }

	// for i := 0; i < 4; i++ {
	// 	tmp := make([]string, len(solTmp))
	// 	copy(tmp, solTmp)
	// 	solutions[i] = tmp
	// }

	calcTime := float64(time.Since(startTime).Milliseconds()) / 1000

	return solutions, calcTime
}

func Rotate90(piece Piece, base Point) Piece {
	newPiece := Piece{color: piece.color, positions: make([]Point, 5)}
	for i, p := range piece.positions {
		newPiece.positions[i] = Point{x: base.y - p.y, y: p.x - base.x}
	}
	return newPiece
}

func Flip(piece Piece, base Point) Piece {
	newPiece := Piece{color: piece.color, positions: make([]Point, 5)}
	for i, p := range piece.positions {
		newPiece.positions[i] = Point{x: base.x - p.x, y: p.y - base.y}
	}
	return newPiece
}
