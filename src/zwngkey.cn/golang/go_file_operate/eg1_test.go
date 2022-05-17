/*
 * @Author: zwngkey
 * @Date: 2022-05-17 17:37:23
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-17 21:57:32
 * @Description:
	go 文件操作
		https://colobu.com/2016/10/12/go-file-operations/#%E4%BB%8B%E7%BB%8D
*/
package gofileopetate

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test1(t *testing.T) {
	// 创建一个硬链接。
	// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
	// 删除和重命名不会影响另一个。
	err := os.Link("./file/1.txt", "./file/2.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Create a symlink
	err = os.Symlink("./file/1.txt", "./file/1_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
	// Symlink在Windows中不工作。
	fileInfo, err := os.Lstat("./file/1_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	//改变软链接的拥有者不会影响原始文件。
	err = os.Lchown("./file/1_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
func Test2(t *testing.T) {
	// 打开原始文件
	originalFile, err := os.Open("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create("./file/test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

func Test3(t *testing.T) {
	file, _ := os.Open("./file/test.txt")
	defer file.Close()

	currentPosition, _ := file.Seek(0, 1)
	fmt.Println("Current position:", currentPosition)

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5
	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	var whence int = 0

	newPosition, err := file.Seek(offset, whence)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Just moved to 5:", newPosition)

	// 从当前位置回退两个字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPosition)

	// 使用下面的技巧得到当前的位置
	currentPosition, _ = file.Seek(0, 1)

	fmt.Println("Current position:", currentPosition)
	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition)

}
func Test4(t *testing.T) {
	// 打开文件，只写
	file, err := os.OpenFile("./file/test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中已用的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 写内存buffer到硬盘
	bufferedWriter.Flush()

	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)

	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

}
func Test5(t *testing.T) {
	file, _ := os.Open("./file/test.txt")
	defer file.Close()

	// file.Read()可以读取一个小文件到大的byte slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSlice := make([]byte, 20)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
func Test6(t *testing.T) {
	file, _ := os.Open("./file/test.txt")
	defer file.Close()

	byteSlice := make([]byte, 512)
	minBytes := 20
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

}
func Test7(t *testing.T) {
	file, _ := os.Open("./file/test.txt")
	defer file.Close()

	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

}
func Test8(t *testing.T) {
	file, _ := os.Open("./file/test.txt")
	defer file.Close()

	bufferedReader := bufio.NewReader(file)

	byteSlice := make([]byte, 5)
	// 得到位置后的5个字节，当前指针不变
	byteSlice, err := bufferedReader.Peek(len(byteSlice))
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %q\n", dataBytes)

	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %q\n", dataString)

	n, err := bufferedReader.WriteTo(os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
