package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("clear")
	// out, err := cmd.CombinedOutput()
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(out))
}
