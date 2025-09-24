package elevatormanager

import "github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"

type elevatorMgr struct {
	elevators []*types.Elevator
}

func NewElevatorManager(elevators []*types.Elevator) ElevatorManager {
	return &elevatorMgr{
		elevators: elevators,
	}
}

func (em *elevatorMgr) FetchElevator(floorNo int, direction types.Direction) error {
	return nil
}

func (em *elevatorMgr) AllocateElevator(elevatorID int, floorNo int) error {
	// add a wait for 3 secs
	return nil
}
