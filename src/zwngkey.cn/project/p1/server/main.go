/*
 * @Author: zwngkey
 * @Date: 2022-05-14 02:39:17
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-15 21:37:39
 * @Description:
 */
package main

func main() {
	server := NewServer("127.0.0.1", 8080)
	server.Run()
}
