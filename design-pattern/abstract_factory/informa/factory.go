package informa

import (
	"go-learn-pattern/design-pattern/abstract_factory"
	"go-learn-pattern/design-pattern/abstract_factory/informa/product"
)

type Informa struct {
}

func (i *Informa) CreateChair() abstract_factory.Chair {
	return &product.BeanBag{}
}

func (i *Informa) CreateCoffeeTable() abstract_factory.CoffeeTable {
	return nil
}

func (i *Informa) CreateSofa() abstract_factory.Sofa {
	return nil
}
