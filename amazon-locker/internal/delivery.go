package internal

var DeliveryMgr DeliveryManager

type DeliveryPartner struct {
	id     int
	name   string
	orders []int
}

type PartnerOrderManager interface {
	DepositOrder(orderId, lockerId int) error
}

func NewDeliveryPartner() PartnerOrderManager {
	return &DeliveryPartner{
		id: 1,
	}
}

func (dp *DeliveryPartner) DepositOrder(orderId int, lockerId int) error {
	// check capacity at locker
	// change order state to deposited
	return nil
}

type DeliveryManager interface {
	AssignDelivery(orderId, deliveryPartnerId int) error
}

type deliveryManager struct {
	orders   []Order
	partners []DeliveryPartner
	lockers  []Locker
}

// can be a singleton
func InitDeliveryManager() {
	if DeliveryMgr != nil {
		return
	}
	DeliveryMgr = &deliveryManager{
		orders:   make([]Order, 0),
		partners: make([]DeliveryPartner, 0),
		lockers: []Locker{
			{
				id: 1,
				capacityAvailable: map[OrderSize]int{
					OrderSizeSmall:  1,
					OrderSizeMedium: 1,
					OrderSizeLarge:  1,
				},
			},
			{
				id: 2,
				capacityAvailable: map[OrderSize]int{
					OrderSizeSmall:  1,
					OrderSizeMedium: 1,
					OrderSizeLarge:  1,
				},
			},
		},
	}
}

func (dm *deliveryManager) AssignDelivery(orderId, deliveryPartnerId int) error {
	for _, p := range dm.partners {
		if p.id == deliveryPartnerId {
			p.orders = append(p.orders, orderId)
			break
		}
	}
	return nil
}
