package golang

import (
	"errors"
	"fmt"
)

const (
	FamilyCarType  = 1
	LuxusryCarType = 2
)

type Car interface {
	NumDoors() int
}

type Motorbike interface {
	GetType() int
}

type Transport interface {
	NumWheels() int
	NumSeats() int
}

type TransportFactory interface {
	NewTransport(v int) (Transport, error)
}

type FamilyCar struct{}

func (fc *FamilyCar) NumDoors() int {
	return 5
}

func (fc *FamilyCar) NumWheels() int {
	return 4
}

func (fc *FamilyCar) NumSeats() int {
	return 5
}

type LuxuryCar struct{}

func (fc *LuxuryCar) NumDoors() int {
	return 2
}

func (fc *LuxuryCar) NumWheels() int {
	return 4
}

func (fc *LuxuryCar) NumSeats() int {
	return 5
}

type CarFactory struct{}

func (c *CarFactory) NewTransport(v int) (Car, error) {
	switch v {
	case FamilyCarType:
		return new(FamilyCar), nil
	case LuxusryCarType:
		return new(LuxuryCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("No car type for %d", v))
	}
}

type Bike struct{}

func (b *Bike) GetType() int {
	return 1
}

type MotorbikeFactory struct{}

func (mf *MotorbikeFactory) NewTransport(m int) (Motorbike, error) {
	return new(Bike), nil
}
