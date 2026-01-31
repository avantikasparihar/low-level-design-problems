package store

import (
	"sync"

	"github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"
)

type ElevatorStore struct {
	mut      sync.Mutex
	elevators []*types.Elevator
}

func (e *ElevatorStore) Get(id int) *types.Elevator {
	e.mut.Lock()
	defer e.mut.Unlock()
	return e.elevators[id]
}

func (e *ElevatorStore) List() []*types.Elevator {
	e.mut.Lock()
	defer e.mut.Unlock()
	return e.elevators
}

func (e *ElevatorStore) Update(id int, elevator *types.Elevator) {
	e.mut.Lock()
	defer e.mut.Unlock()
	e.elevators[id] = elevator
}

func NewElevatorStore(elevators []*types.Elevator) *ElevatorStore {
	return &ElevatorStore{
		elevators: elevators,
	}
}
