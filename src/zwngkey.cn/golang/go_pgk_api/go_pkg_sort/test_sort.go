package go_pkg_sort

import (
	"fmt"
	"sort"
)

func TestSort() {
	//排序int,float,string切片
	// sort.Ints()
	// Ints()底层调用
	// sort.IntSlice(g).Sort()
	// sort.IntSlice(s).Sort()

	// sort.Float64s()
	// sort.Strings()

	family := []struct {
		Name string
		Age  int
	}{
		{"A", 12},
		{"B", 10},
		{"C", 13},
	}
	//排序其他类型切片,使用自定义比较器排序
	//区别:SliceStable会保留相等元素的原始顺序.
	// sort.Slice()
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age > family[j].Age
	})
	fmt.Printf("family: %v\n", family)

	//排序任意数据结构
	//sort.Sort
	//sort.Stable
	//只要任意类型实现了sort.Interface接口,通过上面两个函数即可对任意类型进行排序.

}

type Person struct {
	Name string
	Age  int
}

//ByAge通过对age排序实现了sort.Interface接口
type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age > b[j].Age
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func TestSort2() {
	family := []Person{
		{"A", 12},
		{"B", 10},
		{"C", 13},
	}

	sort.Sort(ByAge(family))
	sort.Stable(ByAge(family))
}
