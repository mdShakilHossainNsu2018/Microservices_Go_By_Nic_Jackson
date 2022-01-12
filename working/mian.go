package main

import (
	"github.com/mdShakilHossainNsu2018/Microservices_Go_By_Nic_Jackson/working/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.Lshortfile)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
