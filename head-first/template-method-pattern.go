package head_first

import "fmt"

type IBeverageTemplate interface {
	boilWater()
	brew()
	pourInCup()
	addCondiments()
	prepareBeverage()
}

type Beverage struct {
	brewFunc func()
	addCondimentsFunc func()
}
func (b *Beverage) boilWater() {
	fmt.Println("water boiling")
}
func (b *Beverage) pourInCup() {
	fmt.Println("pouring beverage in cup")
}
func (b *Beverage) brew() {
	b.brewFunc()
}
func (b *Beverage) addCondiments() {
	b.addCondimentsFunc()
}
func (b *Beverage) prepareBeverage() {
	b.boilWater()
	b.brew()
	b.pourInCup()
	b.addCondiments()
}

type Tea struct {
	Beverage
}
func (t *Tea) brew() {
	fmt.Println("brewing tea")
}
func (t *Tea) addCondiments() {
	fmt.Println("adding tea condiments")
}

func GetNewTea() IBeverageTemplate {
	t := &Tea{}
	t.Beverage.addCondimentsFunc = t.addCondiments
	t.Beverage.brewFunc = t.brew
	return t
}