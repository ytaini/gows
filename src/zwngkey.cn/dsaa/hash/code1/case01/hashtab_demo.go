/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-01 20:49:37
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-01 23:03:55
 */
package case01

import (
	"fmt"
)

type Employee struct {
	id   int
	name string
	next *Employee
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func (e *Employee) String() string {
	return fmt.Sprintf("=> id=%d name=%s\t", e.id, e.name)
}

type EmpLinkedList struct {
	head *Employee
}

func (ell *EmpLinkedList) add(emp *Employee) {
	if ell.findEmpByID(emp.id) != nil {
		return
	}
	if ell.head == nil {
		ell.head = emp
	} else {
		t := ell.head
		if t.id > emp.id {
			ell.head = emp
			emp.next = t
		}
		for t.next != nil && t.next.id < emp.id {
			t = t.next
		}
		t.next, emp.next = emp, t.next
	}
}
func (ell *EmpLinkedList) traversal() {
	if ell.head == nil {
		fmt.Print("当前链表为空")
		return
	}
	t := ell.head
	for t != nil {
		fmt.Print(t)
		t = t.next
	}
}
func (ell *EmpLinkedList) findEmpByID(id int) *Employee {
	t := ell.head
	for t != nil {
		if t.id == id {
			return t
		}
		t = t.next
	}
	return nil
}
func (ell *EmpLinkedList) deleteEmpByID(id int) bool {
	t := ell.head
	if t == nil {
		return false
	}
	if t.id == id {
		ell.head = t.next
		return true
	}
	for t.next != nil {
		if t.next.id == id {
			t.next = t.next.next
			return true
		}
		t = t.next
	}
	return false
}

type EmpHashTable struct {
	employees []*EmpLinkedList
	size      int
}

func NewEmpHashTable(size int) *EmpHashTable {
	emps := make([]*EmpLinkedList, size)
	for i := 0; i < size; i++ {
		emps[i] = &EmpLinkedList{}
	}
	return &EmpHashTable{
		size:      size,
		employees: emps,
	}
}

func (eht *EmpHashTable) Add(emp *Employee) {
	index := eht.hash(emp.id)
	eht.employees[index].add(emp)
}

func (eht *EmpHashTable) Traversal() {
	for i, emps := range eht.employees {
		fmt.Printf("第%d条链表:\t", i)
		emps.traversal()
		fmt.Println()
	}
}

func (eht *EmpHashTable) hash(key int) int {
	return key % eht.size
}

func (eht *EmpHashTable) FindEmpByID(id int) *Employee {
	index := eht.hash(id)
	return eht.employees[index].findEmpByID(id)
}

func (eht *EmpHashTable) DeleteEmpByID(id int) bool {
	index := eht.hash(id)
	return eht.employees[index].deleteEmpByID(id)
}
