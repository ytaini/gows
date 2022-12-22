package registryservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// RegistryServer 注册服务
func RegistryServer(r RegistryInfo) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(&r); err != nil {
		return err
	}
	res, err := http.Post(URL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code %d", res.StatusCode)
	}
	return nil
}

// DeregisterServer 取消注册服务
func DeregisterServer(serviceId ServiceID) error {
	// 创建请求
	req, err := http.NewRequest(http.MethodDelete, URL, bytes.NewBuffer([]byte(serviceId)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service. Registry service responded with code %v", res.StatusCode)
	}
	return nil
}

type serviceUpdateHandler struct{}

func (*serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("test1")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(r.Body)
	var p patch
	if err := dec.Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prov.Update(&p)
}

func UpdateHandler(path string) error {
	serviceUpdateURL, err := url.Parse(path)
	if err != nil {
		return err
	}
	http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{})
	return nil
}

func HeartbeatHandler(path string) error {
	heartbeatURL, err := url.Parse(path)
	if err != nil {
		return err
	}
	http.HandleFunc(heartbeatURL.Path, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	return nil
}

type providers struct {
	services map[ServiceName][]string
	mu       *sync.RWMutex
}

var prov = &providers{
	services: make(map[ServiceName][]string),
	mu:       &sync.RWMutex{},
}

func (p *providers) Update(pat *patch) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, patchEntry := range pat.Added {
		if _, ok := p.services[patchEntry.Name]; !ok {
			p.services[patchEntry.Name] = make([]string, 0)
		}
		p.services[patchEntry.Name] = append(p.services[patchEntry.Name], patchEntry.URL)
	}

	for _, patchEntry := range pat.Removed {
		if providerURLs, ok := p.services[patchEntry.Name]; ok {
			for i := range providerURLs {
				if providerURLs[i] == patchEntry.URL {
					p.services[patchEntry.Name] = append(providerURLs[:i], providerURLs[i+1:]...)
				}
			}
		}
	}
	for name, urls := range prov.services {
		log.Printf("%s : %v\n", name, urls)
	}
}

func (p *providers) get(name ServiceName) (string, error) {
	services, ok := p.services[name]
	if !ok {
		return "", fmt.Errorf("no providers available for service %v", name)
	}
	return services[0], nil
}

func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}
