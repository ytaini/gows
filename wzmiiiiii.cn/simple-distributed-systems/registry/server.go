package registry

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"

const ServicesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mu            *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.registrations = append(r.registrations, reg)
	return nil
}

func (r *registry) remove(url string) error {
	for i := range reg.registrations {
		if reg.registrations[i].ServiceURL == url {
			r.mu.Lock()
			reg.registrations = append(reg.registrations[:i], reg.registrations[i+1:]...)
			r.mu.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at URL %s not found ", url)
}

var reg = registry{
	registrations: make([]Registration, 0),
	mu:            new(sync.Mutex),
}

type Service struct{}

func (*Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")
	switch req.Method {
	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding : %v with URL: %v\n", r.ServiceName, r.ServiceURL)

		err = reg.add(r)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
