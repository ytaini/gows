package test_pointer
//import "fmt"
//
//func main() {
//	var p = f()
//	fmt.Println(p)
//}

func f() *int {
	v := 1
	return &v
}
