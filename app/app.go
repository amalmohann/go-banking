package app

import (
	"log"
	"net/http"

	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// defining routes
	router.HandleFunc("/", healthCheck).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
