package handlers

import (
	"ChallengeBackEndPP/user"
	"encoding/json"
	"fmt"
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
	var message string
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		message = "JSON deserialization error"
		RenderJSON(w, http.StatusBadRequest, message, nil)
		return
	}

	id, err := h.UserService.Create(req.FullName, req.TaxNumber, req.Email, req.Password, req.IsShopkeeper)
	if err != nil {
		message = "error to create user"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	result := fmt.Sprintf("User with id %v created successfully", id)
	RenderJSON(w, http.StatusOK, result, nil)
}
