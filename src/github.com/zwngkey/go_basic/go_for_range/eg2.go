package goforrange

import (
	"fmt"
	"time"
)

/*
	// Go 1.4开始，可以不用带上i和e，直接for range遍历
*/

func Testeg2() {
	for i := range time.Tick(time.Second) {
		fmt.Println(i)
	}
	for range time.Tick(time.Second) {

	}
}
