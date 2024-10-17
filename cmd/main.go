package main

import (
	"log"

	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(httpapi.Server)
	if err := srv.Run("8080", handlers.InitRourers()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
