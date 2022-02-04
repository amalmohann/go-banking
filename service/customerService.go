package service

import (
	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/errs"
)

//services
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

// default repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

//functions
func toStatusCode(status string) string {
	statusCode := ""
	if status == "active" {
		statusCode = "1"
	} else if status == "inactive" {
		statusCode = "0"
	}
	return statusCode
}

// get all customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	var response []dto.CustomerResponse
	statusCode := toStatusCode(status)
	c, err := s.repo.FindAll(statusCode)
	if err != nil {
		return nil, err
	}
	for _, c := range c {
		response = append(response, c.ToResponse())
	}
	return response, nil
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
