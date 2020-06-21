package main

import (
	"fmt"
	"io"
	"os"
)

// open file

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed.")
		fmt.Printf("111\n%#v\n222", file)
		// so when we use *file to get its value, it will panic
		// fmt.Printf("111\n%#v\n222", *file) // because you are trying to type a nil's value
		return
	}

	defer file.Close() // if there is no file opened. *file is nil, we can test it with fmt.printf("%#v", file)
	// this func must have nil protection
	// defer will make this func run between return's two opts(1st is to set return value, 2nd is the real return), if there is a return will be exec'ed
	// so the consequence is
	// 1 set return value
	// 2 defer
	// 3 return

	var tmp = [128]byte{}
	for {
		n, err := file.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("//read to end")
			return
		}
		if err != nil {
			fmt.Printf("Read failed err %v\n", err)
		}
		//fmt.Printf("read %d bytes\n", n)
		fmt.Print(string(tmp[:n])) // [m:n] => including m, excluding n
		// if n < 128 {
		// 	return
		// }
	}

}
