package internal

import (
	"errors"
)

type Locker struct {
	id                int
	location          string
	capacityAvailable map[OrderSize]int
	compartmentList   []*Compartment
}

type Compartment struct {
	id          int
	size        OrderSize
	occupied    bool
	accessToken int
	orderId     int
}

func (l *Locker) AllocateCompartment(size OrderSize, orderId int) (int, error) {
	if l.capacityAvailable[size] == 0 {
		return -1, errors.New("no compartments available, try after some time")
	}

	for _, comp := range l.compartmentList {
		if comp.occupied || comp.size != size {
			continue
		}
		comp.occupied = true
		comp.accessToken = generateAccessToken()
		comp.orderId = orderId
		// update capacity on the locker
		l.capacityAvailable[comp.size]--
		// todo: update locker info on the order details
		return comp.id, nil
	}

	return -1, errors.New("unable to find compartment")
}

func (l *Locker) UnlockCompartment(id, accessToken int) error {
	for _, comp := range l.compartmentList {
		if comp.id != id {
			continue
		}
		if comp.accessToken != accessToken {
			return errors.New("access token doesn't match for the compartment")
		}
		return nil
	}

	return errors.New("compartment id not found")
}

func generateAccessToken() int {
	// generate a random token
	return 123456
}
