package dto

import (
	"strings"

	"github.com/amalmohann/banking/errs"
)

// Request bodies
type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// Validators
// create new account
func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000.00 {
		return errs.ValidationError("Minimum Balance is 5000 or above to create a new account.")
	}
	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return errs.ValidationError("Account type should be checking or savings")
	}
	return nil
}
