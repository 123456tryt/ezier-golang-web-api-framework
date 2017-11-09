package main

import (
	_ "ezier/models"
	"ezier/routers"
	"fmt"

	"flag"
	"net/http"
	"time"
)

func main() {
	addr := flag.String("addr", ":9999", "Port to listen on")
	server := &http.Server{
		Handler:      routers.HttpServeMux,
		Addr:         *addr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	server.ListenAndServe()
	fmt.Println("server is one port 9999")
}
