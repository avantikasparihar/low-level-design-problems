package main

import (
	"fmt"
	"os"
	"time"
	"sync"

	. "github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/elevatormanager"
	"github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/store"
	"github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"
	"github.com/avantikasparihar/low-level-design-problems/elevator-system/pkg/capacityvalidator"
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

var (
	elevators   []*types.Elevator
	elevatorMgr ElevatorManager
)

func main() {
	fmt.Println("Elevator System")
	InitElevators()
	elevatorMgr = NewElevatorManager(store.NewElevatorStore(elevators), capacityvalidator.NewCapacityValidatorByWeight())
	
	// Start background job to display elevator info
	var wg sync.WaitGroup
	wg.Add(1)
	go displayElevatorInfo()
	wg.Add(1)
	go func() {
		err := elevatorMgr.Run(0)
		if err != nil {
			fmt.Printf("failed running elevator: %v", err)
			os.Exit(1)
		}
	}()
	wg.Add(1)
	go func() {
		err := elevatorMgr.Run(1)
		if err != nil {
			fmt.Printf("failed running elevator: %v", err)
			os.Exit(1)
		}
	}()
	defer wg.Wait()

	// Fetch elevator
	id, err := elevatorMgr.FetchElevator(2, types.DirectionUp)
	if err != nil {
		fmt.Printf("failed fetching elevator: %v", err)
		os.Exit(1)
	}
	// Request floor to elevator
	err = elevatorMgr.RequestFloorFromElevator(id, 7)
	if err != nil {
		fmt.Printf("failed allocating elevator: %v", err)
		os.Exit(1)
	}

	// Wait for a few seconds
	time.Sleep(3 * time.Second)

	// Fetch another elevator
	id, err = elevatorMgr.FetchElevator(1, types.DirectionUp)
	if err != nil {
		fmt.Printf("failed fetching elevator: %v", err)
		os.Exit(1)
	}
	// Request floor to elevator
	err = elevatorMgr.RequestFloorFromElevator(id, 9)
	if err != nil {
		fmt.Printf("failed allocating elevator: %v", err)
		os.Exit(1)
	}
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

func displayElevatorInfo() {
	for {
		fmt.Println("*****************************")
		for _, el := range elevators {
			elevator, err := elevatorMgr.GetElevatorInfo(el.ID)
			if err != nil {
				fmt.Printf("failed getting elevator info: %v", err)
				os.Exit(1)
			}
			fmt.Printf("Elevator %d\n", elevator.ID)
			fmt.Printf("floor: %d, direction: %s\n", elevator.CurrentFloor, elevator.CurrentDirection)
		}
		time.Sleep(1 * time.Second)
	}
}
