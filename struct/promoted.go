package _struct

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

type Cat struct {
	Animal
}

func EmbeddedFoo() {
	c := Cat{}

	c.SetName("A")
	fmt.Printf("Name: %s\n", c.Name)

	c.Animal.Name = "B"
	fmt.Printf("Name: %s\n", c.Name)

	c.Name = "C"
	fmt.Printf("Name: %s\n", c.Name)
}
