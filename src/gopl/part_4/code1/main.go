package main

import (
	"fmt"
	"time"
)

type Emp struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
	Salary   int
	MID      int
}

var dilbert Emp

func main() {
	dilbert.Salary -= 5000

	// position := dilbert.Position
	// p := &position
	// fmt.Printf("position: %T\n", p) //*string

	position := &(dilbert.Position)
	position1 := (&dilbert).Position
	fmt.Printf("position: %T\n", position)   //*string
	fmt.Printf("position1: %T\n", position1) //string

	*position = "senios " + *position

	var eOT *Emp = &dilbert
	fmt.Printf("eOT.Position: %T\n", eOT.Position)

}
