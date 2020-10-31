package golang

import (
	"errors"
	"fmt"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Blue  = 2
	Black = 3
)

type Shirt struct {
	Color int
	SKU   string
	Price float32
}

func (s Shirt) GetInfo() string {
	return fmt.Sprintf("Color is %d , price is %0.2f, SKU is %s", s.Color, s.Price, s.SKU)
}

func GetShirtCloner() ShirtCloner {
	return &ShirtsCache{}
}

var whiteShirt = &Shirt{
	Price: 17.0,
	SKU:   "EMPTY",
	Color: White,
}
var blueShirt = &Shirt{
	Price: 17.0,
	SKU:   "EMPTY",
	Color: Blue,
}

var blackShirt = &Shirt{
	Price: 17.0,
	SKU:   "EMPTY",
	Color: Black,
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(t int) (ItemInfoGetter, error) {
	switch t {
	case White:
		return *whiteShirt, nil
	case Blue:
		return *blueShirt, nil
	case Black:
		return *blackShirt, nil
	default:
		return nil, errors.New(fmt.Sprintf("Shirt cache for color %d not found", t))
	}
}
