package internal

const (
	OutcomeTypeUnknown OutcomeType = iota
	OutcomeTypeHorizontal1
	OutcomeTypeVertical1
	OutcomeTypeDiagonal1
	OutcomeTypeHorizontal2
	OutcomeTypeVertical2
	OutcomeTypeDiagonal2
	OutcomeTypeDraw
)

type BoardManager interface {
	DropBall(column int) error
	CheckOutcome() OutcomeType
	GetBoard() [][]int
}

type DefaultBoard struct {
	Matrix [][]int
}

type OutcomeType int

func NewDefaultBoard() BoardManager {
	matrix := make([][]int, 7)
	for i := range matrix {
		matrix[i] = make([]int, 6)
	}
	return &DefaultBoard{
		Matrix: matrix,
	}
}

func (b *DefaultBoard) DropBall(column int) error {
	// impl logic to drop ball in a given column

	return nil
}

func (b *DefaultBoard) CheckOutcome() OutcomeType {
	// evaluate board
	return OutcomeTypeUnknown
}

func (b *DefaultBoard) GetBoard() [][]int {
	return b.Matrix
}
