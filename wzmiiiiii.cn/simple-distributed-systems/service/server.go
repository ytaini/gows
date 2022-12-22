package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"wzmiiiiii.cn/sds/registryservice"
)

func Start(ctx context.Context, reg registryservice.RegistryInfo,
	registerHandlerFunc func()) (context.Context, error) {

	if err := registryservice.HeartbeatHandler(reg.HeartbeatURL); err != nil {
		return ctx, err
	}

	if err := registryservice.UpdateHandler(reg.ServiceUpdateURL); err != nil {
		return ctx, err
	}

	registerHandlerFunc()

	ctx = startService(ctx, reg)

	//服务注册
	if err := registryservice.RegistryServer(reg); err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, reg registryservice.RegistryInfo) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	addr := reg.ServiceHost + ":" + reg.ServicePort
	srv := http.Server{
		Addr: addr,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println(err)
			// 取消服务注册
			if err := registryservice.DeregisterServer(reg.ServiceID); err != nil {
				log.Println(err)
			}
			cancel()
		}
	}()

	go func() {
		log.Printf("%v is running. Server Address at %s\n", reg.ServiceName, reg.ServiceURL)
		log.Println("Press enter key to stop...")
		var s string
		if _, err := fmt.Scanln(&s); err != nil {
			log.Println(err)
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		// 取消服务注册
		if err := registryservice.DeregisterServer(reg.ServiceID); err != nil {
			log.Println(err)
		}
		cancel()
	}()
	return ctx
}
