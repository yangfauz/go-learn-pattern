package product

import "go-learn-pattern/design-pattern/abstract_factory"

type VITTSJO struct{}

func (v VITTSJO) Price() float64 {
	return 899000
}

func (v VITTSJO) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 80,
		Width:  78,
		Height: 40,
	}
}

func (v VITTSJO) IsFoldable() bool {
	return false
}
