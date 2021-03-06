package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amalmohann/banking/dto"
	"github.com/amalmohann/banking/logger"
	"github.com/amalmohann/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h *AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	p := mux.Vars(r)
	customerId := p["customer_id"]
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(strconv.Itoa(http.StatusBadRequest) + " Bad Request : " + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, err := h.service.NewAccount(request)
		if err != nil {
			logger.Error(strconv.Itoa(err.Status) + err.Message)
			writeResponse(w, err.Status, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}

func (h *AccountHandlers) makeTransaction(w http.ResponseWriter, r *http.Request) {
	// get details from url
	p := mux.Vars(r)
	customerId := p["customer_id"]
	accountId := p["account_id"]

	// decoding the Request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(strconv.Itoa(http.StatusBadRequest) + " Bad Request : " + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		// build request
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, err := h.service.MakeTransaction(request)
		if err != nil {
			logger.Error(strconv.Itoa(err.Status) + err.Message)
			writeResponse(w, err.Status, err.Message)
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
