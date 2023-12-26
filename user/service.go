package user

import (
	"errors"
	"gorm.io/gorm"
)

var (
	NoRowsAffectedErr = errors.New("no rows affected")
)

type Service struct {
	Db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{Db: db}
}

func (s *Service) GetById(id uint) (*Response, error) {
	var user *User
	err := s.Db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user.MapUserToResponse(), nil
}

func (s *Service) Create(request Request) (int, error) {
	newUser, err := NewUser(request)
	if err != nil {
		return 0, err
	}

	insert := s.Db.Create(&newUser)
	if insert.Error != nil {
		return 0, insert.Error
	}

	if insert.RowsAffected == 0 {
		return 0, NoRowsAffectedErr
	}

	return int(newUser.ID), nil
}

func (s *Service) Update(id uint, request RequestUpdate) (int, error) {
	var user *User
	err := s.Db.First(&user, id).Error
	if err != nil {
		return 0, err
	}
	user.UpdateDiffFields(request)
	err = s.Db.Save(&user).Error
	if err != nil {
		return 0, err
	}

	return int(user.ID), nil
}

func (s *Service) AddMoney(id uint, value float64) error {
	var user *User
	err := s.Db.First(&user, id).Error
	if err != nil {
		return err
	}

	if user.IsShopkeeper {
		return InvalidAddBalanceErr
	}
	user.Balance += value
	err = s.Db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}
