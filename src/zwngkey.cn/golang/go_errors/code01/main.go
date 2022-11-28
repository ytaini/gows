/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 17:55:42
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 18:30:57
 * @Description:
 */
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	erro "github.com/pkg/errors"
)

func main() {
	buf, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	strpath := filepath.Join(home, ".gitconig")
	fmt.Println(strpath)
	config, err := readFile(strpath)
	return config, erro.Wrap(err, "could not read config")
}
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, erro.Wrap(err, "open failed")
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, erro.Wrap(err, "read failed")
	}
	return buf, nil
}
