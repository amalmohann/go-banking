package service

import "github.com/amalmohann/banking/domain"

//services
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, error)
}

// default repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// get all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// get customers by id
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	return s.repo.ById(id)
}

// helper function to connect to default service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
