/*
  - @Author: zwngkey
  - @Date: 2021-12-11 16:05:03
  - @LastEditTime: 2022-11-21 11:10:45
  - @Description:
    类似实现cat命令
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"zwngkey.cn/gobasic/day02/05cat_cmd_lianxi/cat"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Println("err: please input param!!")
		return
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat.CatCmd(bufio.NewReader(f))
	}
}
