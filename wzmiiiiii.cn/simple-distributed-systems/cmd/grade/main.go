package main

import (
	"context"
	"fmt"
	"log"

	"wzmiiiiii.cn/sds/logservice"

	"github.com/google/uuid"
	"wzmiiiiii.cn/sds/registryservice"

	"wzmiiiiii.cn/sds/grades"

	"wzmiiiiii.cn/sds/service"
)

const (
	host = "localhost"
	port = "6000"
)

var serviceAddr = fmt.Sprintf("http://%s:%s", host, port)

func main() {
	ri := registryservice.RegistryInfo{
		ServiceID:        registryservice.ServiceID(uuid.NewString()),
		ServicePort:      port,
		ServiceHost:      host,
		ServiceName:      registryservice.GradeService,
		ServiceURL:       serviceAddr,
		ServiceUpdateURL: serviceAddr + "/update",
		HeartbeatURL:     serviceAddr + "/heartbeat",
		RequiredServices: []registryservice.ServiceName{
			registryservice.LogService,
		},
	}

	ctx, err := service.Start(context.Background(), ri, grades.RegisterHandlers)
	if err != nil {
		log.Fatalln(err)
	}

	if logProvider, err := registryservice.GetProvider(registryservice.LogService); err == nil {
		log.Printf("Logging service found at: %s\n", logProvider)
		logservice.SetClientLogger(logProvider, ri.ServiceName)
	}
	<-ctx.Done()
	log.Printf("Shutting down %s...\n", ri.ServiceName)
}