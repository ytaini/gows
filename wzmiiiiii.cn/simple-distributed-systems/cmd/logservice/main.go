package main

import (
	"context"
	"fmt"
	stdlog "log"

	"wzmiiiiii.cn/sds/registry"

	"wzmiiiiii.cn/sds/log"
	"wzmiiiiii.cn/sds/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()
	stdlog.Println("Shutting down Log Service!!!")
}
