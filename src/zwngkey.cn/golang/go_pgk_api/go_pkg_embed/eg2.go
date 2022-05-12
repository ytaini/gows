package gopkgembed

import (
	"embed"
	"log"
)

//这是一个指令,而不是普通的注释
//go:embed test/*
var f embed.FS

func Testeg2() {
	data, err := f.ReadFile("test/1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	print(string(data))
	// dirs, _ := f.ReadDir("test")
	// for _, dir := range dirs {
	// 	fmt.Printf("dir.Name(): %v\n", dir.Name())
	// }
}
