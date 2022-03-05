package oop

import "fmt"

type Information interface {
	General()
	Attributes()
}
type Item struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func (itm Item) General() {
	fmt.Printf("\n%s", itm.Name)
	fmt.Printf("\n%s\n", itm.Description)
	fmt.Println(itm.Price)
}
func (itm Item) Attributes() {
	fmt.Println(itm.Weight)
}

type Mobile struct {
	Item
	DisplayFeatures   []string
	ProcessorFeatures []string
}

func (mob Mobile) Attributes() {
	mob.Item.Attributes()
	fmt.Println("\nDisplay Features:")
	for _, key := range mob.DisplayFeatures {
		fmt.Println(key)
	}
	fmt.Println("\nProcessor Features:")
	for _, key := range mob.ProcessorFeatures {
		fmt.Println(key)
	}
}
