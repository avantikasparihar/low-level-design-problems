package internal

import (
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal/types"
	"time"
)

const ratePerHr = 30

type BillingManager interface {
	GetBill(parkingID string) *types.ParkingBill
	PayBill(parkingID string)
}

type billingManager struct {
	parkingLot *types.ParkingLot
}

func (b billingManager) GetBill(parkingID string) *types.ParkingBill {
	return b.getBill(parkingID)
}

func (b billingManager) PayBill(parkingID string) {
	b.getBill(parkingID).Pay()
}

func (b billingManager) getBill(parkingID string) *types.ParkingBill {
	parking := b.parkingLot.GetParking(parkingID)
	elapsedTime := time.Since(parking.StartTime)
	bill := &types.ParkingBill{
		Amount: elapsedTime.Hours() * ratePerHr,
	}
	parking.Bill = bill
	return bill
}

func NewBillingManager(parkingLot *types.ParkingLot) BillingManager {
	return &billingManager{
		parkingLot: parkingLot,
	}
}
