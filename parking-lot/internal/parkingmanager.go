package internal

import (
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal/types"
	"time"
)

type ParkingManager interface {
	CreateParking(v types.Vehicle, floorNo int) *types.VehicleParking
	GetVacancyByVehicleType(vt types.VehicleType) int
	GetVacancyByFloor(floorNo int) map[types.VehicleType]int
}

type parkingMgr struct {
	parkingLot *types.ParkingLot
}

func (p parkingMgr) CreateParking(v types.Vehicle, floorNo int) *types.VehicleParking {
	for _, floor := range p.parkingLot.Floors {
		if floor.FloorNo == floorNo {
			floor.BookSpot(v.Type)
			break
		}
	}
	vp := &types.VehicleParking{
		ID:        "",
		StartTime: time.Now(),
		Vehicle:   v,
		FloorNo:   floorNo,
	}
	p.parkingLot.Parkings = append(p.parkingLot.Parkings, vp)
	return vp
}

func (p parkingMgr) GetVacancyByVehicleType(vt types.VehicleType) int {
	var res = 0
	for _, floor := range p.parkingLot.Floors {
		if count, ok := floor.Capacity.Available[vt]; ok {
			res += count
		}
	}
	return res
}

func (p parkingMgr) GetVacancyByFloor(floorNo int) map[types.VehicleType]int {
	for _, floor := range p.parkingLot.Floors {
		if floor.FloorNo == floorNo {
			return floor.Capacity.Available
		}
	}
	return nil
}

func NewParkingMgr(parkingLot *types.ParkingLot) ParkingManager {
	return &parkingMgr{
		parkingLot: parkingLot,
	}
}
