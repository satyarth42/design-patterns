package golang

import "fmt"

type Swimmer interface {
	Swim()
}

type Athlete interface {
	Train()
}

type Animal interface {
	Eat()
}

type CompositeSwimmer struct {
	Athlete
	Swimmer
}

type CompositeFish struct {
	Animal
	Swimmer
}

type AthleteImpl struct{}

func (a *AthleteImpl) Train() {
	fmt.Println("Training")
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

type AnimalImpl struct{}

func (a *AnimalImpl) Eat() {
	fmt.Println("Eating")
}

var fish = &CompositeFish{
	Animal:  &AnimalImpl{},
	Swimmer: &SwimmerImpl{},
}

var swimmer = &CompositeSwimmer{
	Athlete: &AthleteImpl{},
	Swimmer: &SwimmerImpl{},
}
