package ikea

import (
	"go-learn-pattern/design-pattern/abstract_factory"
	"go-learn-pattern/design-pattern/abstract_factory/ikea/product"
)

type Ikea struct {
}

func (i *Ikea) CreateChair() abstract_factory.Chair {
	return &product.Leifarne{}
}

func (i *Ikea) CreateCoffeeTable() abstract_factory.CoffeeTable {
	return &product.VITTSJO{}
}

func (i *Ikea) CreateSofa() abstract_factory.Sofa {
	return &product.HEMLINGBY{}
}
