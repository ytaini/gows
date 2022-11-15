/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-15 01:08:56
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-15 08:43:25
 * @Description:
 */
package datacompress

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test2(t *testing.T) {
	inPath := "source/data3"
	outPath := "zip/data3-zip"
	HuffmanEncoding(inPath, outPath)
}

func Test1(t *testing.T) {
	// inPath := "source/data3"
	// outPath := "zip/data3-zip"
	// outPath2 := "unzip/data3-unzip"
	inPath := "source/1.jpg"
	outPath := "zip/1-zip.zip"
	outPath2 := "unzip/1-unzip.jpg"
	imap := GetFrequencyMap(inPath)
	// for k, v := range imap {
	// 	fmt.Printf("%c-->%d\n", k, v)
	// }
	nodeList := New(imap)
	// for _, v := range nodeList {
	// 	fmt.Println(v)
	// }
	root := CreateHuffmanTree(nodeList)
	// root.PreOrder()
	codeTable := CreateEncodingTable(root)
	// for k, v := range codeTable {
	// 	fmt.Printf("%v-->%s\n", k, v)
	// }
	Encoding(inPath, outPath, codeTable)
	Decoding(outPath, outPath2, codeTable)
}

func Test(t *testing.T) {
	inFile, err := os.Open("source/data1")
	if err != nil {
		return
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	buf := make([]byte, contentBuffer)
	v, _, _ := reader.ReadLine()
	fmt.Println(v)
	i, _ := reader.Read(buf)
	fmt.Println(buf[:i])
}
