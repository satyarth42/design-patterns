package head_first

import "fmt"

type MenuComponent interface {
	getName() string
	getPrice() float32
	print()
	add(component MenuComponent)
	getChild(i int) MenuComponent
	getType() string
}
type emptyMenuComponent struct {

}

func (e *emptyMenuComponent) getName() string {
	panic("implement me")
	return ""
}

func (e *emptyMenuComponent) getPrice() float32 {
	panic("implement me")
	return 0
}

func (e *emptyMenuComponent) print() {
	panic("implement me")
}

func (e *emptyMenuComponent) add(component MenuComponent) {
	panic("implement me")
}

func (e *emptyMenuComponent) getChild(i int) MenuComponent {
	panic("implement me")
	return nil
}
func (e *emptyMenuComponent) getType(i int) string {
	panic("implement me")
	return ""
}

type FoodItem struct {
	emptyMenuComponent
	name string
	price float32
}
func (f *FoodItem) getName() string {
	return f.name
}
func (f *FoodItem) getPrice() float32 {
	return f.price
}
func (f *FoodItem) print() {
	fmt.Printf("name: %s, price: %f\n",f.name, f.price)
}
func (f *FoodItem) getType() string {
	return "food item"
}

type SubMenu struct {
	items []MenuComponent
	name string
}

func (s *SubMenu) getName() string {
	return s.name
}
func (s *SubMenu) getPrice() float32 {
	panic("implement me")
	return 0
}
func (s *SubMenu) print() {
	fmt.Printf("name: %s\n",s.getName())
	for _, item := range s.items {
		item.print()
	}
}
func (s *SubMenu) add(component MenuComponent) {
	s.items = append(s.items,component)
}
func (s *SubMenu) getChild(i int) MenuComponent {
	return s.items[i]
}
func (s *SubMenu) getType() string {
	return "sub menu"
}
type CompMenu struct {
	items []MenuComponent
}
