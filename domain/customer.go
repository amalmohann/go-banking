package domain

type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zip         string `json:"zip" xml:"zip"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}
