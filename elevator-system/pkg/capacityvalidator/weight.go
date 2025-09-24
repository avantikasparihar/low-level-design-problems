package capacityvalidator

type capacityValidatorByWeight struct {
}

func (cw *capacityValidatorByWeight) Validate() bool {
	return true
}

func NewCapacityValidatorByWeight() CapacityValidator {
	return &capacityValidatorByWeight{}
}
