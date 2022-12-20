package logservice

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"wzmiiiiii.cn/sds/registryservice"
)

func SetClientLogger(serviceURL string, clientServiceName registryservice.ServiceName) {
	log.SetPrefix(fmt.Sprintf("[%s] ", clientServiceName))
	log.SetFlags(log.LstdFlags)
	log.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl *clientLogger) Write(data []byte) (int, error) {
	res, err := http.Post(cl.url+"/log", "text/plain", bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. Service responded with code %d", res.StatusCode)
	}
	return len(data), nil
}
