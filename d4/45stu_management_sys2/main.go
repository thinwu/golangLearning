package main

import (
	"fmt"
	"os"
)

// 同包中不同文件，需要编译以后才能link，直接使用go run main.go 没有办法取到 managementSys
var stuMgrSys managementSys

func main() {
	stuMgrSys = managementSys{
		make(map[int16]student, 10),
	}
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
			stuMgrSys.showStudent()
		case 2:
			stuMgrSys.addStudent()
		case 3:
			stuMgrSys.delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("")
		}
	}

}
