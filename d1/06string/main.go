package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "C:\\github\\2DGame"
	fmt.Printf("\"%s\"\n", path)
	s := "I'm OK"
	fmt.Println(s)
	//反引号原样数出
	s1 := `第一行
	第二行
	第三行
	`

	s2 := `C:\github\2DGame`
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s2))
	name := "aaa"
	world := "bbb"
	//字符串拼接
	ss := name + world
	fmt.Println(ss)
	//返回字符串变量
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	//分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	fmt.Println(strings.Contains(ss, "bb"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "a"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "a"))

	//判断字串位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
