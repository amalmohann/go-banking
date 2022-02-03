package service

import (
	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/errs"
)

//services
type CustomerService interface {
	GetAllCustomers(...string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
}

// default repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// get all customers
func (s DefaultCustomerService) GetAllCustomers(params ...string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(params[0])
}

// get customers by id
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// helper function to connect to default service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
