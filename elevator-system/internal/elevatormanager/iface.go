package elevatormanager

import "github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"

type ElevatorManager interface {
	FetchElevator(floorNo int, direction types.Direction) error
	AllocateElevator(elevatorID int, floorNo int) error
}
