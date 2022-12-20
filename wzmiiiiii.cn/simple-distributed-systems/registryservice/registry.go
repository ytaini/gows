package registryservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	LogService      = "Log Service"
	GradeService    = "Grade Service"
	RegistryService = "Registry Service"
)

type (
	ServiceID   string
	ServiceName string
)

type RegistryInfo struct {
	ServiceID        ServiceID     `json:"server_id"`
	ServiceName      ServiceName   `json:"service_name"`
	ServicePort      string        `json:"service_port"`
	ServiceHost      string        `json:"service_host"`
	ServiceURL       string        `json:"service_url"`
	ServiceUpdateURL string        `json:"service_update_url"`
	RequiredServices []ServiceName `json:"required_services"`
}

type register struct {
	registryInfos []*RegistryInfo
	mu            *sync.RWMutex
}

var reg = register{
	registryInfos: make([]*RegistryInfo, 0),
	mu:            new(sync.RWMutex),
}

func (r *register) add(ri *RegistryInfo) error {
	r.mu.Lock()
	r.registryInfos = append(r.registryInfos, ri)
	r.mu.Unlock()

	// 请求需要依赖的服务
	if err := r.requestRequiredServices(ri); err != nil {
		return err
	}

	r.notify(&patch{
		Added: []*entry{
			{Name: ri.ServiceName, URL: ri.ServiceURL},
		},
	})
	return nil
}
func (r *register) requestRequiredServices(ri *RegistryInfo) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var p patch

	for _, registryInfo := range r.registryInfos {
		for _, requiredService := range ri.RequiredServices {
			if registryInfo.ServiceName == requiredService {
				p.Added = append(p.Added, &entry{
					Name: registryInfo.ServiceName,
					URL:  registryInfo.ServiceURL,
				})
			}
		}
	}

	if err := r.sendPatch(&p, ri.ServiceUpdateURL); err != nil {
		return err
	}
	return nil
}

func (r *register) sendPatch(p *patch, url string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	if res, err := http.Post(url, "application/json", bytes.NewBuffer(data)); err != nil {
		return err
	} else {
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("request required service failed")
		}
	}
	return nil
}

func (r *register) notify(p *patch) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, reg := range r.registryInfos {
		go func(reg *RegistryInfo) {
			for _, requiredService := range reg.RequiredServices {
				tmpP := &patch{[]*entry{}, []*entry{}}
				sendUpdate := false
				for _, added := range p.Added {
					if added.Name == requiredService {
						tmpP.Added = append(tmpP.Added, added)
						sendUpdate = true
					}
				}
				for _, removed := range p.Removed {
					if removed.Name == requiredService {
						tmpP.Removed = append(tmpP.Removed, removed)
						sendUpdate = true
					}
				}
				if sendUpdate {
					if err := r.sendPatch(tmpP, reg.ServiceUpdateURL); err != nil {
						log.Println(err)
						return
					}
				}
			}
		}(reg)
	}
}

func (r *register) remove(serviceId ServiceID) error {
	for i, registryInfo := range r.registryInfos {
		if registryInfo.ServiceID == serviceId {
			r.notify(&patch{
				Removed: []*entry{
					{Name: registryInfo.ServiceName, URL: registryInfo.ServiceURL},
				},
			})
			r.mu.Lock()
			r.registryInfos = append(r.registryInfos[:i], r.registryInfos[i+1:]...)
			r.mu.Unlock()
			return nil
		}
	}
	return fmt.Errorf("ServiceId: [%s] NOT FOUNT !!! ", serviceId)
}

type entry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []*entry
	Removed []*entry
}
