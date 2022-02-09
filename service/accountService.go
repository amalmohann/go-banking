package service

import (
	"time"

	"github.com/amalmohann/banking/domain"
	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

//service
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

// default repository
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToResponse()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.IsWithdrawal() {
		account, err := s.repo.GetAccountById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.ValidationError("Insufficient balance in this account.")
		}
	}
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}

	response := transaction.ToResponse()
	return &response, nil

}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}
