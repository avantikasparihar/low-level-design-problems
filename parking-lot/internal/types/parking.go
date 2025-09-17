package types

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	mut      sync.Mutex
	Floors   map[int]*ParkingFloor
	Parkings map[string]*VehicleParking
}

func (pl *ParkingLot) GetFloor(no int) *ParkingFloor {
	pl.mut.Lock()
	defer pl.mut.Unlock()
	if floor, ok := pl.Floors[no]; ok {
		return floor
	}
	return nil
}

func (pl *ParkingLot) UpdateFloor(floor *ParkingFloor) {
	pl.mut.Lock()
	defer pl.mut.Unlock()
	pl.Floors[floor.FloorNo] = floor
}

func (pl *ParkingLot) CreateParking(parking *VehicleParking) {
	pl.mut.Lock()
	defer pl.mut.Unlock()
	pl.Parkings[parking.ID] = parking
}

func (pl *ParkingLot) GetParking(id string) *VehicleParking {
	return pl.Parkings[id]
}

func (pl *ParkingLot) UpdateParking(parking *VehicleParking) {
	pl.Parkings[parking.ID] = parking
}

func (pl *ParkingLot) DeleteParking(id string) {
	pl.mut.Lock()
	defer pl.mut.Unlock()
	if pl.Parkings[id] != nil {
		pl.Parkings[id].EndTime = time.Now()
	}
}

func NewParkingLot(floors map[int]*ParkingFloor) *ParkingLot {
	return &ParkingLot{
		Floors:   floors,
		Parkings: make(map[string]*VehicleParking),
	}
}

type ParkingFloor struct {
	FloorNo  int
	Capacity Capacity
}

type VehicleParking struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
	Vehicle   Vehicle
	FloorNo   int
	Bill      *ParkingBill
}

func (vp *VehicleParking) IsBillPaid() bool {
	if vp.Bill == nil {
		return false
	}
	return vp.Bill.Paid
}

type Capacity struct {
	Total     map[VehicleType]int
	Available map[VehicleType]int
}

func (pf *ParkingFloor) BookSpot(vt VehicleType) error {
	count, found := pf.Capacity.Available[vt]
	if !found {
		return errors.New(fmt.Sprintf("Vehicle type %s not available", vt))
	}
	if count <= 0 {
		return fmt.Errorf("%s is fully booked", vt)
	}
	pf.Capacity.Available[vt] -= 1
	return nil
}

func (pf *ParkingFloor) FreeSpot(vt VehicleType) {
	if _, ok := pf.Capacity.Available[vt]; ok {
		pf.Capacity.Available[vt] += 1
	}
}

type ParkingBill struct {
	Amount float64
	Paid   bool
}

func (pb *ParkingBill) Pay() {
	pb.Paid = true
}
