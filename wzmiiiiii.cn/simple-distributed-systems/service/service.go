package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"wzmiiiiii.cn/sds/registry"
)

func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {

	registerHandlersFunc()

	ctx = StartService(ctx, reg.ServiceName, host, port)

	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func StartService(ctx context.Context, serviceName registry.ServiceName,
	host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server

	srv.Addr = host + ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 实现按ENTER 键退出
	go func() {
		log.Printf("%v is running. Press enter key to stop...\n", serviceName)
		var s string
		_, _ = fmt.Scanln(&s)
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}()

	return ctx
}
