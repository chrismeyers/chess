package main

import (
	"fmt"
)

func main() {
	drawBoard()
}

func drawBoard() {
	fmt.Print("+ ")
	for i := 0; i < 8; i++ {
		fmt.Print("- ")
	}
	fmt.Print("+")
	fmt.Println()

	for i := 0; i < 8; i++ {
		fmt.Print("| ")
		for j := 0; j < 8; j++ {
			if i == 0 {
				switch j {
				case 0, 7:
					fmt.Print("r")
				case 1, 6:
					fmt.Print("n")
				case 2, 5:
					fmt.Print("b")
				case 3:
					fmt.Print("q")
				case 4:
					fmt.Print("k")
				}
			} else if i == 1 {
				fmt.Print("p")
			} else if i == 6 {
				fmt.Print("P")
			} else if i == 7 {
				switch j {
				case 0, 7:
					fmt.Print("R")
				case 1, 6:
					fmt.Print("N")
				case 2, 5:
					fmt.Print("B")
				case 3:
					fmt.Print("Q")
				case 4:
					fmt.Print("K")
				}
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	fmt.Print("+ ")
	for i := 0; i < 8; i++ {
		fmt.Print("- ")
	}
	fmt.Print("+")
	fmt.Println()
}
