package handler

import (
	"ChallengeBackEndPP/user"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserService user.UseCase
}

func NewUserHandler(service user.UseCase) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req user.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON deserialization error"))
		return
	}

	_, err = h.UserService.Create(req.FullName, req.TaxNumber, req.Email, req.Password, req.IsShopkeeper)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}
