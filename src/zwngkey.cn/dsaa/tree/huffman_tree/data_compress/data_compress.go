/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-15 00:03:07
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-15 08:37:35
 * @Description:
	通过哈夫曼编码对文件压缩与解压缩
		(不适合压缩图片)
*/
package datacompress

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const contentBuffer = 5000000

//统计词频，压缩的基石
//使用map记录文件中字符对应ASCII码出现的次数
func GetFrequencyMap(path string) map[byte]int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	imap := make(map[byte]int)
	// 读入文件数据，readline 记入map中，统计频次
	reader := bufio.NewReader(file)
	buffer := make([]byte, contentBuffer)
	readCount, _ := reader.Read(buffer)
	for i := 0; i < readCount; i++ {
		imap[buffer[i]]++
	}
	return imap
}

type TreeNode struct {
	data  byte
	times int
	left  *TreeNode
	right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("Node [data=%c,times=%d]", t.data, t.times)
}

func (t *TreeNode) PreOrder() {
	var pre func(t *TreeNode)
	pre = func(t *TreeNode) {
		if t != nil {
			fmt.Println(t)
			pre(t.left)
			pre(t.right)
		}
	}
	pre(t)
}

type TreeHeap []*TreeNode

func (t TreeHeap) Len() int {
	return len(t)
}
func (t TreeHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t TreeHeap) Less(i, j int) bool {
	return t[i].times <= t[j].times
}

func (t *TreeHeap) Push(node any) {
	*t = append(*t, node.(*TreeNode))
}
func (t *TreeHeap) Pop() any {
	n := len(*t)
	v := (*t)[n-1]
	*t = (*t)[:n-1]
	return v
}

func New(imap map[byte]int) TreeHeap {
	plist := make(TreeHeap, 0)
	// 遍历map ,按频率排序
	for k, v := range imap {
		plist = append(plist, &TreeNode{data: k, times: v})
	}
	sort.Sort(plist)
	return plist
}

// 使用优先队列构造哈夫曼树
// 使用递归也可以生成哈夫曼树，但是不能保证树的结构是最优的，我们需要保证构建的哈夫曼树具有最小权值路径和，
// 这样才能使得压缩的效率最大化，从代码可以清晰的看到，实现了堆接口后的哈夫曼树构造非常简单。
// 因为heap是优先队列的缘故，每次插入都会按Times（频次）升序排序保证下次合成的两两节点权值之和是所有节点中最小的，
// plist[0]即为构造完成哈夫曼树的根节点.
func CreateHuffmanTree(plist TreeHeap) *TreeNode {
	heap.Init(&plist)
	for plist.Len() > 1 {
		t1 := heap.Pop(&plist).(*TreeNode)
		t2 := heap.Pop(&plist).(*TreeNode)
		root := &TreeNode{times: t1.times + t2.times}
		if t1.times > t2.times {
			root.right, root.left = t1, t2
		} else {
			root.right, root.left = t2, t1
		}
		heap.Push(&plist, root)
	}
	return plist[0]
}

// 使用回溯法得到经过哈夫曼树编码以后的table
// encodeMap中key为字符的ASCII码值，用byte表示，
// value为哈夫曼树从根节点到叶子节点的左右路径编码值，用string表示，
// 其中tmp记录递归过程中所走过的路径，往左走用'0'表示，往右走用'1'表示，最后再转换成string就得到了对应ASCII码的编码值。
func CreateEncodingTable(node *TreeNode) map[byte]string {
	encodeMap := make(map[byte]string)
	tmp := make([]byte, 0)
	var depth func(*TreeNode)
	depth = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.left == nil && tn.right == nil {
			encodeMap[tn.data] = string(tmp)
		}
		tmp = append(tmp, '0')
		depth(tn.left)
		tmp[len(tmp)-1] = '1'
		depth(tn.right)
		tmp = tmp[:len(tmp)-1]
	}
	depth(node)
	return encodeMap
}

func Encoding(inPath, outpath string, encodeMap map[byte]string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	buf := make([]byte, contentBuffer)
	var buff bytes.Buffer
	for {
		count, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		for i := 0; i < count; i++ {
			code := encodeMap[buf[i]]
			buff.WriteString(code)
		}
	}
	// fmt.Println(buff.String())
	res := make([]byte, 0)
	codeStr := buff.String()
	// fmt.Println(codeStr)
	var i int
	for i = 0; i+8 <= len(codeStr); i += 8 {
		strByte := codeStr[i : i+8]
		v, _ := strconv.ParseInt(strByte, 2, 64)
		res = append(res, byte(v))
	}
	// fmt.Println(res)
	outfile, err := os.Create(outpath)
	if err != nil {
		return
	}
	defer outfile.Close()
	outfile.WriteString(codeStr[i:] + "\n") //用第一行来存储最后不足8位的位.
	wcount, err := outfile.Write(res)
	if err != nil {
		fmt.Println(wcount)
		return
	}
}

//
func HuffmanEncoding(inPath, outPath string) {
	imap := GetFrequencyMap(inPath)
	plist := New(imap)
	rootNode := CreateHuffmanTree(plist)
	codeTable := CreateEncodingTable(rootNode)
	Encoding(inPath, outPath, codeTable)
}

func Decoding(inPath, outpath string, encodeMap map[byte]string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	buf := make([]byte, contentBuffer)
	var buff []byte
	last, _, _ := reader.ReadLine() //读取第一行,第一行记录的是最后不足8位的位.

	for {
		count, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		for i := 0; i < count; i++ {
			v1 := buf[i]
			str := fmt.Sprintf("%08b", v1)
			buff = append(buff, []byte(str)...)
		}

	}
	buff = append(buff, last...) //将最后的几位加入到buff中
	buffstr := string(buff)

	// fmt.Println(buffstr)
	codeTable := ReverseMap(encodeMap)

	var res []byte
	k := 0
	for i := 1; i <= len(buffstr); i++ {
		v := buffstr[k:i]
		if v1, ok := codeTable[v]; ok {
			res = append(res, v1)
			k = i
		}
	}
	// fmt.Println(string(res))
	outfile, err := os.Create(outpath)
	if err != nil {
		return
	}
	defer outfile.Close()
	wcount, err := outfile.Write(res)
	if err != nil {
		fmt.Println(wcount)
		return
	}
}

// 反转map
func ReverseMap(encodeMap map[byte]string) map[string]byte {
	imap := make(map[string]byte)
	for k, v := range encodeMap {
		imap[v] = k
	}
	return imap
}
