package gobasic

// import "fmt"

/*
   nil用于表示interface、函数、maps、slices,指针和channels的“零值”。
   如果你不指定变量的类型，编译器将无法编译你的代码，因为它猜不出具体的类型。
*/
func Testeg31() {
	//以下三种变量声明都错误
	//var x  = nil
	//var x int = nil
	//x:= nil

	//nil也不可以分配string类型
	/*
	   var x string = nil
	   _ = x
	*/
}
