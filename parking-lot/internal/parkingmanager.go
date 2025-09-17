package internal

import (
	"fmt"
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal/types"
	"math/rand/v2"
	"time"
)

type ParkingManager interface {
	CreateParking(v types.Vehicle, floorNo int) (*types.VehicleParking, error)
	GetVacancyByVehicleType(vt types.VehicleType) int
	GetVacancyByFloor(floorNo int) map[types.VehicleType]int
	ExitParking(parkingID string) error
}

type parkingMgr struct {
	parkingLot *types.ParkingLot
}

func (p *parkingMgr) CreateParking(v types.Vehicle, floorNo int) (*types.VehicleParking, error) {
	floor := p.parkingLot.GetFloor(floorNo)
	err := floor.BookSpot(v.Type)
	if err != nil {
		return nil, err
	}
	p.parkingLot.UpdateFloor(floor)
	newParking := &types.VehicleParking{
		ID:        fmt.Sprintf("id-%d", rand.IntN(100)),
		StartTime: time.Now(),
		Vehicle:   v,
		FloorNo:   floorNo,
	}
	p.parkingLot.CreateParking(newParking)
	return newParking, nil
}

func (p *parkingMgr) GetVacancyByVehicleType(vt types.VehicleType) int {
	var res = 0
	for _, floor := range p.parkingLot.Floors {
		if count, ok := floor.Capacity.Available[vt]; ok {
			res += count
		}
	}
	return res
}

func (p *parkingMgr) GetVacancyByFloor(floorNo int) map[types.VehicleType]int {
	floor := p.parkingLot.GetFloor(floorNo)
	if floor != nil {
		return floor.Capacity.Available
	}
	return nil
}

func (p *parkingMgr) ExitParking(parkingID string) error {
	parking := p.parkingLot.GetParking(parkingID)
	if !parking.IsBillPaid() {
		return fmt.Errorf("bill not paid for parking %s", parkingID)
	}
	p.parkingLot.DeleteParking(parkingID)
	floor := p.parkingLot.GetFloor(parking.FloorNo)
	floor.FreeSpot(parking.Vehicle.Type)
	p.parkingLot.UpdateFloor(floor)
	return nil
}

func NewParkingMgr(parkingLot *types.ParkingLot) ParkingManager {
	return &parkingMgr{
		parkingLot: parkingLot,
	}
}
