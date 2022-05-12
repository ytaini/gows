package main

import (
	"fmt"
	"log"
	"os/exec"
)

type student struct {
	id   int
	name string
}

var id int

func main() {
	var flag = true

	var stuList = make([]student, 0, 50)

	// stuList = []student{
	// 	{id: 1, name: "zs"},
	// 	{id: 2, name: "ls"},
	// 	{id: 3, name: "ww"},
	// }

	for flag {
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
		1.查看所有学生信息
		2.新增学生信息
		3.删除学生信息
		0.退出
	`)
		fmt.Print("请输入操作:")
		var num int
		fmt.Scanln(&num)

		switch num {

		case 1:
			// clear()
			lookAllStudent(stuList)
			fmt.Println()
		case 2:
			stuList = addStudent(stuList)
		case 3:
			deleteStu(&stuList)
		case 0:
			flag = false
			fmt.Println("退出系统")
		default:
			fmt.Println("输入错误请重新输入!!")
			fmt.Println()
		}
	}
}

func deleteStu(student *[]student) {
	var id int
	fmt.Print("请输入学生id")
	fmt.Scanln(&id)

}

func addStudent(stuList []student) []student {
	var name string
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	stu := student{
		id:   id,
		name: name,
	}
	stuList = append(stuList, stu)
	fmt.Println(name, "学生添加成功!")
	fmt.Println()
	id++
	return stuList
}

func lookAllStudent(stuList []student) {
	fmt.Println()
	if len(stuList) == 0 {
		fmt.Println("暂无数据!")
		return
	}
	for i, stu := range stuList {
		fmt.Printf("学生信息%d:id:%04d,name:%s\n", i, stu.id, stu.name)
	}
	fmt.Println()
}

func clear() {
	cmd := exec.Command("clear")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("错误")
	}
	fmt.Println(string(out))
}
