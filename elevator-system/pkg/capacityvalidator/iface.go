package capacityvalidator

type CapacityValidator interface {
	Validate(elevatorID int) bool
}
