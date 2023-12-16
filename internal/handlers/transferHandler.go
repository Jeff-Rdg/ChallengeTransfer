package handlers

import (
	"ChallengeBackEndPP/transfer"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransferHandler struct {
	TransferService transfer.UseCase
}

func NewTransferHandler(service transfer.UseCase) *TransferHandler {
	return &TransferHandler{TransferService: service}
}

func (t *TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var req transfer.Request
	var message string

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		message = "JSON deserialization error"
		RenderJSON(w, http.StatusBadRequest, message, nil)
		return
	}

	id, err := t.TransferService.Create(req)
	if err != nil {
		message = "error to create transfer"
		RenderJSON(w, http.StatusBadRequest, message, err)
		return
	}

	result := fmt.Sprintf("transfer with id %v successful", id)
	RenderJSON(w, http.StatusCreated, result, nil)
}
