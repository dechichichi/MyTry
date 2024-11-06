package main

import (
	"context"
	"fenbus/log"
	"fenbus/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./fenbus.log")
	host, port := "localhost", "8080"
	ctx, err := service.Start(context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Log Service stopped")
}
