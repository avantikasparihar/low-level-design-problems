package internal

// LockerMgr is a singleton for managing all lockers within the system
var LockerMgr LockerManager

type LockerManager interface {
	Deposit(lockerName string, orderId int) (int, error)
	Pickup(lockerName string, compId, accessToken int) error
}

type lockerMgr struct {
	lockers  map[string]*Locker
	orderSvc OrderService
}

func GetLockerMgr() LockerManager {
	if LockerMgr != nil {
		return LockerMgr
	}
	LockerMgr = &lockerMgr{
		lockers: map[string]*Locker{
			"locker-1": {
				id: 1,
				capacityAvailable: map[OrderSize]int{
					OrderSizeSmall:  1,
					OrderSizeMedium: 1,
					OrderSizeLarge:  1,
				},
				compartmentList: []*Compartment{
					{
						id:   1,
						size: OrderSizeSmall,
					},
					{
						id:   2,
						size: OrderSizeMedium,
					},
					{
						id:   3,
						size: OrderSizeLarge,
					},
				},
			},
			"locker-2": {
				id: 2,
				capacityAvailable: map[OrderSize]int{
					OrderSizeSmall: 1,
				},
				compartmentList: []*Compartment{
					{
						id:   1,
						size: OrderSizeSmall,
					},
				},
			},
		},
		orderSvc: NewOrderSvc(),
	}
	return LockerMgr
}

func (lm *lockerMgr) Deposit(lockerName string, orderId int) (int, error) {
	order := lm.orderSvc.GetOrder(orderId)

	locker := lm.lockers[lockerName]
	compId, err := locker.AllocateCompartment(order.size, orderId)
	if err != nil {
		return -1, err
	}

	return compId, nil
}

func (lm *lockerMgr) Pickup(lockerName string, compId, accessToken int) error {
	locker := lm.lockers[lockerName]
	err := locker.UnlockCompartment(compId, accessToken)
	if err != nil {
		return err
	}
	return nil
}
