package main

import "fmt"

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

type managementSys struct {
	students map[int16]student
}

func (m managementSys) addStudent() {
	var (
		id   int16
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	m.students[s.id] = *s
}
func (m managementSys) showStudent() {
	for _, v := range m.students {
		fmt.Printf("学号：%v， 姓名：%v\n", v.id, v.name)
	}
}
func (m managementSys) delStudent() {
	fmt.Println("请输入要删除的学号")
	var id int16
	fmt.Scanln(&id)
	delete(m.students, id)
}
