package elevatormanager

import (
	"fmt"
	"time"

	"github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/store"
	"github.com/avantikasparihar/low-level-design-problems/elevator-system/internal/types"
	"github.com/avantikasparihar/low-level-design-problems/elevator-system/pkg/capacityvalidator"
)

const maxFloors = 10

type elevatorMgr struct {
	elevator *store.ElevatorStore
	capacity capacityvalidator.CapacityValidator
}

func NewElevatorManager(elevator *store.ElevatorStore, capacity capacityvalidator.CapacityValidator) ElevatorManager {
	return &elevatorMgr{
		elevator: elevator,
		capacity: capacity,
	}
}

func (em *elevatorMgr) GetElevatorInfo(elevatorID int) (*types.Elevator, error) {
	el := em.elevator.Get(elevatorID)
	return el, nil
}

func (em *elevatorMgr) Run(elevatorID int) error {
	for {
		el := em.elevator.Get(elevatorID)
		// add a wait for 3 secs if it stops at current floor
		if _, found := el.NextFloors[el.CurrentFloor]; found {
			time.Sleep(3 * time.Second)
			delete(el.NextFloors, el.CurrentFloor)
			em.elevator.Update(elevatorID, el)
		}
		// move elevator to next floor if it has further requests
		if len(el.NextFloors) > 0 {
			err := em.moveElevator(el)
			if err != nil {
				return err
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (em *elevatorMgr) FetchElevator(floorNo int, direction types.Direction) (int, error){
	elList := em.elevator.List()
	// TODO: have different strategies for selecting elevator
	for _, el := range elList {
		if floorNo > el.CurrentFloor && el.CurrentDirection == types.DirectionUp {
			err := em.allocateFloor(floorNo, el)
			if err != nil {
				return -1, err
			}
			return el.ID, nil
		} else if floorNo < el.CurrentFloor && el.CurrentDirection == types.DirectionDown {
			err := em.allocateFloor(floorNo, el)
			if err != nil {
				return -1, err
			}
			return el.ID, nil
		}
	}
	// Choose the first elevator if the above conditions don't match
	targetEl := elList[0]
	err := em.allocateFloor(floorNo, targetEl)
	if err != nil {
		return -1, err
	}
	return targetEl.ID, nil
}

func (em *elevatorMgr) RequestFloorFromElevator(elevatorID int, floorNo int) error {
	targetEl := em.elevator.Get(elevatorID)
	return em.allocateFloor(floorNo, targetEl)
}

func (em *elevatorMgr) allocateFloor(floorNo int, elevator *types.Elevator) error {
	if _, found := elevator.NextFloors[floorNo]; found {
		return nil
	}
	if !em.capacity.Validate(elevator.ID) {
		return fmt.Errorf("weight limit exceeded for elevator %d", elevator.ID)
	}
	elevator.NextFloors[floorNo] = struct{}{}
	em.elevator.Update(elevator.ID, elevator)

	return nil
}

func (em *elevatorMgr) moveElevator(el *types.Elevator) error {
	switch el.CurrentDirection {
	case types.DirectionUp:
		if el.CurrentFloor < maxFloors {
			el.CurrentFloor++
		} else {
			el.CurrentDirection = types.DirectionDown
			el.CurrentFloor--
		}
	case types.DirectionDown:
		if el.CurrentFloor > 0 {
			el.CurrentFloor--
		} else {
			el.CurrentDirection = types.DirectionUp
			el.CurrentFloor++
		}
	default:
		return fmt.Errorf("invalid current direction %s", el.CurrentDirection)
	}
	em.elevator.Update(el.ID, el)
	return nil
}
