package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"wzmiiiiii.cn/sds/registryservice"
)

func main() {
	// 开启心跳检查
	registryservice.SetupRegistryService()

	http.Handle(registryservice.Path, &registryservice.Server{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := http.Server{
		Addr: registryservice.Addr,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		log.Printf("%v is running. Server Address at %s\n",
			registryservice.RegistryService, registryservice.URL)
		log.Println("Press enter key to stop...")
		var s string
		if _, err := fmt.Scanln(&s); err != nil {
			log.Println(err)
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}()

	<-ctx.Done()
	log.Printf("Shutting down %s...\n", registryservice.RegistryService)
}
