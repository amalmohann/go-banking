package service

import (
	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/errs"
)

//services
type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
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
func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToResponse()
	return &response, nil
}

// helper function to connect to default service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
