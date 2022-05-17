/*
 * @Author: zwngkey
 * @Date: 2022-05-17 21:59:12
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 00:49:42
 * @Description:
 	go 文件操作
		https://colobu.com/2016/10/12/go-file-operations/#%E4%BB%8B%E7%BB%8D
*/
package gofileopetate

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test9(t *testing.T) {
	file, err := os.Open("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)
	// scanner.Split(bufio.ScanBytes)

	// 也可以定制一个SplitFunc类型的分隔函数
	// mySplitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 	return 0, nil, nil
	// }
	// scanner.Split(mySplitFunc)

	for scanner.Scan() {
		// 得到数据，Bytes() 或者 Text()
		fmt.Printf("First word found:%q\n", scanner.Text())
	}
	// 出现错误或者EOF是返回Error
	if scanner.Err() != nil {
		log.Fatal(err)
	}

}
func Zip(dst, src string) (err error) {
	// 创建准备写入的文件
	fw, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fw.Close()

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		if err := zw.Close(); err != nil {
			log.Println(err)
		}
	}()

	walkFileFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fi, err := d.Info()
		if err != nil {
			return err
		}

		// 通过文件信息，创建 zip 的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		// 替换文件信息中的文件名
		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		// 这步开始没有加，会发现解压的时候说它不是个目录
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 zip.Write 结构
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		// 将打开的文件 Copy 到 w
		n, err := io.Copy(w, fr)
		if err != nil {
			return err
		}

		// 输出压缩的内容
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)

		return nil
	}

	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
	return filepath.WalkDir(src, walkFileFunc)

}

func Test10(t *testing.T) {
	// 源档案（准备打包的文件或目录）
	var src = "./file"
	// 目标文件，打包后的文件
	var dst = "./file.zip"

	if err := Zip(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func UnZip(dst, src string) (err error) {
	// 打开压缩文件 , zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(src)
	if err != nil {
		return
	}
	defer zr.Close()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	return nil
}

func Test11(t *testing.T) {
	// 压缩包
	var src = "./file.zip"
	// 解压后保存的位置，为空表示当前目录
	var dst = "./file2"

	if err := UnZip(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func Test12(t *testing.T) {
	err := os.MkdirAll("./a/b/c", 0766)
	// err := os.Mkdir("./a", 0766)
	if err != nil {
		log.Fatalln(err)
	}
}
