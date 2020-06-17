package main

import "fmt"

type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// one structure implements multiple interfaces
func (c *cat) move() {
	fmt.Println("cat moves")
}

func (c *cat) eat(food string) {
	fmt.Printf("cat eats %v\n", food)
}

// interface nesting
type animal interface {
	mover
	eater
}

func main() {
	c1 := cat{
		name: "Tom",
		feet: 4,
	}
	var a1 animal
	a1 = &c1
	a1.move()
	a1.eat("mouse")

}
