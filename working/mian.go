package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mdShakilHossainNsu2018/Microservices_Go_By_Nic_Jackson/working/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.Lshortfile)
	//hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)
	gh := handlers.NewGoodBye(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	putRouter.Use(ph.MiddlewareProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProduct)
	//sm.Handle("/", ph)
	//sm.Handle("/products", ph)
	sm.Handle("/goodbye", gh)

	s := &http.Server{Addr: ":9090", Handler: sm, IdleTimeout: 120 * time.Second, ReadTimeout: 1 * time.Second, WriteTimeout: 1 * time.Second}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	println("\nTerminating ", sig.String())
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(tc)
	if err != nil {
		l.Fatalln(err)
	}
}
