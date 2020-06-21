package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed.")
		return
	}
	defer file.Close()
	// create a reader
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read line failed")
		}

		fmt.Print(line)
	}

}
