package handlers

import (
	"bank24/domain"
	"bank24/internal/database/requests"
	"encoding/json"
	"io"
	"net/http"
)

func (is *Interstate) Withdraw(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logError("Withdraw", r, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request := struct {
		Id       int     `json:"id"`
		Withdraw float64 `json:"withdraw"`
	}{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		logError("Withdraw", r, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	serviceRequest := requests.Withdraw(domain.UserID(request.Id), request.Withdraw)
	rsp := is.service.Send(serviceRequest)
	if rsp.Err != nil {
		logError("Withdraw", r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
