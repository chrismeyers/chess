package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Piece int

const (
	Empty Piece = iota
	Pawn
	Rook
	Knight
	Bishop
	Queen
	King
)

type Color int

const (
	None Color = iota
	White
	Black
)

type Square struct {
	Piece  Piece
	Color  Color
	Symbol rune
}

type Board [8][8]Square

func NewBoard() Board {
	board := Board{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == 0 {
				switch j {
				case 0, 7:
					board[i][j] = Square{
						Piece:  Rook,
						Color:  Black,
						Symbol: 'r',
					}
				case 1, 6:
					board[i][j] = Square{
						Piece:  Knight,
						Color:  Black,
						Symbol: 'n',
					}
				case 2, 5:
					board[i][j] = Square{
						Piece:  Bishop,
						Color:  Black,
						Symbol: 'b',
					}
				case 3:
					board[i][j] = Square{
						Piece:  Queen,
						Color:  Black,
						Symbol: 'q',
					}
				case 4:
					board[i][j] = Square{
						Piece:  King,
						Color:  Black,
						Symbol: 'k',
					}
				}
			} else if i == 1 {
				board[i][j] = Square{
					Piece:  Pawn,
					Color:  Black,
					Symbol: 'p',
				}
			} else if i == 6 {
				board[i][j] = Square{
					Piece:  Pawn,
					Color:  White,
					Symbol: 'P',
				}
			} else if i == 7 {
				switch j {
				case 0, 7:
					board[i][j] = Square{
						Piece:  Rook,
						Color:  White,
						Symbol: 'R',
					}
				case 1, 6:
					board[i][j] = Square{
						Piece:  Knight,
						Color:  White,
						Symbol: 'N',
					}
				case 2, 5:
					board[i][j] = Square{
						Piece:  Bishop,
						Color:  White,
						Symbol: 'B',
					}
				case 3:
					board[i][j] = Square{
						Piece:  Queen,
						Color:  White,
						Symbol: 'Q',
					}
				case 4:
					board[i][j] = Square{
						Piece:  King,
						Color:  White,
						Symbol: 'K',
					}
				}
			} else {
				board[i][j] = Square{
					Piece:  Empty,
					Color:  None,
					Symbol: '.',
				}
			}
		}
	}

	return board
}

func (board *Board) Render() {
	drawHorizontalBorder()

	for i := 0; i < 8; i++ {
		fmt.Printf("%d | ", 8-i)
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board[i][j].Symbol)
		}
		fmt.Println("|")
	}

	drawHorizontalBorder()

	drawFileCoordinates()
}

func RenderFEN(fen string) {
	parts := strings.Split(fen, " ")
	pieces := parts[0]
	ranks := strings.Split(pieces, "/")

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

	drawFileCoordinates()
}

func drawHorizontalBorder() {
	fmt.Print("  + ")
	for i := 0; i < 8; i++ {
		fmt.Print("- ")
	}
	fmt.Print("+")
	fmt.Println()
}

func drawFileCoordinates() {
	fmt.Print("    ")
	for i := 0; i < 8; i++ {
		fmt.Printf("%c ", rune(97+i))
	}
	fmt.Println()
}
