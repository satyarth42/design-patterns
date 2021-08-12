package head_first

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type NYPizza struct{}

func (p *NYPizza) Prepare() {

}
func (p *NYPizza) Bake() {

}
func (p *NYPizza) Cut() {

}
func (p *NYPizza) Box() {

}

type IPizzaStore interface {
	createPizza(pizzaType string) IPizza
	OrderPizza(pizzaType string) IPizza
}

type PizzaStore struct{}

func (ps *PizzaStore) MakePizza(pizza IPizza) IPizza {

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

type NYPizzaStore struct {
	PizzaStore
}

func (ps *NYPizzaStore) createPizza(pizzaType string) IPizza {
	switch pizzaType {
	case "ny":
		return &NYPizza{}
	default:
		return nil
	}
}
func (ps *NYPizzaStore) OrderPizza(pizzaType string) IPizza {
	pizza := ps.createPizza(pizzaType)

	ps.PizzaStore.MakePizza(pizza)

	return pizza
}

func orderPizza() IPizza {
	store := NYPizzaStore{PizzaStore{}}
	return store.OrderPizza("ny")
}
