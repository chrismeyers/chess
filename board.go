package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type PieceType int

const (
	NoType PieceType = iota
	Pawn
	Rook
	Knight
	Bishop
	Queen
	King
)

type PieceColor int

const (
	NoColor PieceColor = iota
	White
	Black
)

type Piece struct {
	Type   PieceType
	Color  PieceColor
	Symbol string
}

type Square struct {
	Piece *Piece
}

type Board [8][8]Square

func NewBoard() Board {
	board := Board{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == 0 {
				switch j {
				case 0, 7:
					board[i][j] = Square{Piece: &Piece{Type: Rook, Color: Black, Symbol: "r"}}
				case 1, 6:
					board[i][j] = Square{Piece: &Piece{Type: Knight, Color: Black, Symbol: "n"}}
				case 2, 5:
					board[i][j] = Square{Piece: &Piece{Type: Bishop, Color: Black, Symbol: "b"}}
				case 3:
					board[i][j] = Square{Piece: &Piece{Type: Queen, Color: Black, Symbol: "q"}}
				case 4:
					board[i][j] = Square{Piece: &Piece{Type: King, Color: Black, Symbol: "k"}}
				}
			} else if i == 1 {
				board[i][j] = Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}}
			} else if i == 6 {
				board[i][j] = Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}}
			} else if i == 7 {
				switch j {
				case 0, 7:
					board[i][j] = Square{Piece: &Piece{Type: Rook, Color: White, Symbol: "R"}}
				case 1, 6:
					board[i][j] = Square{Piece: &Piece{Type: Knight, Color: White, Symbol: "N"}}
				case 2, 5:
					board[i][j] = Square{Piece: &Piece{Type: Bishop, Color: White, Symbol: "B"}}
				case 3:
					board[i][j] = Square{Piece: &Piece{Type: Queen, Color: White, Symbol: "Q"}}
				case 4:
					board[i][j] = Square{Piece: &Piece{Type: King, Color: White, Symbol: "K"}}
				}
			} else {
				board[i][j] = Square{Piece: nil}
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
			symbol := "."
			if board[i][j].Piece != nil {
				symbol = board[i][j].Piece.Symbol
			}
			fmt.Fprint(w, symbol+" ")
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
