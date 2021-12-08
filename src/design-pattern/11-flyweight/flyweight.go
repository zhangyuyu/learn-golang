package flyweight

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// 其他棋子
}

// ChessPieceUnit 棋子享元
type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

// NewChessPieceUint 工厂
func NewChessPieceUint(id int) *ChessPieceUnit {
	return units[id]
}

// ChessPiece 棋子
type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

// ChessBoard 棋局
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUint(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}
