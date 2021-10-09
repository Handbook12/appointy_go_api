package main

import (
	"go-rest-mongodb/config"
	"go-rest-mongodb/routers"
	"net/http"
	"os"
	"time"
)

var config Config

func init() {
	config.Read()
}

func main() {
	r := routers.Routers()
	srv := &http.Server {
		Handler: r,
		Addr:    config.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
