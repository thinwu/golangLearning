package main

import "fmt"

type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Printf("%s 会动\n", (*a).name)
}

type dog struct {
	leg uint8
	animal
}

func (d *dog) bark() {
	fmt.Printf("%s 汪汪汪\n", (*d).name)
}
func main() {
	d1 := dog{
		leg: 4,
		animal: animal{
			name: "dog",
		},
	}
	d1.move()            // from animal
	d1.bark()            // from dog
	fmt.Println(d1.name) // from animal
	fmt.Println(d1.leg)  // from dog
}
