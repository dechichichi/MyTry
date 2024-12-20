package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServiceURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to register service: %s", res.Status)
	}
	return nil
}

func ShutdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url,
		bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to shutdown service: %s", res.Status)
	}
	return nil
}

type providers struct {
	services map[Servicename][]string
	mutex    *sync.RWMutex
}

func (p *providers)Update(pat patch){
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _,patchEntry:=range pat.Added{
		if _,ok:=p.services[patchEntry.Name];!ok{
			p.services[patchEntry.Name]=make([]string,0)
		}
		p.services[patchEntry.Name]=append(p.services[patchEntry.Name],patchEntry.URL)
	}
	for _,patchEntry:=range pat.Removed{
		if procide
	}
}

var prov=providers{
	services: make(map[Servicename][]string),
	mutex:    new(sync.RWMutex),
}
