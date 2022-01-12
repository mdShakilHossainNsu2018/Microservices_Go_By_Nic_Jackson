package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l: l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_, err := rw.Write([]byte("Good bye"))
	if err != nil {
		log.Fatalf("Error while writing in Good Bye")

	}
}
