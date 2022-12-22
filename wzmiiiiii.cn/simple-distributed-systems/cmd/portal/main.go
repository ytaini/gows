package main

import (
	"context"
	"fmt"
	"log"

	"wzmiiiiii.cn/sds/logservice"

	"wzmiiiiii.cn/sds/service"

	"github.com/google/uuid"
	"wzmiiiiii.cn/sds/registryservice"

	"wzmiiiiii.cn/sds/portal"
)

const (
	port = "7000"
	host = "localhost"
)

var serviceAddr = fmt.Sprintf("http://%s:%s", host, port)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		log.Fatalln(err)
	}
	ri := registryservice.RegistryInfo{
		ServiceID:        registryservice.ServiceID(uuid.NewString()),
		ServiceName:      registryservice.PortalService,
		ServicePort:      port,
		ServiceHost:      host,
		ServiceURL:       serviceAddr,
		ServiceUpdateURL: serviceAddr + "/update",
		HeartbeatURL:     serviceAddr + "/heartbeat",
		RequiredServices: []registryservice.ServiceName{
			registryservice.LogService,
			registryservice.GradeService,
		},
	}
	ctx, err := service.Start(context.Background(), ri, portal.RegisterHandlers)

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
