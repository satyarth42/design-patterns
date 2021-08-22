package head_first

import "fmt"

type IteratorItrfc interface {
	hasNext() bool
	next() *MenuItem
}

type DinerMenuIterator struct {
	items []*MenuItem
	position int
}
func (i *DinerMenuIterator) hasNext() bool {
	return i.items!=nil && i.position<len(i.items)
}
func (i *DinerMenuIterator) next() *MenuItem {
	i.position += 1
	return i.items[i.position-1]
}

type MenuItem struct {
	name string
	price float32
}

type IMenu interface {
	getIterator() IteratorItrfc
}

type DinerMenu struct {
	items []*MenuItem
}
func (m *DinerMenu) getIterator() IteratorItrfc {
	return &DinerMenuIterator{
		items: m.items,
	}
}
func (m *DinerMenu) addItem(item *MenuItem) {
	m.items = append(m.items, item)
}

type Waitress struct {
	dinerMenu DinerMenu
}
func (w *Waitress) printDinerMenu() {
	iterator := w.dinerMenu.getIterator()
	for iterator.hasNext() {
		fmt.Println(*iterator.next())
	}
}

func printMenu() {
	menu := DinerMenu{}
	menu.addItem(&MenuItem{"ITEM 1", 10.2})
	menu.addItem(&MenuItem{"ITEM 2", 1.0})

	waitress := Waitress{
		dinerMenu: menu,
	}

	waitress.printDinerMenu()
}