package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAllCustomers() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Amal Mohan N", City: "Kozhikode", Zip: "673611", DateOfBirth: "1996-07-27", Status: "1"},
		{Id: "1002", Name: "Amrutha TP", City: "Kozhikode", Zip: "673611", DateOfBirth: "1997-10-30", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
