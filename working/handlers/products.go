package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdShakilHossainNsu2018/Microservices_Go_By_Nic_Jackson/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//log.Println("products called")
	if request.Method == http.MethodGet {
		log.Println("products called")
		p.GetProducts(writer, request)
		return
	}

	if request.Method == http.MethodPost {
		p.AddProduct(writer, request)
	}

	if request.Method == http.MethodPut {
		p.l.Println("Put Method called.")
		r := regexp.MustCompile("/([0-9]+)")
		g := r.FindAllStringSubmatch(request.URL.Path, -1)

		if len(g) != 1 {
			http.Error(writer, "Invalid URL ", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(writer, "Invalid URL ", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(writer, "Invalid URL ", http.StatusBadRequest)
			return
		}
		p.l.Println("id: ", id)
		p.UpdateProduct(writer, request)
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)

}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {

	lp := data.GetProducts()

	err := lp.ToJSON(writer)
	if err != nil {
		sErr := fmt.Sprintf("Unable to marshal json: %v", err)
		http.Error(writer, sErr, http.StatusInternalServerError)

	}
}

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("add Product called")
	prod := request.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Update product called")
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {

	}
	prod := request.Context().Value(KeyProduct{}).(data.Product)

	err3 := data.UpdateProduct(id, &prod)
	if err3 == data.ErrProductNotFound {
		http.Error(writer, "Not found", http.StatusNotFound)
	}

	if err3 != nil {
		http.Error(writer, "Unknown error", http.StatusInternalServerError)
	}
}

type KeyProduct struct {
}

func (p Products) MiddlewareProduct(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		prod := data.Product{}
		err2 := prod.FromJSON(request.Body)

		if err2 != nil {
			http.Error(writer, "Unable to unmarshal", http.StatusBadRequest)
			return
		}
		//p.l.Println(prod)

		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		req := request.WithContext(ctx)
		next.ServeHTTP(writer, req)

	})
}
