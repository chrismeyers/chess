package main

import (
	"fmt"
	"strconv"
	"strings"
)

const STARTING_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	draw(STARTING_FEN)
}

func draw(fen string) {
	parts := strings.Split(fen, " ")
	pieces := parts[0]
	ranks := strings.Split(pieces, "/")

	drawHorizontalBorder := func() {
		fmt.Print("  + ")
		for i := 0; i < 8; i++ {
			fmt.Print("- ")
		}
		fmt.Print("+")
		fmt.Println()
	}

	drawHorizontalBorder()

	for i := 0; i < 8; i++ {
		fmt.Printf("%d | ", 8-i)
		rank := strings.Split(ranks[i], "")
		for _, square := range rank {
			empties, err := strconv.Atoi(square)
			if err != nil {
				fmt.Print(square + " ")
			} else {
				for j := 0; j < empties; j++ {
					fmt.Print(". ")
				}
			}
		}
		fmt.Println("|")
	}

	drawHorizontalBorder()

	fmt.Print("    ")
	for i := 0; i < 8; i++ {
		fmt.Printf("%c ", rune(97+i))
	}
	fmt.Println()
}
