package app

import (
	"log"
	"net/http"
	"time"

	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	// router
	router := mux.NewRouter()

	// get dbClient
	dbClient := getDbClient()

	// repositories
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandlers{service.NewAccountService(accountRepositoryDb)}

	// dummy wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// defining routes
	router.HandleFunc("/", healthCheck).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)

	// start server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDbClient() *sqlx.DB {
	dbClient, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return dbClient
}
