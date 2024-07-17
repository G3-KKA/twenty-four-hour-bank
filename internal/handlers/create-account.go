package handlers

import (
	"bank24/domain"
	"bank24/internal/database/requests"
	log "bank24/internal/logger"
	"encoding/json"
	"io"
	"net/http"
)

func (is *Interstate) CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logError("CreateAccount", r, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request := struct {
		Id      int     `json:"id"`
		Balance float64 `json:"balance"`
	}{}

	err = json.Unmarshal(body, &request)

	log.Info("CreateAccount handling for", r.RemoteAddr, " as:", r.Method, " to: ", r.URL.Path, " userid:", request.Id)
	if err != nil {
		logError("CreateAccount", r, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	serviceRequest := requests.NewAccount(domain.UserID(request.Id), request.Balance)
	rsp := is.service.Send(serviceRequest)
	if rsp.Err != nil {
		logError("CreateAccount", r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Info("CreateAccount success for", r.RemoteAddr, " as:", r.Method, " to: ", r.URL.Path)
	w.WriteHeader(http.StatusOK)
}
