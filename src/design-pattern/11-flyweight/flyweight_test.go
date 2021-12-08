package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChessBoard(t *testing.T) {
	board := NewChessBoard()
	board.Move(1, 1, 2)
	assert.Equal(t, board.chessPieces[1].X, 1)
	assert.Equal(t, board.chessPieces[1].Y, 2)
}
