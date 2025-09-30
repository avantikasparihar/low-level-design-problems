package types

const (
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

type Elevator struct {
	ID               int
	CurrentFloor     int
	NextFloors       map[int]struct{}
	CurrentDirection Direction
	Capacity         int
}

type Direction string
