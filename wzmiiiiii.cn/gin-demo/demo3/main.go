// @author: wzmiiiiii
// @since: 2022/12/25 23:56:28
// @desc: TODO

package main

import (
	"github.com/gin-gonic/gin"
	"wzmiiiiii.cn/gind/demo3/controller"
	"wzmiiiiii.cn/gind/demo3/tool"
)

func main() {
	cfg, err := tool.ParseConfig("config/app.yaml")
	if err != nil {
		panic(err)
	}

	gin.SetMode(cfg.AppMode)

	r := gin.Default()

	controller.RegisterRouter(r)

	if err := r.Run(cfg.Address()); err != nil {
		panic(err)
	}
}
