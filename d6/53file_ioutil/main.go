package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("./main.go")
	if err == io.EOF {
		return
	}
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
