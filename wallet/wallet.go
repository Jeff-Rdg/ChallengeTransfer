package wallet

import (
	"errors"
	"gorm.io/gorm"
)

var UserIdNilErr = errors.New("user_id is required")

type Reader interface {
	GetById(id uint) (*Wallet, error)
	GetWalletByUserId(id uint) (*Wallet, error)
}

type Writer interface {
	Create(u *Wallet) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetById(id uint) (*Wallet, error)
	Create(request Request) (int, error)
}

type Wallet struct {
	gorm.Model
	Balance float64 `json:"balance"`
	UserID  uint    `json:"user_id"`
}

type Request struct {
	Balance float64 `json:"balance"`
	UserID  uint    `json:"user_id"`
}

func NewWallet(request Request) (*Wallet, error) {
	err := validateWallet(request.UserID)
	if err != nil {
		return nil, err
	}
	return &Wallet{
		Balance: request.Balance,
		UserID:  request.UserID,
	}, nil
}

func validateWallet(userid uint) error {
	if userid == 0 {
		return UserIdNilErr
	}
	return nil
}
