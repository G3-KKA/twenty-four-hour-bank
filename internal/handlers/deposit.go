package handlers

import (
	"bank24/domain"
	"bank24/internal/database/requests"
	"encoding/json"
	"io"
	"net/http"
)

func (is *Interstate) Deposit(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logError("Deposit", r, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request := struct {
		Id      int     `json:"id"`
		Deposit float64 `json:"deposit"`
	}{}

	err = json.Unmarshal(body, &request)
	if err != nil {
		logError("Deposit", r, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	serviceRequest := requests.Deposit(domain.UserID(request.Id), request.Deposit)
	rsp := is.service.Send(serviceRequest)
	if rsp.Err != nil {
		logError("Deposit", r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
