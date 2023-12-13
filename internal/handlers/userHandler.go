package handlers

import (
	"ChallengeBackEndPP/user"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
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

	id, err := h.UserService.Create(req)
	if err != nil {
		message = "error to create user"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	result := fmt.Sprintf("User with id %v created successfully", id)
	RenderJSON(w, http.StatusCreated, result, nil)
}

func (h *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	var message string

	if param == "" {
		message = "no id informed"
		RenderJSON(w, http.StatusBadRequest, message, nil)
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		message = "error to find user"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	res, err := h.UserService.GetById(uint(id))
	if err != nil {
		message = "error to find user"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	RenderJSON(w, http.StatusOK, "", res)
}

func (h *UserHandler) AddMoney(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	var message string
	var req user.AddValueRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		message = "JSON deserialization error"
		RenderJSON(w, http.StatusBadRequest, message, nil)
		return
	}

	if param == "" {
		message = "no id informed"
		RenderJSON(w, http.StatusBadRequest, message, nil)
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		message = "error to add money"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	err = h.UserService.AddMoney(uint(id), req.Value)
	if err != nil {
		message = "error to add money"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	result := fmt.Sprintf("User with id %v updated successfully", id)
	RenderJSON(w, http.StatusCreated, result, nil)
}
