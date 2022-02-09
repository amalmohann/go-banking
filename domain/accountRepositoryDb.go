package domain

import (
	"strconv"

	"github.com/amalmohann/banking/errs"
	"github.com/amalmohann/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

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

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
