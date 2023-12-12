package handlers

import (
	"ChallengeBackEndPP/wallet"
	"net/http"
)

type WalletHandler struct {
	WalletService wallet.UseCase
}

func NewWalletHandler(service wallet.UseCase) *WalletHandler {
	return &WalletHandler{WalletService: service}
}

func (s *WalletHandler) FindWalletById(w http.ResponseWriter, r *http.Request) {

}

func (s *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {

}
