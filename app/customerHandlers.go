package app

import (
	"encoding/json"
	"fmt"
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
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// find customer by id handler
func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id := p["customer_id"]
	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		writeResponse(w, err.Status, err.ToResponse())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
