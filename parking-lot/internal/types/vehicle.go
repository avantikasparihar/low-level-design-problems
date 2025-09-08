package types

const (
	SmallCar VehicleType = "small-car"
	LargeCar VehicleType = "large-car"
	Bike     VehicleType = "bike"
)

type Vehicle struct {
	Type             VehicleType
	RegisteredNumber string
}

type VehicleType string
