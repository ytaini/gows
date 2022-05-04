package gobasic

import (
	"bytes"
	"fmt"
)

func Eg91() {
	path := []byte("AAA/BBBBBBBBBB")
	sep := bytes.IndexByte(path, '/')
	dir1 := path[:sep]
	dir2 := path[sep+1:]

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{
		dir1,
		dir2,
	}, []byte{
		'/',
	})

	fmt.Printf("%s\n", dir1)
	fmt.Printf("%s\n", dir2)
	fmt.Printf("%s\n", path)
	fmt.Printf("%s\n", bytes.Join([][]byte{
		[]byte("aaa"),
		[]byte("bbbb"),
		[]byte("ccccc"),
	}, []byte("//")))

	//s[l:h:m]
	//len(s)= h-l
	//cap(s)= m-l
	a := path[:4:4]           //a还是引用path数组.但是此时a的cap为4
	a = append(a, "hello"...) //当再向a中添加元素时,append会创建新的底层数组,a将引用新的数组,这时就不会改变path数组了.
	fmt.Println(a, &a[0], &a)
	fmt.Println(path, &path[0], &path)
	fmt.Println(cap(a))
	fmt.Println(cap(dir1))
}
