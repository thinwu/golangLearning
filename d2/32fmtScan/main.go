package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		name  string
		class string
		age   int
	)
	fmtScan(&name, &class, &age)
	useBufio()
}

func fmtScan(name *string, class *string, age *int) {
	fmt.Println("fmtScan, Please input：")
	//fmt.Scanf("%s %d %s", &name, &age, &class)
	//fmt.Println(name, age, class)

	fmt.Scanln(name, age, class) // split with end (\n, 0, space...), so age = 60 -> output age = 6
	fmt.Println(*name, *age, *class)
}

func useBufio() {
	fmt.Println("useBufio, Please input：")
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("%s\n", s)
}
