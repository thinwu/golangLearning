package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int16
	name string
}

func newStudent(id int16, name string) *student {
	return &student{
		id,
		name,
	}
}

var (
	allStudent = make(map[int16]*student, 10) // 声明并初始化
)

func showAllStudents() {
	for _, v := range allStudent {
		fmt.Printf("学号：%v， 姓名：%v\n", v.id, v.name)
	}
}
func addStudent() {
	var (
		id   int16
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	allStudent[id] = newStudent(id, name)
}
func deleteStudent() {
	fmt.Println("请输入要删除的学号")
	var id int16
	fmt.Scanln(&id)
	delete(allStudent, id)

}
func main() {
	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统")
		// 2. 等待用户选择 , 反引号原样输出
		fmt.Println(`
			1. 查看所有学生
			2. 新增学生
			3. 删除学生
			4. 退出
			`)
		fmt.Println("请输入你要做什么")
		var input int
		fmt.Scanln(&input)
		fmt.Printf("你选择了 %d\n", input)
		// 3. 执行对应的函数
		switch input {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("")
		}
	}

}
