package gopkgtime

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	time.NewTicker与time.Tick函数的使用
*/

func Eg11() {
	ticker := time.NewTicker(5 * time.Second)
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(p int) {
			temp := rand.Intn(5)
			time.Sleep(time.Duration(temp) * time.Second)
			fmt.Println("I want to sleep", temp, "seconds")
			c <- p
		}(i)
	}

loop:
	for {
		select {
		case x := <-c:
			fmt.Println("goroutine-", x, "run done")
		case <-ticker.C:
			// case <-time.Tick(5 * time.Second):
			fmt.Println("time out")
			// os.Exit(2)
			break loop
		}
	}
}
