package main

import (
	. "github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/elevatormanager"
	"github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"
)

/*
Entities:
	Elevator
		- ID
		- currentFloor
		- nextFloorsQueue
		- Direction <up, down>
		- Capacity

Interfaces:
	ElevatorManager
		- FetchElevator(floorNo int, direction string)
		- AllocateElevator(floorNo int) -> checks capacity(calls capacityMgr.Validate())
	CapacityValidator
		- Validate(elevatorID string) -> can have multiple impl depending upon the strategy(max weight, max count, etc)
*/

var elevators []*types.Elevator

func main() {
	InitElevators()
	elevatorMgr := NewElevatorManager(elevators)

	elevatorMgr.FetchElevator(1, types.DirectionUp)
}

func InitElevators() {
	elevators = []*types.Elevator{
		{
			ID:               0,
			CurrentFloor:     0,
			NextFloors:       make(map[int]struct{}),
			CurrentDirection: types.DirectionUp,
			Capacity:         0,
		},
		{
			ID:               1,
			CurrentFloor:     0,
			NextFloors:       make(map[int]struct{}),
			CurrentDirection: types.DirectionUp,
			Capacity:         0,
		},
	}
}
