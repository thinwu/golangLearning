package main

import (
	"encoding/json"
	"fmt"
)

// capitalized as to public, because json package needs to access the attributes.

type person struct {
	Name string
	Age  int
}

// or adding tag, means, in these package, the field will be displayed as the tag defined field.
type person2 struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age" db:"age" ini:"age"`
}

// json encode /decode
//
func main() {
	p1 := person{
		Name: "Daniel",
		Age:  37,
	}
	jStr, error := json.Marshal(p1)
	p2 := person2{
		Name: "Daniel",
		Age:  37,
	}
	jStr2, error2 := json.Marshal(p2)
	if error != nil {
		fmt.Println("json.Marshal error")
	} else {
		fmt.Printf("%#v\n", string(jStr))
	}
	if error2 != nil {
		fmt.Println("json.Marshal error")
	} else {
		fmt.Printf("%#v\n", string(jStr2))
	}
	// json to obj
	str := `{"Name":"Daniel","Age":37}`
	var p3 = new(person)
	err := json.Unmarshal([]byte(str), p3)
	if err != nil {
		fmt.Println("json.Unmarshal error")
		return
	}
	fmt.Printf("%#v\n", *p3)
}
