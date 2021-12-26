package handlers

import (
	"log"
	"net/http"
)

type Golang struct {
	l *log.Logger
}

func NewGolang(l *log.Logger) *Golang {
	return &Golang{l}
}

func (h *Golang) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Golang")
	rw.Write([]byte("Hello Golang!!"))

}
