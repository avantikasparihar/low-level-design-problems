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
	DropBall(column, player int) error
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

func (b *DefaultBoard) DropBall(column, player int) error {
	// impl logic to drop ball in a given column
	j := column - 1
	for i := 6; i >= 0; i-- {
		if b.Matrix[i][j] == 0 {
			b.Matrix[i][j] = player
			break
		}
	}

	return nil
}

func (b *DefaultBoard) CheckOutcome() OutcomeType {
	// evaluate board

	// 1. check horizontal row
	for i := 6; i >= 0; i-- {
		cnt1, cnt2 := 0, 0
		for j := 0; j <= 5; j++ {
			switch b.Matrix[i][j] {
			case 0:
				cnt1, cnt2 = 0, 0
			case 1:
				cnt1++
				cnt2 = 0
			case 2:
				cnt2++
				cnt1 = 0
			}
			if cnt1 == 4 {
				return OutcomeTypeHorizontal1
			}
			if cnt2 == 4 {
				return OutcomeTypeHorizontal2
			}
		}
	}

	// 2. check vertical row

	// 3. check diagonal row
	return OutcomeTypeUnknown
}

func (b *DefaultBoard) GetBoard() [][]int {
	return b.Matrix
}
