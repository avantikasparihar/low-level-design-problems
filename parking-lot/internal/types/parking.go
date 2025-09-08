package types

import "time"

type ParkingLot struct {
	Floors   []*ParkingFloor
	Parkings []*VehicleParking
}

func (pl *ParkingLot) GetParking(id string) *VehicleParking {
	for _, parking := range pl.Parkings {
		if parking.ID == id {
			return parking
		}
	}
	return nil
}

func (pl *ParkingLot) GetFloor(no int) *ParkingFloor {
	for _, floor := range pl.Floors {
		if floor.FloorNo == no {
			return floor
		}
	}
	return nil
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{}
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

type Capacity struct {
	Total     map[VehicleType]int
	Available map[VehicleType]int
}

func (pf *ParkingFloor) BookSpot(vt VehicleType) {
	if _, ok := pf.Capacity.Available[vt]; ok {
		pf.Capacity.Available[vt] -= 1
	}
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
