package domain

import (
	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/errs"
)

type Customer struct {
	Id          string `json:"id" db:"customer_id"`
	Name        string `json:"name" db:"name"`
	City        string `json:"city" db:"city"`
	Zip         string `json:"zip" db:"zipcode"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status" db:"status"`
}

func (c Customer) toStatusText() string {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}

func (c Customer) ToResponse() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zip:         c.Zip,
		DateOfBirth: c.DateOfBirth,
		Status:      c.toStatusText(),
	}
	return response
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
