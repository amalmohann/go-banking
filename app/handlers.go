package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/amalmohann/banking/service"
	"github.com/gorilla/mux"
)

//handler service
type CustomerHandlers struct {
	service service.CustomerService
}

// health check handler
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// Find all customers handler
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
}

// find customer by id handler
func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	p := r.URL.Query()
	log.Print("back: ", p)
	log.Print(params)
	id := params["customer_id"]
	customer, _ := ch.service.GetCustomerById(id)
	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Customer Not Found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
