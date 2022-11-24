/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:42:12
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 19:28:12
 * @Description:
 */
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
