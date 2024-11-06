package main

import (
	"context"
	"fenbus/log"
	"fenbus/registry"
	"fenbus/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./fenbus.log")
	host, port := "localhost", "8080"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		Servicename: "Log Service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Log Service stopped")
}
