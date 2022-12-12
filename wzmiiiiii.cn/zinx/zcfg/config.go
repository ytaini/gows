package zcfg

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

/*
	存储一切有关zinx框架的全局参数,供其他模块使用,
	一些参数是可以通过zinx.json由用户进行配置.
*/

type GlobalConfig struct {
	// 当前服务器主机监听的IP
	Host string `json:"host" yaml:"host"`
	// 当前服务器主机监听的端口号
	TcpPort int `json:"port" yaml:"port"`
	// 当前服务器的名称
	Name string `json:"name" yaml:"name"`
	// 当前zinx版本号
	Version string `json:"version" yaml:"version"`
	// 当前服务器主机允许的最大连接数
	MaxConn int `json:"max_conn" yaml:"max_conn"`
	// 当前zinx框架数据包的最大值
	MaxPackageSize uint32 `json:"max_package_size" yaml:"max_package_size"`
	// 日志位置
	LogPath string `json:"log_path" yaml:"log_path"`
	// worker池的goroutine数量
	WorkerPoolSize uint32 `json:"worker_pool_size" yaml:"worker_pool_size"`
	// 每个worker对应的消息队列的任务数量的最大值
	MaxWorkerTaskLen uint32
}

// Config 定义一个全局的对外Config
var Config *GlobalConfig

// 初始化GlobalObject对象
func init() {
	// 默认值
	Config = &GlobalConfig{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		Host:             "0.0.0.0",
		TcpPort:          8999,
		MaxConn:          100,
		MaxPackageSize:   4096,
		LogPath:          "log/zinx.log",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	// 尝试从conf/zinx.json去加载一些用户自定义的参数.
	if err := Config.reloadByYaml(); err != nil {
		if err = Config.reloadByJson(); err != nil {
			log.Println("lack zcfg file...")
		}
	}
}

// reloadByJson 从zinx.json去加载用户自定义的参数
func (c *GlobalConfig) reloadByJson() error {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return nil
}

// reloadByYaml 从zinx.json去加载用户自定义的参数
func (c *GlobalConfig) reloadByYaml() error {
	data, err := os.ReadFile("conf/zinx.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return nil
}
