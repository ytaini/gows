https://www.golangtc.com/t/5e1c5f34b17a82532765a0e5

# uintptr 和 unsafe普及
    uintptr表示的指针地址的值，可以用来进行数值计算
    ```
    var i int8 = 10
    //输出i变量本身的地址的十进制与十六进制
	fmt.Printf("%d,%v\n", &i, &i) //1374390714370,0x14000120002

    //p变量中存的是i变量的地址.
    p := unsafe.Pointer(&i)
	fmt.Println(p)  //0x14000120002

    //u变量中存的是 i变量中存的值.
	u := uintptr(i)
	fmt.Println(u) //10
    
	k := &i
    //u变量中存的是 k变量中存的值.而k变量中存的是一个地址值.
	u = uintptr(unsafe.Pointer(k))
	fmt.Println(u) //1374390714370
    ```
    即uintptr类型的变量 可以存 任何指针中存的值. 如果它存的是指针地址,它存的是十进制的,而不是十六进制的.
    而且它可以进行数值计算.


# unsafe 包
```
//ArbitraryType的类型也是int，但它被赋予特殊的含义，代表一个Go的任意表达式类型
type ArbitraryType int

//Pointer是一个int指针类型，在Go种，它是所有指针类型的父类型，也就是说所有的指针类型都可以转化为Pointer, uintptr和Pointer可以相互转化
type Pointer *ArbitraryType

//返回指针变量在内存中占用的字节数(记住，不是变量对应的值占用的字节数)
func Sizeof(x ArbitraryType) uintptr

/*Offsetof返回变量指定属性的偏移量，这个函数虽然接收的是任何类型的变量，但是有一个前提，就是变量要是一个struct类型，且还不能直接将这个struct类型的变量当作参数，只能将这个struct类型变量的属性当作参数*/
func Offsetof(x ArbitraryType) uintptr

//返回变量对齐字节数量
func Alignof(x ArbitraryType) uintptr
```

# 什么是内存对齐?为什么要内存对齐?

