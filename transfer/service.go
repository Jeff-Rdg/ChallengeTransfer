package transfer

import (
	"ChallengeBackEndPP/user"
	"gorm.io/gorm"
)

type Service struct {
	repo     Repository
	userRepo user.Repository
	db       *gorm.DB
}

func NewService(rep Repository, usRepo user.Repository, _db *gorm.DB) *Service {
	return &Service{
		repo:     rep,
		userRepo: usRepo,
		db:       _db,
	}
}
func (s *Service) GetById(id uint) (*Transfer, error) {
	return nil, nil
}

func (s *Service) Create(request Request) (int, error) {
	transfer, err := NewTransfer(request)
	if err != nil {
		return 0, err
	}
	// verificar se usuários existem
	sender, err := s.userRepo.GetById(request.SenderId)
	if err != nil {
		return 0, err
	}

	receiver, err := s.userRepo.GetById(request.ReceiverId)
	if err != nil {
		return 0, err
	}

	if sender.IsShopkeeper {
		return 0, InvalidTransferErr
	}

	if sender.Balance < request.Value {
		return 0, InvalidTransferErr
	}

	sender.Balance -= request.Value
	receiver.Balance += request.Value

	// implementar transações
	s.db.Transaction(func(tx *gorm.DB) error {
	err = tx.Updates(&sender)
		if err != nil {
		return err	
		}
	err = tx.Updates(&receiver)
		if err != nil {
		return err	
		}
	err = tx.Create(&transfer).Error
		if err != nil {
		return err	
		}
	})
	
	/*
	tx := s.db.Begin()
	_, err = s.userRepo.Update(sender)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	_, err = s.userRepo.Update(receiver)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	transactionId, err := s.repo.Create(transfer)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
 */
	return int(transfer.ID), nil
}
