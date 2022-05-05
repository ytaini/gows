/*
 * @Author: zwngkey
 * @Date: 2022-05-02 16:09:26
 * @LastEditTime: 2022-05-02 17:34:55
 * @Description: flag 包的使用
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var usage = `Usage: %s [options]
Options are:
    -n requests     Number of requests to perform
    -c concurrency  Number of multiple requests to make at a time
    -s timeout      Seconds to max. wait for each response
    -m method       Method name
 `

var (
	r, c, timeout int
	m, u          string
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
	flag.IntVar(&r, "n", 1000, "")
	flag.IntVar(&c, "c", 100, "")
	flag.IntVar(&timeout, "s", 10, "")
	flag.StringVar(&m, "m", "GET", "")
	flag.Parse()

	if flag.NArg() != 1 {
		exit("Invalid url")
	}
	m = strings.ToUpper(m)
	u = flag.Args()[0]

	if m != "GET" {
		exit("Invalid url")
	}
	log.Println(r, c, timeout, m, u)

	log.Println("app done")
}

func exit(msg string) {
	flag.Usage()
	fmt.Fprintln(os.Stderr, "\n[Error]: "+msg)
	os.Exit(1)
}
