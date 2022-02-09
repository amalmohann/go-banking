package domain

import (
	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToResponse() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) ToRequest() dto.NewAccountRequest {
	return dto.NewAccountRequest{
		CustomerId:  a.CustomerId,
		AccountType: a.AccountType,
		Amount:      a.Amount,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
