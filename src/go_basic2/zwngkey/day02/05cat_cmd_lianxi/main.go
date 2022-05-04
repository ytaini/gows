/*
 * @Author: zwngkey
 * @Date: 2021-12-11 16:05:03
 * @LastEditTime: 2022-04-30 20:51:59
 * @Description:
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"go_basic2/zwngkey/day02/05cat_cmd_lianxi/cat"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalln("err: please input param!!")
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
