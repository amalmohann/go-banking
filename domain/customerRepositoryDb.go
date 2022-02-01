package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// repository
type CustomerRepositoryDb struct {
	dbClient *sql.DB
}

// implementing the interfaces

// FindAll()
func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)
	query := "SELECT * FROM customers"
	row, err := db.dbClient.Query(query)
	if err != nil {
		log.Print("Error Fetching from database: ", err.Error())
		return nil, err
	}
	for row.Next() {
		var c Customer
		err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zip, &c.Status)
		if err != nil {
			log.Print("Error scanning Customer list from database: ", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// FindById()
func (db CustomerRepositoryDb) ById(id string) (*Customer, error) {
	var c Customer
	query := "SELECT * FROM customers WHERE customer_id = ?"
	row := db.dbClient.QueryRow(query, id)
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zip, &c.Status)
	if err != nil {
		log.Print("Error while scanning the customer: ", err.Error())
		return nil, err
	}
	return &c, nil
}

// Helper Function to create new Db connection
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbClient, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return CustomerRepositoryDb{dbClient: dbClient}
}
