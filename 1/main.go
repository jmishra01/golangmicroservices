package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
