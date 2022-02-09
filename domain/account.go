package domain

import (
	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
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

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	GetAccountById(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}
