/*
 * @Author: zwngkey
 * @Date: 2022-05-14 06:48:04
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-14 06:59:46
 * @Description:
	解析配置文件
*/
package main

type MysqlConfig struct {
	Addr     string `ini:"addr"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Database int    `ini:"database"`
}

func loadIni(fileName string, data any) {

}

func main() {

}
