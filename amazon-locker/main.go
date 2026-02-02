package main

import (
	. "github.com/avantikasparihar/low-level-design-problems/amazon-locker/internal"
	"log"
)

/*
Entities:
	Order:
		- id: string
		- size: enum("SMALL", "MEDIUM", "LARGE")
		- customerId: int
		- deliveryPartnerId: int
		- state: enum("CREATED", "DEPOSITED", "PICKED-UP")
		- customerAccessCode: int
		- locker: Locker

	Locker:
		- id: int
		- location: string
		- capacityAvailable: map[size]int

	customer:
		- id: int
		- name: string
		- orders: []int

	DeliveryPartner:
		- id: int
		- name: string
		- orders: []int

Interfaces:
	CustomerOrderManager:
		- CreateOrder(order OrderRequest) error
		- GetOrderAccessCode(orderId int) int

	PartnerOrderManager:
		- DepositOrder(orderId int, lockerId int) error

	DeliveryManager:
		- AssignDelivery(orderId int, deliveryPartnerId int) error
*/

func main() {
	lockerMgr := GetLockerMgr()

	compId, err := lockerMgr.Deposit("locker-1", 1)
	if err != nil {
		log.Fatalln(err)
	}

	err = lockerMgr.Pickup("locker-1", compId, 123456)
	if err != nil {
		log.Fatalln(err)
	}
}
