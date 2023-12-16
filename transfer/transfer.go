package transfer

import (
	"errors"
	"gorm.io/gorm"
)

var (
	SenderIdNilErr   = errors.New("sender_id is required")
	ReceiverIdNilErr = errors.New("receiver_id is required")
	ValueNilErr      = errors.New("value is required")
)

type Transfer struct {
	gorm.Model
	SenderId   uint    `json:"sender_id"`
	ReceiverId uint    `json:"receiver_id"`
	Value      float64 `json:"value"`
}

type Request struct {
	SenderId   uint    `json:"sender_id"`
	ReceiverId uint    `json:"receiver_id"`
	Value      float64 `json:"value"`
}

type Reader interface {
	GetById(id uint) (*Transfer, error)
}

type Writer interface {
	Create(t *Transfer) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetById(id uint) (*Transfer, error)
	Create(request Request) (int, error)
}

func NewTransfer(request Request) (*Transfer, error) {
	err := validateTransfer(request.SenderId, request.ReceiverId, request.Value)
	if err != nil {
		return nil, err
	}

	return &Transfer{
		Model:      gorm.Model{},
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Value:      request.Value,
	}, nil
}

func validateTransfer(senderId, receiverId uint, value float64) error {
	var errs []error
	if senderId == 0 {
		errs = append(errs, SenderIdNilErr)
	}
	if receiverId == 0 {
		errs = append(errs, ReceiverIdNilErr)
	}
	if value == 0 {
		errs = append(errs, ValueNilErr)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
