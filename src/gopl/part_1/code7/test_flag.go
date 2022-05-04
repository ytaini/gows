package main

import (
	"flag"
	"fmt"
)

var (
	intFlag    int
	boolFlag   bool
	stringFlag string
)

func init() {
	flag.IntVar(&intFlag, "i", 0, "int flag value")
	flag.BoolVar(&boolFlag, "b", false, "bool flag value")
	flag.StringVar(&stringFlag, "s", "abc", "string flag value")
}

// func init() {
//   intflag := flag.Int("intflag", 0, "int flag value")
//   boolflag := flag.Bool("boolflag", false, "bool flag value")
//   stringflag := flag.String("stringflag", "default", "string flag value")
// }

func main() {
	flag.Parse()

	fmt.Println("int flag:", intFlag)
	fmt.Println("bool flag:", boolFlag)
	fmt.Println("string flag:", stringFlag)

}
