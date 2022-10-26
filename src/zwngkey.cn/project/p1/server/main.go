/*
 * @Author: zwngkey
 * @Date: 2022-05-14 02:39:17
 * @LastEditors: imzw 1714894407@qq.com
 * @LastEditTime: 2022-07-16 07:35:48
 * @Description:
 */
package main

func main() {
	server := NewServer("127.0.0.1", 8080)
	server.Run()
}
