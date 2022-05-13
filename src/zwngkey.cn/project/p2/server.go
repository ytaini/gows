/*
 * @Author: zwngkey
 * @Date: 2022-05-14 05:53:56
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 05:57:47
 * @Description:
 */
package main

import (
	"fmt"
	"sort"
	"strconv"
)

type idx = string

type Student struct {
	id   string
	name string
	age  int
}

type Server struct {
	stus     map[idx]Student
	choice   string
	stuCount int
}

func NewServer() *Server {
	return &Server{stus: make(map[idx]Student)}
}

func (s *Server) run() {
loop:
	for {
		s.print()

		fmt.Print("请选择功能[序号]:")

		fmt.Scanf("%s", &s.choice)

		if !s.checkChoice() {
			fmt.Println("非法输入,请输入对应的序号!!")
			s.pause()
			continue
		}

		switch s.choice {
		case "1":
			s.showAllStus()
			s.pause()

		case "2":
			var name string
			var age int
			fmt.Print("请输入学生姓名:")
			fmt.Scanln(&name)
			fmt.Print("请输入学生年龄:")
			fmt.Scanln(&age)
			s.addStu(name, age)
			fmt.Println("提示:学生信息添加成功")
			s.pause()
		case "3":
			var id string
			for {
				fmt.Print("请输入学生学号:")
				fmt.Scanln(&id)
				if s.removeStu(id) {
					fmt.Println("学生信息删除成功")
					break
				}
				if id == "x" {
					break
				}
				fmt.Println("没有此学号,重新输入[按x退出]")
			}
			s.pause()
		case "4":
			var id string
			for {
				fmt.Print("请输入学生学号:")
				fmt.Scanln(&id)
				stu, ok := s.findStu(id)
				if ok {
					fmt.Println("学生信息查询成功")
					fmt.Printf("学号: %v\t姓名:%-10v\t年龄:%v\n", stu.id, stu.name, stu.age)
					break
				}
				if id == "x" {
					break
				}
				fmt.Println("没有此学号,重新输入[按x退出]")
			}
			s.pause()
		case "0":
			s.exit()
			break loop
		}
	}
}

func (s *Server) findStu(id string) (Student, bool) {
	stu, ok := s.stus[id]
	return stu, ok
}

func (s *Server) removeStu(id string) bool {
	_, ok := s.stus[id]
	delete(s.stus, id)
	s.stuCount--
	return ok

}

func (s *Server) addStu(name string, age int) {
	s.stuCount++
	id := fmt.Sprintf("%04d", s.stuCount)
	stu := Student{
		id,
		name,
		age,
	}
	s.stus[id] = stu
}

func (s *Server) showAllStus() {
	if s.stuCount == 0 {
		fmt.Println("提示: 没有学生信息!")
		return
	}
	var keys []string
	for key := range s.stus {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("学号: %v\t姓名:%-10v\t年龄:%v\n", s.stus[key].id, s.stus[key].name, s.stus[key].age)
	}

}
func (s *Server) checkChoice() bool {
	v, err := strconv.Atoi(s.choice)
	if err != nil || v < 0 || v > 4 {
		return false
	}
	return true
}
func (s *Server) exit() {
	fmt.Println("exit system")
}

func (*Server) print() {
	const welcome string = `
---------go 方法版学生管理系统------------
            1.查看所有学生信息
            2.增加学生信息
            3.删除学生信息
            4.查询学生信息
            0.退出
	`
	fmt.Println(welcome)
}

func (*Server) pause() {
	fmt.Println()
	fmt.Print("按ENTER继续")
	fmt.Scanln()
}
