package golang

import "fmt"

type Director struct {
	builder BuildProcess
}

type VehicleProduct interface {
	Drive()
}

type BuildProcess interface {
	SetWheels(w int) BuildProcess
	SetSeats(s int) BuildProcess
	SetStructure(s string) BuildProcess
	Build() VehicleProduct
}

type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

func (v *Vehicle) Drive() {
	fmt.Printf("Driving %s with %d seats and %d wheels", v.Structure, v.Seats, v.Wheels)
}

type CarBuilder struct {
	v Vehicle
}

func (cb *CarBuilder) SetWheels(w int) BuildProcess {
	cb.v.Wheels = w
	return cb
}

func (cb *CarBuilder) SetSeats(s int) BuildProcess {
	cb.v.Seats = s
	return cb
}

func (cb *CarBuilder) SetStructure(s string) BuildProcess {
	cb.v.Structure = s
	return cb
}

func (cb *CarBuilder) Build() VehicleProduct {
	return &cb.v
}

func BuildCar() VehicleProduct {
	director := Director{}
	director.builder = &CarBuilder{}

	director.builder.SetWheels(4)
	director.builder.SetSeats(5)
	director.builder.SetStructure("CAR")

	car := director.builder.Build()
	return car
}
