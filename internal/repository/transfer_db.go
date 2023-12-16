package repository

import (
	"ChallengeBackEndPP/transfer"
	"gorm.io/gorm"
)

type TransferDb struct {
	Db *gorm.DB
}

func NewTransfer(db *gorm.DB) *TransferDb {
	return &TransferDb{Db: db}
}

func (tr *TransferDb) Create(transaction *transfer.Transfer) (int, error) {
	insert := tr.Db.Create(&transaction)

	if insert.Error != nil {
		return 0, insert.Error
	}

	if insert.RowsAffected == 0 {
		return 0, NoRowsAffectedErr
	}

	return int(transaction.ID), nil
}

func (tr *TransferDb) GetById(id uint) (*transfer.Transfer, error) {
	var findTransfer *transfer.Transfer

	err := tr.Db.Where("id = ?", id).First(&findTransfer).Error
	if err != nil {
		return nil, err
	}

	return findTransfer, nil
}
