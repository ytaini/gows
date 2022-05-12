package goforrange

import "fmt"

func Test2() {

	map1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	for k := range map1 {
		fmt.Println(k) //map在遍历时，起始遍历索引是一个随机数，所以这里的输出结果是不能确定的
	}

}

func Test1() {
	sli1 := []int{1, 2, 3}
	sli2 := make([]*int, len(sli1))

	for i, v := range sli1 {
		sli2[i] = &v
	}
	for _, v := range sli2 {
		fmt.Println(v) //输出 3 3 3
	}
	/*
		因为for range在遍历值类型时，其中的v变量是一个值的拷贝.
		当使用&获取v变量的指针时，v变量在for range中只会声明一次，v的地址一直不变.之后循环中只是改变v中存储的值.
		在给sli[i]赋值的时,都是v变量的指针，而&v最终会指向sli1最后一个元素的值拷贝。
	*/
}
