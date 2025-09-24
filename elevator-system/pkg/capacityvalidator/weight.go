package capacityvalidator

const maxAllowedWeightInKg = 500

type capacityValidatorByWeight struct {}

func (cw *capacityValidatorByWeight) Validate(elevatorID int) bool {
	if getCurrentWeight(elevatorID) <= maxAllowedWeightInKg {
		return true
	}
	return false
}

func getCurrentWeight(elevatorID int) int {
	// TODO: must fetch weight from sensors
	return 0
}

func NewCapacityValidatorByWeight() CapacityValidator {
	return &capacityValidatorByWeight{}
}
