package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/amalmohann/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	// []domain.Customer{
	// 	{Id: "1001", Name: "Amal Mohan N", City: "Kozhikode", Zip: "673611", DateOfBirth: "1996-07-27", Status: "1"},
	// 	{Id: "1002", Name: "Amrutha TP", City: "Kozhikode", Zip: "673611", DateOfBirth: "1997-10-30", Status: "1"},
	// }
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Add("Content-Type", "application/json")
// 	// json.NewEncoder(w).Encode(vars)
// 	fmt.Fprintf(w, "Post method")
// }

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprintf(w, "%s\n", vars["customer_id"])

// }
