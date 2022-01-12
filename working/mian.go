package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(request.Body)

		if err != nil {
			http.Error(writer, "Opps!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(writer, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Good bye")
	})

	http.ListenAndServe(":9090", nil)
}
