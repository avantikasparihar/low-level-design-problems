package internal

const (
	OrderSizeSmall OrderSize = iota
	OrderSizeMedium
	OrderSizeLarge
)

const (
	OrderStateCreated OrderState = iota
	OrderStateDeposited
	OrderStatePickedUp
)

type Order struct {
	id                int
	item              string
	size              OrderSize
	customerId        int
	deliveryPartnerId int
	lockerId          int
	state             OrderState
}

type OrderSize int

type OrderState int

type CreateOrderRequest struct {
	Item string
	Size OrderSize
}

type OrderService interface {
	GetOrder(id int) Order
}

type orderSvc struct {
	orders []Order
}

func NewOrderSvc() OrderService {
	return &orderSvc{
		orders: []Order{
			{
				id:   1,
				size: OrderSizeSmall,
			},
		},
	}
}

func (os *orderSvc) GetOrder(id int) Order {
	for _, o := range os.orders {
		if o.id == id {
			return o
		}
	}
	return Order{}
}
