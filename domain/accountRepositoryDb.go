package domain

import (
	"database/sql"
	"strconv"

	"github.com/amalmohann/banking/errs"
	"github.com/amalmohann/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

// save account
func (db AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	insQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?,?,?,?,?)"
	result, err := db.client.Exec(insQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Unexpected error from Database : " + err.Error())
		return nil, errs.InternalServerError("Unexpected error from Database : " + err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Unexpected error from Database : " + err.Error())
		return nil, errs.InternalServerError("Unexpected error from Database : " + err.Error())
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

// Find account by id
func (db AccountRepositoryDb) GetAccountById(id string) (*Account, *errs.AppError) {
	var acc Account
	fetchQuery := "SELECT * from accounts where account_id = ?"
	if err := db.client.Get(&acc, fetchQuery, id); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No account found : " + err.Error())
			return nil, errs.NotFoundError("No account found : " + err.Error())
		} else {
			logger.Error("Unexpected error from Database : " + err.Error())
			return nil, errs.InternalServerError("Unexpected error from Database : " + err.Error())
		}
	}
	return &acc, nil
}

// save transaction
func (db AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// begin transaction
	txn, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.InternalServerError("Error while starting a new transaction for bank account transaction: " + err.Error())
	}

	// adding transaction record
	query := `INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`
	result, err := txn.Exec(query, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while processing transaction: " + err.Error())
		return nil, errs.InternalServerError("Error while processing transaction: " + err.Error())
	}

	// updating balance
	if t.IsWithdrawal() {
		_, err = txn.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = txn.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// rollback in case of error
	if err != nil {
		txn.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.InternalServerError("Unexpected database error")
	}

	// commit if everything goes well
	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.InternalServerError("UUnexpected database error")
	}

	//getting transaction id from the result
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.InternalServerError("Unexpected database error")
	}

	// fetching the latest account details
	account, appErr := db.GetAccountById(t.AccountId)
	if appErr != nil {
		logger.Error("Error while fetching account details")
		return nil, appErr
	}

	// updating the txn object
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
