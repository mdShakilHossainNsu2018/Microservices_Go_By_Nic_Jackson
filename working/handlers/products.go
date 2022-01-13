package handlers

import (
	"fmt"
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
		p.getProducts(writer, request)
		return
	}

	if request.Method == http.MethodPost {
		p.addProduct(writer, request)
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
		p.updateProduct(id, writer, request)
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

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("add Product called")
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)

	if err != nil {
		http.Error(writer, "Unable to unmarshal", http.StatusBadRequest)
	}
	data.AddProduct(*prod)
	p.l.Printf("Prod: %v", prod)
}

func (p *Products) updateProduct(id int, writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Update product called")
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)

	if err != nil {
		http.Error(writer, "Unable to unmarshal", http.StatusBadRequest)
	}

	err2 := prod.UpdateProduct(id, prod)
	if err2 == data.ErrProductNotFound {
		http.Error(writer, "Not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(writer, "Unknown error", http.StatusInternalServerError)
	}
}
