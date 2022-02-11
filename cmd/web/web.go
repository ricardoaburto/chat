package main

import (
	"WebSockets/Chat/ws/internal/handlers"
	"log"
	"net/http"
)

func main()  {
	mux:=routes()

	log.Println("starting channel listener")

	go handlers.ListenToWsChannel()

	log.Println("starting web server on port 8080")

	err :=http.ListenAndServe(":8080",mux)
	if err != nil {
		log.Println(err)
	}

}
