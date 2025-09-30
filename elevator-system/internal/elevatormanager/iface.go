package elevatormanager

import "github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"

type ElevatorManager interface {
	GetElevatorInfo(elevatorID int) (*types.Elevator, error)
	Run(elevatorID int) error
	FetchElevator(floorNo int, direction types.Direction) (int, error)
	RequestFloorFromElevator(elevatorID int, floorNo int) error
}
