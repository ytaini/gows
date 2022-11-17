/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-15 00:03:07
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-16 22:36:33
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
	"strconv"
	"strings"
)

// 缓存大小
const bufferSize = 5000000

//统计词频，压缩的基石
//使用map记录文件中字符对应ASCII码出现的次数
func GetFrequencyMap(path string) map[byte]int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	imap := make(map[byte]int)
	reader := bufio.NewReader(file)       //
	buffer := make([]byte, bufferSize)    // 缓存数组
	readedCount, _ := reader.Read(buffer) // 读入文件数据
	for i := 0; i < readedCount; i++ {    // 统计词频
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

// 实现了heap接口的结构体既是可排序数组，也可以当作优先队列来使用
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
	// 根据imap构造节点
	for k, v := range imap {
		plist = append(plist, &TreeNode{data: k, times: v})
	}
	// 按频率排序
	// sort.Sort(plist)
	return plist
}

// 使用优先队列构造哈夫曼树
// 我们需要保证构建的哈夫曼树具有最小权值路径和，
// 这样才能使得压缩的效率最大化，从代码可以清晰的看到，实现了heap接口后的哈夫曼树构造非常简单。
// 因为treeHeap是优先队列的缘故，所有每次Pop出来的一定是times最小的结点,这样就保证下次合成的节点权值之和是最小的.
func CreateHuffmanTree(plist TreeHeap) *TreeNode {
	heap.Init(&plist) // 生成最小堆
	for plist.Len() > 1 {
		// head.Pop(): removes and returns the minimum element (according to Less) from the heap.
		// head.Pop(): 通过小顶堆实现.
		t1 := heap.Pop(&plist).(*TreeNode)
		t2 := heap.Pop(&plist).(*TreeNode)
		newNode := &TreeNode{times: t1.times + t2.times}
		if t1.times > t2.times {
			newNode.right, newNode.left = t1, t2
		} else {
			newNode.right, newNode.left = t2, t1
		}
		heap.Push(&plist, newNode) //head.Push()
	}
	return plist[0] // plist[0]即为构造完成哈夫曼树的根节点.
}

// 使用回溯法得到经过哈夫曼树编码以后的哈夫曼编码
// encodeMap中key为字符的ASCII码值，用byte表示，
// value为哈夫曼树从根节点到叶子节点的左右路径编码值，用string表示，
// 其中tmp记录递归过程中所走过的路径，往左走用'0'表示，往右走用'1'表示，最后再转换成string就得到了对应ASCII码的编码值。
// node: 哈夫曼树root结点
func CreateEncodingTable(node *TreeNode) map[byte]string {
	encodeTable := make(map[byte]string)
	tmp := make([]byte, 0)
	var depth func(*TreeNode)
	depth = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.left == nil && tn.right == nil {
			encodeTable[tn.data] = string(tmp)
		}
		tmp = append(tmp, '0')
		depth(tn.left)
		tmp[len(tmp)-1] = '1' //回溯
		depth(tn.right)
		tmp = tmp[:len(tmp)-1] //回溯
	}
	depth(node)
	return encodeTable
}

func Encoding(inPath, outpath string, encodeMap map[byte]string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	buf := make([]byte, bufferSize)
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
	buf := make([]byte, bufferSize)
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

// https://www.jianshu.com/p/bd7fa3ca2b1f
// 另一种将字符串"1010101010101010.."按每八位存储为一个byte的方式
func Encoding1(inPath string, outPath string, encodeMap map[int]string) {
	/* 一次性读入，存到string或者buffer.string中 */
	/*1.先尝试一次性读入*/
	inFile, err := os.Open(inPath)
	if err != nil {
		return
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	fileContent := make([]byte, bufferSize)
	count, _ := reader.Read(fileContent)
	var buff bytes.Buffer
	//string编码
	for i := 0; i < count; i++ {
		v := fileContent[i]
		if code, ok := encodeMap[int(v)]; len(code) != 0 && ok {
			buff.WriteString(code)
		}
	}
	/*
		将字符串"1010101010101010"按每八位存储为一个byte:
			每个byte的初始状态都是00000000（八位）
			假设要存入的字符串为"0111011110101"（13位）,那么我们需要两个byte来记录字符串的信息，
				从前往后遍历字符串，用pos记录当前遍历到的字符串下标，读到'1'时，对byte的对应比特位进行或操作,
				这样进行八次读取，第一个byte就变成了"011101111"的倒序排列byte:111101110，因为后续我们读入的时候也是倒序读入，所以这样做是可行的，
				还剩最后五位"10101"用一个新的byte来记录，同理byte2 : 00010101
				byte2的前三位是空余位，读入的时候我们并不考虑，所以用left来记录最后一个byte要读入的bit位数即可。left = 13 % 8 = 5 。
	*/
	res := make([]byte, 0)
	var tmp byte = 0
	//编码 : 将字符串"1010101010101010"按每八位存储为一个byte.
	for idx, bit := range buff.Bytes() {
		//每八个位使用一个byte读取，结果存入res数组即可
		pos := idx % 8
		if pos == 0 && idx > 0 {
			// fmt.Printf("%b\n", buf)
			res = append(res, tmp)
			tmp = 0 //
		}
		if bit == '1' {
			tmp |= 1 << pos
		}
	}
	//TODO 这个left是剩余待处理的位数
	left := buff.Len() % 8
	res = append(res, tmp)

	// 将编码数组写入文件 , TODO 先将码表和left数写入文件，解码时在开头读取
	writeTable(outPath, encodeMap, left)
	outFile, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	wcount, err := outFile.Write(res)
	if err != nil {
		fmt.Println(wcount)
		return
	}
}

func writeTable(path string, codeMap map[int]string, left int) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	// 第一行，写入文件头的长度
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(len(codeMap)+1) + "\n")
	for k, v := range codeMap {
		buff.WriteString(strconv.Itoa(k) + ":" + v + "\n")
	}
	buff.WriteString(strconv.Itoa(left) + "\n")
	file.WriteString(buff.String())
	file.Close()
}

func Depress(inPath, depressPath string) {
	// originPath 原文件(或者可以传入码表)， inPath 读入被压缩的文件 , depressPath 还原后的输出路径
	encodeMap := make(map[int]string)
	decodeMap := make(map[string]int)
	//2.读入压缩文件
	compressFile, _ := os.Open(inPath)
	// br 读取文件头 ,返回偏移量
	br := bufio.NewReader(compressFile)
	left, offset := readTable(*br, encodeMap)
	for idx, v := range encodeMap {
		decodeMap[v] = idx
	}
	// 解码string暂存区
	var buff bytes.Buffer
	// 编码bytes暂存区
	codeBuff := make([]byte, bufferSize)
	codeLen, err := compressFile.ReadAt(codeBuff, int64(offset))
	if err != nil {
		fmt.Println(err)
		return
	}
	//遍历解码 , 读取比特
	for i := 0; i < codeLen; i++ {
		//对每个byte单独进行位运算转string
		perByte := codeBuff[i]
		for j := 0; j < 8; j++ {
			//与运算
			buff.WriteString(strconv.Itoa(int((perByte >> j) & 1)))
		}
	}
	// 对照码表，解码string , 对8取余目的是解决正好读满8个bit的情况发生
	contentStr := buff.String()[:buff.Len()-(8-left)%8]
	bytes := make([]byte, 0)
	//用切片读contenStr即可
	for star, end := 0, 1; end <= len(contentStr); {
		charValue, ok := decodeMap[contentStr[star:end]]
		if ok {
			bytes = append(bytes, byte(charValue))
			star = end
		}
		end++
	}

	depressFile, _ := os.Create(depressPath)
	depressFile.Write(bytes)
	depressFile.Close()
}

func readTable(br bufio.Reader, encodeMap map[int]string) (int, int) {
	lineStr, _, _ := br.ReadLine()
	lines, _ := strconv.Atoi(string(lineStr))
	for i := 0; i < lines-1; i++ {
		lineContent, _, _ := br.ReadLine()
		kvArr := strings.Split(string(lineContent), ":")
		k, v := kvArr[0], kvArr[1]
		kNum, _ := strconv.Atoi(k)
		encodeMap[kNum] = v
	}
	leftStr, _, _ := br.ReadLine()
	left, _ := strconv.Atoi(string(leftStr))
	return left, br.Size() - br.Buffered()
}
