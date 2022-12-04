package chessplay_test

import (
	"testing"

	"github.com/notnil/chess"
)

func initGame(t *testing.T, fen string) *chess.Game {
	startingFen, err := chess.FEN(fen)
	if err != nil {
		t.Fatal(err)
	}
	game := chess.NewGame(chess.UseNotation(chess.UCINotation{}), startingFen)
	return game
}
