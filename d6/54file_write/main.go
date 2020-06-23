package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// openfile write content

func main() {
	demo2()
	demo3()
}
func demo1() {
	fileObj, err := os.OpenFile("./aa.txt", os.O_TRUNC|os.O_CREATE, 0664)
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open file failed!\n")
		return
	}

	fileObj.Write([]byte("test\n"))
	fileObj.WriteString("writeString\n")
}
func demo2() {
	fileObj, err := os.OpenFile("./aa.txt", os.O_TRUNC|os.O_CREATE, 0664)
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open file failed!\n")
		return
	}
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("testing bufio")
	wr.Flush() // remember to flush the cache to push the content to the disk

}
func demo3() {
	str := "hello Dan"
	err := ioutil.WriteFile("./aa.txt", []byte(str), 0664)
	if err != nil {
		fmt.Println("write file failed, err")
		return
	}
}
