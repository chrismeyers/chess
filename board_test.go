package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBoard_NewBoard(t *testing.T) {
	board := NewBoard()

	emptyRank := [8]Square{
		{Piece: nil},
		{Piece: nil},
		{Piece: nil},
		{Piece: nil},
		{Piece: nil},
		{Piece: nil},
		{Piece: nil},
		{Piece: nil},
	}

	expected := Board{
		{
			Square{Piece: &Piece{Type: Rook, Color: Black, Symbol: "r"}},
			Square{Piece: &Piece{Type: Knight, Color: Black, Symbol: "n"}},
			Square{Piece: &Piece{Type: Bishop, Color: Black, Symbol: "b"}},
			Square{Piece: &Piece{Type: Queen, Color: Black, Symbol: "q"}},
			Square{Piece: &Piece{Type: King, Color: Black, Symbol: "k"}},
			Square{Piece: &Piece{Type: Bishop, Color: Black, Symbol: "b"}},
			Square{Piece: &Piece{Type: Knight, Color: Black, Symbol: "n"}},
			Square{Piece: &Piece{Type: Rook, Color: Black, Symbol: "r"}},
		},
		{
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
			Square{Piece: &Piece{Type: Pawn, Color: Black, Symbol: "p"}},
		},
		emptyRank,
		emptyRank,
		emptyRank,
		emptyRank,
		{
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
			Square{Piece: &Piece{Type: Pawn, Color: White, Symbol: "P"}},
		},
		{
			Square{Piece: &Piece{Type: Rook, Color: White, Symbol: "R"}},
			Square{Piece: &Piece{Type: Knight, Color: White, Symbol: "N"}},
			Square{Piece: &Piece{Type: Bishop, Color: White, Symbol: "B"}},
			Square{Piece: &Piece{Type: Queen, Color: White, Symbol: "Q"}},
			Square{Piece: &Piece{Type: King, Color: White, Symbol: "K"}},
			Square{Piece: &Piece{Type: Bishop, Color: White, Symbol: "B"}},
			Square{Piece: &Piece{Type: Knight, Color: White, Symbol: "N"}},
			Square{Piece: &Piece{Type: Rook, Color: White, Symbol: "R"}},
		},
	}

	if !cmp.Equal(board, expected) {
		t.Errorf("NewBoard did not generate correct starting position: %s", cmp.Diff(board, expected))
	}
}

func TestBoard_Render(t *testing.T) {
	board := NewBoard()

	var output bytes.Buffer
	board.Render(&output)

	expected :=
		"  + - - - - - - - - +\n" +
			"8 | r n b q k b n r |\n" +
			"7 | p p p p p p p p |\n" +
			"6 | . . . . . . . . |\n" +
			"5 | . . . . . . . . |\n" +
			"4 | . . . . . . . . |\n" +
			"3 | . . . . . . . . |\n" +
			"2 | P P P P P P P P |\n" +
			"1 | R N B Q K B N R |\n" +
			"  + - - - - - - - - +\n" +
			"    a b c d e f g h \n"

	if expected != output.String() {
		t.Errorf("Board rendered incorrectly \nwant:\n%s\ngot:\n%s\n", expected, output.String())
	}
}

func TestBoard_RenderFEN(t *testing.T) {
	tests := []struct {
		name string
		FEN  string
		want string
	}{
		{
			name: "Successfully renders starting position",
			FEN:  STARTING_FEN,
			want: "  + - - - - - - - - +\n" +
				"8 | r n b q k b n r |\n" +
				"7 | p p p p p p p p |\n" +
				"6 | . . . . . . . . |\n" +
				"5 | . . . . . . . . |\n" +
				"4 | . . . . . . . . |\n" +
				"3 | . . . . . . . . |\n" +
				"2 | P P P P P P P P |\n" +
				"1 | R N B Q K B N R |\n" +
				"  + - - - - - - - - +\n" +
				"    a b c d e f g h \n",
		},
		{
			name: "Successfully renders Fool's mate",
			FEN:  "rnb1kbnr/pppp1ppp/4p3/8/6Pq/5P2/PPPPP2P/RNBQKBNR w KQkq - 1 3",
			want: "  + - - - - - - - - +\n" +
				"8 | r n b . k b n r |\n" +
				"7 | p p p p . p p p |\n" +
				"6 | . . . . p . . . |\n" +
				"5 | . . . . . . . . |\n" +
				"4 | . . . . . . P q |\n" +
				"3 | . . . . . P . . |\n" +
				"2 | P P P P P . . P |\n" +
				"1 | R N B Q K B N R |\n" +
				"  + - - - - - - - - +\n" +
				"    a b c d e f g h \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			RenderFEN(&output, tt.FEN)

			if tt.want != output.String() {
				t.Errorf("Board rendered incorrectly \nwant:\n%s\ngot:\n%s\n", tt.want, output.String())
			}
		})
	}
}
