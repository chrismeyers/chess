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

func TestBoard_Move(t *testing.T) {
	tests := []struct {
		name      string
		color     PieceColor
		move      string
		wantEmpty []int
		wantMoved []int
		wantPiece *Piece
	}{
		{
			name:      "Successfully moves a white pawn",
			color:     White,
			move:      "f3",
			wantEmpty: []int{6, 5},
			wantMoved: []int{5, 5},
			wantPiece: &Piece{Type: Pawn, Color: White, Symbol: "P"},
		},
		{
			name:      "Successfully moves a black pawn",
			color:     Black,
			move:      "e6",
			wantEmpty: []int{1, 4},
			wantMoved: []int{2, 4},
			wantPiece: &Piece{Type: Pawn, Color: Black, Symbol: "p"},
		},
		// TODO: This test case uncovered a complication in the moving logic.
		// In the cases where there are multiple of the same pieces (knight,
		// bishop, rook), the correct piece needs to be moved based on the
		// possible moves it can make instead of the first piece found when
		// searching left to right through the ranks.
		// {
		// 	name:      "Successfully moves a white knight",
		// 	color:     White,
		// 	move:      "Nf3",
		// 	wantEmpty: []int{7, 6},
		// 	wantMoved: []int{5, 5},
		// 	wantPiece: &Piece{Type: Knight, Color: White, Symbol: "N"},
		// },
		{
			name:      "Successfully moves a black knight",
			color:     Black,
			move:      "Nc6",
			wantEmpty: []int{0, 1},
			wantMoved: []int{2, 2},
			wantPiece: &Piece{Type: Knight, Color: Black, Symbol: "n"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := NewBoard()

			board.Move(tt.color, tt.move)

			if board[tt.wantEmpty[0]][tt.wantEmpty[1]].Piece != nil {
				t.Errorf("Expected square %v to be empty", tt.wantEmpty)
			}

			moved := board[tt.wantMoved[0]][tt.wantMoved[1]].Piece
			if !cmp.Equal(moved, tt.wantPiece) {
				t.Errorf("Moved piece is incorrect: %s", cmp.Diff(moved, tt.wantPiece))
			}
		})
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
