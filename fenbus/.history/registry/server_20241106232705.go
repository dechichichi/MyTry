package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":8080"
const ServiceURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.RWMutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	err := r.sendRequiredServices(reg)
	if err != nil {
		return err
	}
	return nil
}

func (r *registry) sendRequiredServices(reg Registration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var p patch
	for _, s := range r.registrations {
		for _, r := range reg.RegistryServices {
			p.Added = append(p.Added, patchEntry{
				Name: s.Servicename,
				URL:  s.ServiceURL,
			})
		}
	}
	err := r.sendPatch(p, reg.ServiceURL)
	if err != nil {
		return err
	}
	return nil

}

func (r *registry) sendPatch(p patch, url string) error {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewReader(d))
}

func (r *registry) remove(url string) error {
	for i := range reg.registrations {
		if reg.registrations[i].ServiceURL == url {
			r.mutex.Lock()
			reg.registrations = append(reg.registrations[:i], reg.registrations[:i+1]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("Service with URL %s not found", url)
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.RWMutex),
}

type RegistryService struct{}

func (r RegistryService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request service")
	switch req.Method {
	case http.MethodGet:
		dec := json.NewDecoder(req.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service:%v with URL %s\n", r.Servicename, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Panicln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(payload)
		log.Printf("Removing service with URL %s\n", url)
		err = reg.remove(url)
		if err != nil {
			log.Panicln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
