package solvepkg

import (
	"fmt"
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

var board [][]string
var solutions [][]string

func Fill(piece Piece, pointIdx int, reset bool) bool {
	width, height := len(board[0]), len(board)
	targets := make([]Point, 5)
	for i, p := range piece.positions {
		newTarget := Point{x: pointIdx%width + p.x, y: pointIdx/width + p.y}
		if newTarget.x >= width || newTarget.y >= height || newTarget.x < 0 || newTarget.y < 0 {
			return false
		}
		if board[newTarget.y][newTarget.x] != "" && !reset {
			return false
		}
		targets[i] = newTarget
	}

	for _, p := range targets {
		if reset {
			board[p.y][p.x] = ""
		} else {
			board[p.y][p.x] = piece.color
		}
	}
	return true
}

func Completes() bool {
	width, height := len(board[0]), len(board)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func Search(pieces []Piece) {
	width, height := len(board[0]), len(board)

	idx := 0
	for i := 0; i < 60; i++ {
		x, y := i%width, i/width
		if board[y][x] == "" {
			idx = i
			break
		}
	}

	for i, piece := range pieces {
		solution := make([]string, 60)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				solution[i*width+j] = board[i][j]
			}
		}
		solutions = append(solutions, solution)
		if Fill(piece, idx, false) {
			// fmt.Println("len=", len(pieces), ",i=", i)
			if Completes() {
				fmt.Println("found")
				solution := make([]string, 60)
				for i := 0; i < height; i++ {
					for j := 0; j < width; j++ {
						solution[i*width+j] = board[i][j]
					}
				}
				solutions = append(solutions, solution)
			} else {
				var remainingPieces []Piece
				remainingPieces = append(remainingPieces, append(pieces[:i], pieces[i+1:]...)...)
				Search(remainingPieces)
			}
			Fill(piece, idx, true)
		}
	}
}

func Solve(width int, height int) ([][]string, float64) {
	solutions := make([][]string, 0)

	board = make([][]string, height)
	for i := range board {
		board[i] = make([]string, width)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			board[i][j] = ""
		}
	}
	pieces := []Piece{
		Piece{color: "#fdf100", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 2, y: 1}}},
		Piece{color: "#29005d", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 0, y: 2}, Point{x: 1, y: 0}, Point{x: 1, y: -1}}},
		Piece{color: "#66e25a", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 2, y: -1}}},
		Piece{color: "#BB0000", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: 1}, Point{x: 1, y: -1}}},
		Piece{color: "#996e5b", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: 1}, Point{x: 1, y: -1}, Point{x: 1, y: -2}}},
		Piece{color: "#234c83", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}, Point{x: 2, y: -2}}},
		Piece{color: "#808080", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}, Point{x: 2, y: -1}}},
		Piece{color: "#000080", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 0, y: 2}, Point{x: 0, y: 3}, Point{x: 0, y: 4}}},
		Piece{color: "#dad400", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 0, y: 2}, Point{x: 0, y: 3}, Point{x: 1, y: 0}}},
		Piece{color: "#62b7ff", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 0, y: 2}, Point{x: 1, y: 0}, Point{x: 2, y: 0}}},
		Piece{color: "#ffc0cb", positions: []Point{Point{x: 0, y: 0}, Point{x: 0, y: 1}, Point{x: 0, y: 2}, Point{x: 1, y: 1}, Point{x: 1, y: 2}}},
		Piece{color: "#004900", positions: []Point{Point{x: 0, y: 0}, Point{x: 1, y: 0}, Point{x: 2, y: 0}, Point{x: 1, y: -1}, Point{x: 1, y: -2}}},
	}

	rotatedPieces := []Piece{
		Rotate90(pieces[0], Point{x: 0, y: 1}),
		Rotate180(pieces[0], Point{x: 2, y: 1}),
		Rotate270(pieces[0], Point{x: 2, y: 0}),
		Rotate90(pieces[1], Point{x: 0, y: 2}),
		Rotate180(pieces[1], Point{x: 1, y: 0}),
		Rotate270(pieces[1], Point{x: 1, y: -1}),
		Rotate90(pieces[2], Point{x: 0, y: 1}),
		Rotate180(pieces[2], Point{x: 2, y: -1}),
		Rotate270(pieces[2], Point{x: 2, y: -1}),
		Rotate90(pieces[4], Point{x: 1, y: 1}),
		Rotate180(pieces[4], Point{x: 1, y: 1}),
		Rotate270(pieces[4], Point{x: 1, y: -2}),
		Rotate90(pieces[5], Point{x: 0, y: 0}),
		Rotate90(pieces[6], Point{x: 0, y: 0}),
		Rotate180(pieces[6], Point{x: 2, y: -1}),
		Rotate270(pieces[6], Point{x: 1, y: -2}),
		Rotate90(pieces[7], Point{x: 0, y: 4}),
		Rotate90(pieces[8], Point{x: 0, y: 3}),
		Rotate180(pieces[8], Point{x: 1, y: 0}),
		Rotate270(pieces[8], Point{x: 1, y: 0}),
		Rotate90(pieces[9], Point{x: 0, y: 2}),
		Rotate180(pieces[9], Point{x: 2, y: 0}),
		Rotate270(pieces[9], Point{x: 2, y: 0}),
		Rotate90(pieces[10], Point{x: 0, y: 2}),
		Rotate180(pieces[10], Point{x: 1, y: 2}),
		Rotate270(pieces[10], Point{x: 0, y: 0}),
		Rotate90(pieces[11], Point{x: 0, y: 0}),
		Rotate180(pieces[11], Point{x: 2, y: 0}),
		Rotate270(pieces[11], Point{x: 1, y: -2}),
	}

	pieces = append(pieces, rotatedPieces...)

	var flippedPieces []Piece
	for _, piece := range pieces {
		target := piece.positions[0]
		for _, position := range piece.positions {
			if position.x > target.x {
				target = position
			} else if position.x == target.x && position.y < target.y {
				target = position
			}
		}
		flippedPieces = append(flippedPieces, Flip(piece, target))
	}
	pieces = append(pieces, flippedPieces...)

	startTime := time.Now()

	Search(pieces)

	for i := range solutions {
		for j := range solutions[i] {
			if solutions[i][j] == "" {
				solutions[i][j] = "#000000"
			}
		}
	}

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

func Rotate180(piece Piece, base Point) Piece {
	newPiece := Piece{color: piece.color, positions: make([]Point, 5)}
	for i, p := range piece.positions {
		newPiece.positions[i] = Point{x: base.x - p.x, y: base.y - p.y}
	}
	return newPiece
}

func Rotate270(piece Piece, base Point) Piece {
	newPiece := Piece{color: piece.color, positions: make([]Point, 5)}
	for i, p := range piece.positions {
		newPiece.positions[i] = Point{x: p.y - base.y, y: base.x - p.x}
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
