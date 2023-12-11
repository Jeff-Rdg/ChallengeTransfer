package repository

import (
	"ChallengeBackEndPP/wallet"
	"gorm.io/gorm"
)

type WalletDB struct {
	Db *gorm.DB
}

func NewWallet(db *gorm.DB) *WalletDB {
	return &WalletDB{Db: db}
}

func (w *WalletDB) Create(u *wallet.Wallet) (int, error) {
	insert := w.Db.Create(&u)
	if insert.RowsAffected == 0 {
		return 0, NoRowsAffectedErr
	}

	if insert.Error != nil {
		return 0, insert.Error
	}

	return int(u.ID), nil
}

func (w *WalletDB) GetById(id uint) (*wallet.Wallet, error) {
	var findWallet *wallet.Wallet

	err := w.Db.Where("id = ?", id).First(&findWallet).Error
	if err != nil {
		return nil, err
	}
	return findWallet, nil
}

func (w *WalletDB) GetWalletByUserId(id uint) (*wallet.Wallet, error) {
	var findWallet *wallet.Wallet

	err := w.Db.Where("user_id = ?", id).First(&findWallet).Error
	if err != nil {
		return nil, err
	}
	return findWallet, nil
}
