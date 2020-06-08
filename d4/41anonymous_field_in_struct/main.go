package main

import "fmt"

type person struct {
	string // by default type will be the name
	int
	// then if there is another string will not be able to put into the struct.
}

func main() {
	p1 := person{
		"aa",
		111,
	}
	fmt.Println(p1.string)
}
