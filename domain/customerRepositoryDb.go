package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	dbClient *sql.DB
}

func (db CustomerRepositoryDb) FindAllCustomers() ([]Customer, error) {
	customers := make([]Customer, 0)
	findAllQuery := "SELECT * FROM customers"
	row, err := db.dbClient.Query(findAllQuery)
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
