package head_first

import "fmt"

type DuckItrfc interface {
	quack()
	fly()
}

type MallardDuck struct {}

func (d *MallardDuck) quack()  {
	fmt.Println("duck quacking")
}
func (d *MallardDuck) fly()  {
	fmt.Println("duck flying")
}

type ITurkey interface {
	gobble()
	fly()
}

type Turkey struct {}

func (t *Turkey) gobble()  {
	fmt.Println("turkey quacking")
}
func (t *Turkey) fly()  {
	fmt.Println("turkey flying short distance")
}

type DuckTurkeyAdapter struct {
	turkey ITurkey
}
func (a *DuckTurkeyAdapter) quack() {
	a.turkey.gobble()
}
func (a *DuckTurkeyAdapter) fly() {
	for i:=0;i<5;i++ {
		a.turkey.fly() //turkey flies short distance so making it fly many times
	}
}
