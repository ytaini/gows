package main

import (
	"log"

	"wzmiiiiii.cn/sds/logservice"
)

func main() {
	logservice.SetClientLogger("http://localhost:4000", "Grade Service")
	log.Println("12313")
	log.Println("12313")
	log.Println("12313")
	log.Println("12313")
	log.Println("12313")
}
