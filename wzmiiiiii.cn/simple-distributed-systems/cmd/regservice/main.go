package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"wzmiiiiii.cn/sds/registry"
)

func main() {
	http.Handle("/services", &registry.Service{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		log.Println("Registry Service is running. Press enter key to stop...")
		var s string
		_, _ = fmt.Scanln(&s)
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}()

	<-ctx.Done()
	log.Println("Shutting down registry service...")
}
