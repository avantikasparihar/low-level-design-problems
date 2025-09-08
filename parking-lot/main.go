package main

import (
	"fmt"
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal"
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal/types"
	"time"
)

/*
Entities:
	- Vehicle
		- Type
	- ParkingFloor
		- FloorNo
		- Capacity
			- Total
			- Available
	- VehicleParking
		- StartTime
		- Vehicle
*/

/*
Interfaces:
	- ParkingManager
		- CreateParking
		- GetParking
	- BillingManager
		- GetBill
*/

/*
Requirements:
1. entity ParkingLot which has a list of parking floors
2. parkings should happen parallely and not sequentially
3.
4.
5.
6.
7. Get Vacancy func- GetVacancyByVehicleType()
8. Entity Capacity in ParkingFloor has a map of vehicle type to quantity
9.
10.
11. Get Vacancy func- GetVacancyByFloor()
12. Interface- BillingManager
*/

var (
	parkingLot *types.ParkingLot
	parkingMgr internal.ParkingManager
	billingMgr internal.BillingManager
)

func main() {
	InitParkingLot()

	// display details for floor 1
	displayFloorDetails(1)

	// create a parking of type small-car
	vehicle := types.Vehicle{
		Type:             types.SmallCar,
		RegisteredNumber: "KA01 LN9999",
	}
	parking, err := parkingMgr.CreateParking(vehicle, 0)
	if err != nil {
		fmt.Printf("failed to park: %v", err)
	}
	fmt.Printf("parking created: %+v\n", parking)

	// display details for floor 0
	displayFloorDetails(0)

	time.Sleep(5 * time.Second)

	// pay bill for parking
	bill := billingMgr.GetBill(parking.ID)
	fmt.Printf("bill created for parking %s: %+v\n", parking.ID, bill)
	billingMgr.PayBill(parking.ID)
	fmt.Printf("bill paid for parking: %s\n", parking.ID)

	// exit parking
	err = parkingMgr.ExitParking(parking.ID)
	if err != nil {
		fmt.Printf("failed to exit parking: %v", err)
	}
	fmt.Printf("parking exited: %+v\n", parking)

	// display details for floor 0
	displayFloorDetails(0)
}

func InitParkingLot() {
	floors := []*types.ParkingFloor{
		{
			FloorNo: 0,
			Capacity: types.Capacity{
				Total: map[types.VehicleType]int{
					types.SmallCar: 10,
					types.LargeCar: 5,
					types.Bike:     10,
				},
				Available: map[types.VehicleType]int{
					types.SmallCar: 10,
					types.LargeCar: 5,
					types.Bike:     10,
				},
			},
		},
		{
			FloorNo: 1,
			Capacity: types.Capacity{
				Total: map[types.VehicleType]int{
					types.SmallCar: 10,
					types.LargeCar: 5,
					types.Bike:     10,
				},
				Available: map[types.VehicleType]int{
					types.SmallCar: 10,
					types.LargeCar: 5,
					types.Bike:     10,
				},
			},
		},
		{
			FloorNo:  2,
			Capacity: types.Capacity{},
		},
	}
	parkingLot = types.NewParkingLot(floors)
	parkingMgr = internal.NewParkingMgr(parkingLot)
	billingMgr = internal.NewBillingManager(parkingLot)
}

func displayFloorDetails(floorNo int) {
	fmt.Printf("Floor %d\n", floorNo)
	floor := parkingLot.GetFloor(floorNo)
	for vt, count := range floor.Capacity.Total {
		fmt.Printf("type: %s, total: %d, available: %d\n", vt, count, floor.Capacity.Available[vt])
	}
}
