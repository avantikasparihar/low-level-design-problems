package main

import (
	"fmt"
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal"
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal/types"
	"sync"
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

	// 1. display details for floor 1
	displayFloorDetails(1)

	// 2. create a parking of type small-car on floor 0
	id := createSingleParking()

	// 3. display details for floor 0
	displayFloorDetails(0)

	time.Sleep(5 * time.Second)

	// 4. pay bill for parking
	bill := billingMgr.GetBill(id)
	fmt.Printf("bill created for parking %s: %+v\n", id, bill)
	billingMgr.PayBill(id)
	fmt.Printf("bill paid for parking: %s\n", id)

	// 5. exit parking
	err := parkingMgr.ExitParking(id)
	if err != nil {
		fmt.Printf("failed to exit parking: %v", err)
	}
	fmt.Printf("parking exited: %s\n", id)

	// 6. display details for floor 0
	displayFloorDetails(0)

	// 7. create multiple parkings concurrently
	_ = createMultipleParking()

	// 8. display details for floor 0
	displayFloorDetails(0)
}

func InitParkingLot() {
	floors := map[int]*types.ParkingFloor{
		0: {
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
		1: {
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
		2: {
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

func createSingleParking() string {
	vehicle := types.Vehicle{
		Type:             types.SmallCar,
		RegisteredNumber: "KA01 LN9999",
	}
	parking, err := parkingMgr.CreateParking(vehicle, 0)
	if err != nil {
		fmt.Printf("failed to park: %v", err)
		return ""
	}
	fmt.Printf("parking created: %+v\n", parking)
	return parking.ID
}

func createMultipleParking() []string {
	var (
		cnt = 11
		ids []string
	)
	var wg sync.WaitGroup
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createSingleParking()
			ids = append(ids, id)
		}()
	}
	wg.Wait()
	return ids
}
