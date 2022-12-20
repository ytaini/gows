package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	"wzmiiiiii.cn/sds/registryservice"

	"wzmiiiiii.cn/sds/service"

	"wzmiiiiii.cn/sds/logservice"
)

const (
	host = "localhost"
	port = "4000"
)

var serviceAddr = fmt.Sprintf("http://%s:%s", host, port)

func main() {
	logservice.Init("./distributed.log")

	r := registryservice.RegistryInfo{
		ServiceID:        registryservice.ServiceID(uuid.NewString()),
		ServiceName:      registryservice.LogService,
		ServicePort:      port,
		ServiceHost:      host,
		ServiceURL:       serviceAddr,
		ServiceUpdateURL: serviceAddr + "/update",
		RequiredServices: make([]registryservice.ServiceName, 0),
	}

	ctx, err := service.Start(context.Background(), r, logservice.RegisterHandlerFunc)

	if err != nil {
		log.Fatalln(err)
	}
	<-ctx.Done()
	log.Printf("Shutting down %s...\n", registryservice.LogService)
}
