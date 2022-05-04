package main

import (
	"log"
	"os"
)

func init() {
	file, _ := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetPrefix("[zwngkey] ")
	log.SetOutput(file)
}

func main() {

	log.Println("普通日志")
	log.Printf("putong%s\n", "日志")
	log.Fatalln("123")
	log.Panicln("1234")
}
