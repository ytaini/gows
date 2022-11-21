/*
 * @Author: zwngkey
 * @Date: 2022-05-15 08:30:56
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 20:21:52
 * @Description:
 */
package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Addr     string `ini:"addr"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

var supportConfig = []string{"mysql", "redis"}

func findSlice(s []string, tar string) bool {
	for _, v := range s {
		if v == tar {
			return true
		}
	}
	return false
}

func checkDataType(data any) error {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Pointer {
		return errors.New("data 不是指针类型")
	}

	valEle := val.Elem()

	if valEle.Kind() != reflect.Struct {
		return errors.New("data 的基类型不是结构体类型")
	}

	_, ok := valEle.Interface().(Config)

	if !ok {
		return errors.New("data 的基类型不是Config类型")
	}
	return nil
}

func loadIni(fileName string, data any) error {
	err := checkDataType(data)

	if err != nil {
		return err
	}

	buf, err := os.ReadFile(fileName)

	if err != nil {
		return errors.New("open file 失败")
	}

	var fileContent = string(buf)

	lineSlice := strings.Split(fileContent, "\n")

	var structName string

	for lineNo, line := range lineSlice {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			configType := line[1 : len(line)-1]

			if !findSlice(supportConfig, configType) {
				return fmt.Errorf("解析失败: 第%v行语法错误", lineNo+1)
			}

			v := reflect.TypeOf(data).Elem()

			for i := 0; i < v.NumField(); i++ {
				if configType == v.Field(i).Tag.Get("ini") {
					structName = v.Field(i).Name
					break
				}
			}
		} else {

			if !strings.Contains(line, "=") || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				return fmt.Errorf("第%v行语法错误", lineNo+1)
			}

			ss := strings.Split(line, "=")

			fType, _ := reflect.TypeOf(data).Elem().FieldByName(structName)

			fVal := reflect.ValueOf(data).Elem().FieldByName(structName)

			ftt := fType.Type

			for j := 0; j < ftt.NumField(); j++ {
				tagName := ftt.Field(j).Tag.Get("ini")
				if tagName == ss[0] {
					rf := fVal.Field(j)
					if rf.CanInt() {
						sa, err := strconv.Atoi(ss[1])
						if err != nil {
							return fmt.Errorf("第%v行语法错误", lineNo+1)
						}
						rf.SetInt(int64(sa))
					} else {
						rf.SetString(ss[1])
					}
				}
			}
		}
	}

	return nil
}
