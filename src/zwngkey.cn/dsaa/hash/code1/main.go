/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-01 20:48:46
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-01 22:48:01
 */
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"zwngkey.cn/dsaa/hash/code1/case01"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	linkLikedCount := 16
	hashTable := case01.NewEmpHashTable(linkLikedCount)
	for i := 0; i < 2*linkLikedCount; i++ {
		randInt := rand.Intn(300)
		randStr := strconv.Itoa(randInt)
		emp := case01.NewEmployee(randInt, "张三"+randStr)
		hashTable.Add(emp)
	}
	hashTable.Traversal()
	fmt.Println()
	fmt.Println(hashTable.FindEmpByID(143))
	fmt.Println(hashTable.DeleteEmpByID(143))
	fmt.Println()
	hashTable.Traversal()
}
