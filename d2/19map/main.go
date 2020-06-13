package main

import "fmt"

func main() {
	//var m1 map[string]int // nil 没有在内存中开辟空间
	var m1 = make(map[string]int, 10) // 要估算好map的容量，避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["Daniel"] = 37
	fmt.Println(m1)

	// map 键值对初始化， 不带容量估算
	// m1 ：=map[string]int{
	// 	"stu1":100,
	// 	"stu2":200,
	// 	"stu3":300,
	// }

	// ok, bool. v value， 只用一个变量接收的时候，只会返回一个值，没有找到，就用对应类型的零值。
	// 用两个变量接收，第二个规则上用ok，是一个bool类型。来判读是否取到了。

	v, ok := m1["Cyrus"]

	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)
	}
	fmt.Println(m1["Cyrus"]) // 拿到对应类型0值， 不推荐直接使用

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}
	for _, v := range m1 {
		fmt.Println(v)
	}
	//del
	delete(m1, "Daniel")
	fmt.Println(m1)
	delete(m1, "Daniel") // 删除不存在的不报错
	fmt.Println(m1)
	// 查看 delete 函数
	// go doc buildin.delete
}
