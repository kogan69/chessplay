package chessplay_test

import (
	"fmt"
	"testing"
)

func Test_Promotion(t *testing.T) {
	fen := "k1r5/5P2/8/8/8/8/8/KR6 w - - 0 1"

	promoteTo := []string{"q", "r", "b", "n"}
	for _, piece := range promoteTo {
		fmt.Printf("FEN before promotion: %s", fen)
		game := initGame(t, fen)
		move := "f7f8" + piece
		err := game.MoveStr(move)

		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("board after promotion: %s", fen)
		fmt.Println(game.Position().Board().Draw())
	}
}

func Test_EnPassant(t *testing.T) {
	fen := "k7/8/8/1p6/2p5/8/1PP5/K7 w - - 0 1"
	game := initGame(t, fen)
	fmt.Println(game.Position().Board().Draw())
	err := game.MoveStr("b2b4")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(game.Position().Board().Draw())

	err = game.MoveStr("c4b3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(game.Position().Board().Draw())
}

func Test_Castling(t *testing.T) {
	fen := "rnbqk2r/pppp1ppp/3b1n2/4p3/4P3/3B1N2/PPPP1PPP/RNBQK2R w KQkq - 4 4"

	game := initGame(t, fen)
	fmt.Println(game.Position().Board().Draw())
	err := game.MoveStr("e1g1")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(game.Position().Board().Draw())

}
