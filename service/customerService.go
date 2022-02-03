package service

import (
	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/errs"
)

//services
type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
}

// default repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// get all customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {

	statusCodes := map[string]string{
		"inactive": "0",
		"active":   "1",
	}

	if status == "active" || status == "inactive" {
		status = statusCodes[status]
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

// get customers by id
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// helper function to connect to default service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
