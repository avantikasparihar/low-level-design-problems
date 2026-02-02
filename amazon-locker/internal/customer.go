package internal

type customer struct {
	id     int
	name   string
	orders []int
}

func NewCustomer() CustomerOrderManager {
	return &customer{}
}

type CustomerOrderManager interface {
	CreateOrder(req *CreateOrderRequest) error
	GetOrderAccessCode(orderId int) (int, error)
	PickupOrder(orderId, accessCode int) error
}

func (c *customer) CreateOrder(req *CreateOrderRequest) error {
	return nil
}

func (c *customer) GetOrderAccessCode(orderId int) (int, error) {
	// return access code from order details
	return -1, nil
}

func (c *customer) PickupOrder(orderId, accessCode int) error {
	// validate access code against order id
	// update order state to picked-up
	return nil
}
