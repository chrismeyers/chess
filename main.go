package main

import (
	"fmt"
	"os"
)

const STARTING_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	board := NewBoard()
	board.Render(os.Stdout)

	fmt.Println()

	RenderFEN(os.Stdout, STARTING_FEN)
}
