package registryservice

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	Port = "3000"
	Host = "localhost"
	Addr = Host + ":" + Port
	Path = "/registry"
	URL  = "http://" + Addr + Path
)

type Server struct{}

func (*Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var ri RegistryInfo
		if err := dec.Decode(&ri); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := reg.add(&ri); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Adding Server : [%v] with URL: %v\n", ri.ServiceName, ri.ServiceURL)
		for _, regInfo := range reg.registryInfos {
			fmt.Printf("%+v\n", regInfo)
		}
	case http.MethodDelete:
		idBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		serviceId := string(idBytes)

		if err := reg.remove(ServiceID(serviceId)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Removing Server : serviceId [%s]\n", serviceId)
		for _, regInfo := range reg.registryInfos {
			fmt.Printf("%+v\n", regInfo)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
