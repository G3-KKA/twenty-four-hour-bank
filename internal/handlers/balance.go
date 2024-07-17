package handlers

import (
	"bank24/domain"
	"bank24/internal/database/requests"
	"encoding/json"
	"net/http"
)

// GET
func (is *Interstate) Balance(w http.ResponseWriter, r *http.Request) {
	requestId, err := urlGetId(r.URL.Path, "accounts")
	serviceRequest := requests.Balance(domain.UserID(requestId))
	rsp := is.service.Send(serviceRequest)
	if rsp.Err != nil {
		logError("Balance", r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	brsp, ok := rsp.Data.(requests.BalanceResponse)
	if !ok {
		logError("Balance", r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(brsp)
	if err != nil {
		logError("Balance", r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
