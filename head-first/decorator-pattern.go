package head_first

type IBeverage interface {
	Cost() float32
	GetDescription() string
}

type BaseInfo struct {
	cost        float32
	description string
}

func (b *BaseInfo) Cost() float32 {
	return b.cost
}
func (b *BaseInfo) GetDescription() string {
	return b.description
}

type Condiments struct {
	Beverage IBeverage
	BaseInfo
}
func (c *Condiments) Cost() float32 {
	if c.Beverage!=nil {
		return c.BaseInfo.Cost()+c.Beverage.Cost()
	}
	return c.BaseInfo.Cost()
}
func (c *Condiments) Description() string {
	if c.Beverage!=nil {
		return c.BaseInfo.GetDescription()+", "+c.Beverage.GetDescription()
	}
	return c.BaseInfo.GetDescription()
}

func getDecaf() IBeverage {
	 return &BaseInfo{
		cost: 0.99,
		description: "decaf",
	}
}
func getEspresso() IBeverage {
	return &BaseInfo{
		cost: 0.8,
		description: "espresso",
	}
}

func addMocha(beverage IBeverage) IBeverage {
	return &Condiments{
		Beverage: beverage,
		BaseInfo:BaseInfo{
			cost: 0.2,
			description:"mocha",
		},
	}
}

func addWhip(beverage IBeverage) IBeverage {
	return &Condiments{
		Beverage: beverage,
		BaseInfo:BaseInfo{
			cost: 0.1,
			description:"whip",
		},
	}
}

func getWhipDoubleMochaDecaf() IBeverage {
	decaf := getDecaf()

	mochaDecafCoffee := addMocha(decaf)

	doubleMochaDecaf := addMocha(mochaDecafCoffee)

	whipDoubleMochaDecaf := addWhip(doubleMochaDecaf)

	return whipDoubleMochaDecaf
}