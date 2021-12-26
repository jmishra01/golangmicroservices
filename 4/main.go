package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"microservice.com/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "[Golang] ", log.LstdFlags)
	new_hello := handlers.NewHello(l)
	new_golang := handlers.NewGolang(l)
	new_products := handlers.NewProducts(l)

	newServeMux := http.NewServeMux()
	newServeMux.Handle("/", new_hello)
	newServeMux.Handle("/golang", new_golang)
	newServeMux.Handle("/product", new_products)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      newServeMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)
}
