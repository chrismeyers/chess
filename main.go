package main

import (
	"fmt"
	"os"
	"time"
)

const STARTING_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	clearScreen()

	board := NewBoard()

	board.Render(os.Stdout)
	time.Sleep(1 * time.Second)

	moves := []string{"f3", "e6", "g4", "Qh4"}

	for i, move := range moves {
		clearScreen()

		color := Black
		if i%2 == 0 {
			color = White
		}

		board.Move(color, move)
		board.Render(os.Stdout)
		time.Sleep(1 * time.Second)
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
