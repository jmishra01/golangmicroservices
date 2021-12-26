package main

import (
	"log"
	"net/http"
	"os"

	"microservice.com/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "[Golang] ", log.LstdFlags)
	new_hello := handlers.NewHello(l)
	new_golang := handlers.NewGolang(l)

	newServeMux := http.NewServeMux()
	newServeMux.Handle("/", new_hello)
	newServeMux.Handle("/golang", new_golang)

	http.ListenAndServe(":9090", newServeMux)
}
