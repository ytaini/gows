/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 17:54:06
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 04:01:24
 */
package filetest

import (
	"fmt"
	"testing"

	"zwngkey.cn/designpattern/idioms/functionaloptions/file"
)

func Test(t *testing.T) {

	// 默认创建一个空文件
	err := file.New("./tmp")
	if err != nil {
		panic(err)
	}

	// 创建并写入内容,并改变文件拥有者.
	content := `functional option write file`
	err = file.New("./tmp1", file.UID(501), file.GID(20), file.Contents(content))
	if err != nil {
		panic(err)
	}
}

func Test1(t *testing.T) {
	// 下面的代码展示了如何使用这个构造函数来创建一个文件：
	err := file.New("/tmp/test.txt",
		file.UID(1000),
		file.GID(1000),
		file.Contents("hello world"),
		file.Permissions(0644),
	)
	if err != nil {
		fmt.Println(err)
	}
}
