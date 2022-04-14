package main

import (
	"fmt"
	"io"
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
	Symbol string
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
						Symbol: "r",
					}
				case 1, 6:
					board[i][j] = Square{
						Piece:  Knight,
						Color:  Black,
						Symbol: "n",
					}
				case 2, 5:
					board[i][j] = Square{
						Piece:  Bishop,
						Color:  Black,
						Symbol: "b",
					}
				case 3:
					board[i][j] = Square{
						Piece:  Queen,
						Color:  Black,
						Symbol: "q",
					}
				case 4:
					board[i][j] = Square{
						Piece:  King,
						Color:  Black,
						Symbol: "k",
					}
				}
			} else if i == 1 {
				board[i][j] = Square{
					Piece:  Pawn,
					Color:  Black,
					Symbol: "p",
				}
			} else if i == 6 {
				board[i][j] = Square{
					Piece:  Pawn,
					Color:  White,
					Symbol: "P",
				}
			} else if i == 7 {
				switch j {
				case 0, 7:
					board[i][j] = Square{
						Piece:  Rook,
						Color:  White,
						Symbol: "R",
					}
				case 1, 6:
					board[i][j] = Square{
						Piece:  Knight,
						Color:  White,
						Symbol: "N",
					}
				case 2, 5:
					board[i][j] = Square{
						Piece:  Bishop,
						Color:  White,
						Symbol: "B",
					}
				case 3:
					board[i][j] = Square{
						Piece:  Queen,
						Color:  White,
						Symbol: "Q",
					}
				case 4:
					board[i][j] = Square{
						Piece:  King,
						Color:  White,
						Symbol: "K",
					}
				}
			} else {
				board[i][j] = Square{
					Piece:  Empty,
					Color:  None,
					Symbol: ".",
				}
			}
		}
	}

	return board
}

func (board *Board) Render(w io.Writer) {
	drawHorizontalBorder(w)

	for i := 0; i < 8; i++ {
		fmt.Fprintf(w, "%d | ", 8-i)
		for j := 0; j < 8; j++ {
			fmt.Fprint(w, board[i][j].Symbol+" ")
		}
		fmt.Fprintln(w, "|")
	}

	drawHorizontalBorder(w)

	drawFileCoordinates(w)
}

func RenderFEN(w io.Writer, fen string) {
	parts := strings.Split(fen, " ")
	pieces := parts[0]
	ranks := strings.Split(pieces, "/")

	drawHorizontalBorder(w)

	for i := 0; i < 8; i++ {
		fmt.Fprintf(w, "%d | ", 8-i)
		rank := strings.Split(ranks[i], "")
		for _, square := range rank {
			empties, err := strconv.Atoi(square)
			if err != nil {
				fmt.Fprint(w, square+" ")
			} else {
				for j := 0; j < empties; j++ {
					fmt.Fprint(w, ". ")
				}
			}
		}
		fmt.Fprintln(w, "|")
	}

	drawHorizontalBorder(w)

	drawFileCoordinates(w)
}

func drawHorizontalBorder(w io.Writer) {
	fmt.Fprint(w, "  + ")
	for i := 0; i < 8; i++ {
		fmt.Fprint(w, "- ")
	}
	fmt.Fprint(w, "+")
	fmt.Fprintln(w)
}

func drawFileCoordinates(w io.Writer) {
	fmt.Fprint(w, "    ")
	for i := 0; i < 8; i++ {
		fmt.Fprintf(w, "%c ", rune(97+i))
	}
	fmt.Fprintln(w)
}
