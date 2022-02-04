package domain

import (
	"database/sql"
	"time"

	"github.com/amalmohann/banking/errs"
	"github.com/amalmohann/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// repository
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// implementing the interfaces
// FindAll()
func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	query := "SELECT * FROM customers"
	if status == "1" || status == "0" {
		query = query + " where status = " + status
	}
	err := db.client.Select(&customers, query)
	if err != nil {
		logger.Error("Error Fetching from database: " + err.Error())
		return nil, errs.InternalServerError("Error Fetching from database: " + err.Error())
	}
	return customers, nil
}

// FindById()
func (db CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	query := "SELECT * FROM customers WHERE customer_id = ?"
	logger.Info(id)
	err := db.client.Get(&c, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error Fetching from database: " + err.Error())
			return nil, errs.NotFoundError("Customer Not Found!")
		} else {
			logger.Error("Error Fetching from database: " + err.Error())
			return nil, errs.InternalServerError("Unexpected database error: " + err.Error())
		}
	}
	return &c, nil
}

// Helper Function to create new Db connection
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbClient, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: dbClient}
}
