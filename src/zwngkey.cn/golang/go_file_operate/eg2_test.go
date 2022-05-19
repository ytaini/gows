/*
 * @Author: zwngkey
 * @Date: 2022-05-17 21:59:12
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 15:50:27
 * @Description:
	go 文件打包,拆包,压缩,解压
*/
package gofileopetate

import (
	"archive/zip"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}
		err := unZip(path, file)
		if err != nil {
			return err
		}
	}
	return nil
}
func unZip(path string, file *zip.File) error {
	// 获取到 Reader
	fr, err := file.Open()
	if err != nil {
		return err
	}
	defer fr.Close()

	// 创建要写出的文件对应的 Write
	fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer fw.Close()

	n, err := io.Copy(fw, fr)
	if err != nil {
		return err
	}

	// 将解压的结果输出
	log.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

	// 因为是在循环中，无法使用 defer ，直接放在最后
	// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
	// 可以把它单独放在一个函数中，这里是个实验，就这样了
	// fw.Close()
	// fr.Close()
	return nil
}

func Test11(t *testing.T) {
	// 压缩包
	var src = "./file.zip"
	// 解压后保存的位置，为空表示当前目录
	var dst = "./file_copy"

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

func Test13(t *testing.T) {
	f, err := os.Create("./test.zip.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	gzipW := gzip.NewWriter(f)
	defer gzipW.Close()

	f1, err := os.Open("./file.zip")
	if err != nil {
		log.Fatalln(err)
	}

	n, err := io.Copy(gzipW, f1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("成功压缩 ，共写入了 %d 个字符的数据\n", n)
}

func Test14(t *testing.T) {
	f, err := os.Open("./test.zip.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	gzipR, err := gzip.NewReader(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer gzipR.Close()

	f1, err := os.Create("./file_copy.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()

	n, err := io.Copy(f1, gzipR)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("成功解压 ，共写入了 %d 个字符的数据\n", n)
}

func Test15(t *testing.T) {
	fmt.Printf("os.TempDir(): %v\n", os.TempDir())
	// 在系统临时文件夹中创建一个临时文件夹
	tempDirPath, err := os.MkdirTemp("", "myTempDir")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(tempDirPath)

	// 在临时文件夹中创建临时文件
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Temp file created:", tempFile.Name())

	// 关闭文件
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 删除我们创建的资源
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}

}

func Test16(t *testing.T) {
	data, err := ioutil.ReadFile("./file/1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// 计算Hash
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}

func Test17(t *testing.T) {
	f, err := os.Open("./file/1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	hs := md5.New()

	_, err = io.Copy(hs, f)
	if err != nil {
		log.Fatalln(err)
	}

	// 计算hash并打印结果。
	// 传递 nil 作为参数，因为我们不通过参数传递数据，而是通过writer接口。
	sum := hs.Sum(nil)
	fmt.Printf("Md5 checksum: %x\n", sum)
}
