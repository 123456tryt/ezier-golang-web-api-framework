package main

import (
	"flag"
	"fmt"
	"my_go_web_framework/routers"
	"net/http"
	"time"
)

func main() {
	addr := flag.String("addr", ":9123", "Port to listen on")

	srv := &http.Server{
		Handler:      routers.HttpServeMux,
		Addr:         *addr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("serve@127.0.01", addr)

}
