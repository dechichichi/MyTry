package main

import (
	"context"
	"fenbus/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort
	go func() {
		log.Panicln(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%Registry service started. Press any key to stop.\n")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
}
