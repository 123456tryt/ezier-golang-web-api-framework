package main

import (
	_ "ezier/models"
	"ezier/routers"
	"flag"
	"fmt"

	"net/http"
	"time"
)

func init() {
}

func main() {
	port := flag.String("port", ":1234", "help message for flagname")
	flag.Parse()

	//http://qefee.com/2014/02/02/go%E8%AF%AD%E8%A8%80%E7%9A%84flag%E5%8C%85%E7%AE%80%E5%8D%95%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/
	server := &http.Server{
		Handler:      routers.HttpServeMux,
		Addr:         *port,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	errrr := server.ListenAndServe()
	if errrr != nil {
		fmt.Println(errrr)

	} else {
		fmt.Println("server is one port", *port)

	}
}
