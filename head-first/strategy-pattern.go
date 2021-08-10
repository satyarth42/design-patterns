package head_first

import "fmt"

type IDuck interface {
	Swim()
	Display()
	SetFlyBehavior(flyBehavior IFly)
	SetQuackBehavior(quackBehavior IQuack)
	PerformQuack()
	PerformFly()
}

type IFly interface {
	Fly()
}

type IQuack interface {
	Quack()
}

type Duck struct {
	FlyBehaviour IFly
	QuackBehavior IQuack
}

func (d *Duck) Swim() {
	fmt.Println("Duck Swims")
}

func (d *Duck) Display() {
	fmt.Println("Duck Display")
}

func (d *Duck) SetFlyBehavior(flyBehavior IFly) {
	d.FlyBehaviour = flyBehavior
}

func (d *Duck) SetQuackBehavior(quackBehavior IQuack) {
	d.QuackBehavior = quackBehavior
}

func (d *Duck) PerformQuack() {
	d.QuackBehavior.Quack()
}

func (d *Duck) PerformFly() {
	d.FlyBehaviour.Fly()
}

type Flightless struct {}
func (f *Flightless) Fly() {
	fmt.Println("cant fly")
}

type Silent struct {}
func (s *Silent) Quack() {
	fmt.Println("cannot quack")
}

type WoodenDuck struct {
	Duck
}

func CreateWoodenDuck() IDuck {
	duck := WoodenDuck{}

	flightless := &Flightless{}
	silent := &Silent{}

	duck.SetFlyBehavior(flightless)
	duck.SetQuackBehavior(silent)

	duck.Display()
	duck.Swim()

	duck.PerformFly()
	duck.PerformQuack()

	return &duck
}