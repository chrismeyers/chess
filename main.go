package main

import "fmt"

const STARTING_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	board := NewBoard()
	board.Render()

	fmt.Println()

	RenderFEN(STARTING_FEN)
}
