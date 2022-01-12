package handlers

import (
	"fmt"
	"github.com/mdShakilHossainNsu2018/Microservices_Go_By_Nic_Jackson/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//log.Println("products called")
	if request.Method == http.MethodGet {
		log.Println("products called")
		p.getProducts(writer, request)
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)

}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(writer)
	if err != nil {
		sErr := fmt.Sprintf("Unable to marshal json: %v", err)
		http.Error(writer, sErr, http.StatusInternalServerError)

	}
}
