package internal

import (
	"github.com/avantikasparihar/low-level-design-problems/parking-lot/internal/types"
	"time"
)

type BillingManager interface {
	GetBill(parkingID string) *types.ParkingBill
	PayBill(parkingID string)
}

type billingManager struct {
	parkingLot *types.ParkingLot
}

func (b billingManager) GetBill(parkingID string) *types.ParkingBill {
	parking := b.parkingLot.GetParking(parkingID)
	elapsedTime := time.Since(parking.StartTime)
	return &types.ParkingBill{
		Amount: elapsedTime.Hours() * 30,
	}
}

func (b billingManager) PayBill(parkingID string) {
	parking := b.parkingLot.GetParking(parkingID)
	parking.Bill.Pay()
}

func NewBillingManager() BillingManager {
	return &billingManager{}
}
