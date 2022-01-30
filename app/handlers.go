package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zip  string `json:"zip" xml:"zip"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Live")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Amal Mohan N", City: "Kakkodi", Zip: "673611"},
		{Name: "Amal M", City: "Karaparamba", Zip: "673005"},
	}
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(vars)
	fmt.Fprintf(w, "Post method")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s\n", vars["customer_id"])

}
