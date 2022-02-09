package dto

import (
	"strings"

	"github.com/amalmohann/banking/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	CustomerId      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 00.00 {
		return errs.ValidationError("Transaction cannot be proceeded with a negative amount.")
	}
	if !r.IsWithdrawal() && !r.IsDeposit() {
		return errs.ValidationError("Transaction should be withdraw or deposit.")
	}
	return nil
}

func (r TransactionRequest) IsWithdrawal() bool {
	return strings.ToLower(r.TransactionType) == WITHDRAWAL
}

func (r TransactionRequest) IsDeposit() bool {
	return strings.ToLower(r.TransactionType) == DEPOSIT
}
