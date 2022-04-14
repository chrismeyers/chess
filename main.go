package main

import (
	"fmt"
	"os"
	"time"
)

const STARTING_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	clearScreen()

	fmt.Println("Render FEN")
	RenderFEN(os.Stdout, STARTING_FEN)
	time.Sleep(1 * time.Second)

	clearScreen()

	board := NewBoard()
	for i := 0; i < 3; i++ {
		fmt.Printf("Render Board %d\n", i+1)
		board.Render(os.Stdout)
		time.Sleep(1 * time.Second)
		clearScreen()
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
