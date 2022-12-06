/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 17:34:14
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 03:58:09
 */
// https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md

// Functional Options (功能选项)
// Functional options are a method of implementing clean/eloquent APIs in Go.
// Options implemented as a function set the state of that option.

// go中实现默认参数
package file

import (
	"os"
)

type Options struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type Option func(args *Options)

func UID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func GID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func Contents(c string) Option {
	return func(args *Options) {
		args.Contents = c
	}
}
func Permissions(perms os.FileMode) Option {
	return func(args *Options) {
		args.Permissions = perms
	}
}

// 实现了一个名为 New 的构造函数，它用于创建文件。
// 该函数接收两个参数：filepath 表示文件路径，setters 表示一组选项。
// 在函数内部，它首先设置了一些默认值，然后使用 setters 中的选项来覆盖这些默认值。
// 最后，它使用 os.OpenFile 函数打开文件，并将内容写入文件中。最后，它使用 f.Chown 方法来更改文件的所有者。
// 通过使用选项来设置文件的各种属性，该函数提供了一种灵活的方式来创建文件。
// 例如，你可以使用 UID 选项来指定文件的所有者，使用 Contents 选项来指定文件的内容，使用 Permissions 选项来指定文件的权限等。
func New(filepath string, setters ...Option) error {
	// 默认值
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getgid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY, //os.O_EXCL: 与 os.O_CREATE一起使用,文件不能已存在.
	}

	for _, setter := range setters {
		setter(args)
	}

	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		return err
	} else {
		defer f.Close()
	}
	if _, err := f.WriteString(args.Contents); err != nil {
		return err
	}
	return f.Chown(args.UID, args.GID)
}

// 例如，下面的代码展示了如何使用这个构造函数来创建一个文件：
// New("/tmp/test.txt",
//     UID(1000),
//     GID(1000),
//     Contents("hello world"),
//     Permissions(0644),
// )
