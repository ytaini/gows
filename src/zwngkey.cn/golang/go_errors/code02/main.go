/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 17:55:42
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 19:29:34
 * @Description:
 */
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	return config, fmt.Errorf("could not read config: %w", err)
}
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open failed: %w", err)
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}
	return buf, nil
}

// could not read config: open failed: open /Users/imzw/.gitconig: no such file or directory
// could not read config: open failed: open /Users/imzw/.gitconig: no such file or directory
