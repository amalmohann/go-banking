package dto

import "github.com/amalmohann/banking/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000.00 {
		return errs.ValidationError("Minimum Balance is 5000 or above to create a new account.")
	}
	if r.AccountType != "savings" && r.AccountType != "checking" {
		return errs.ValidationError("Account type should be checking or savings")
	}
	return nil
}
