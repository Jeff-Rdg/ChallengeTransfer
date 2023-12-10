package repository

import (
	"ChallengeBackEndPP/user"
	"errors"
	"gorm.io/gorm"
)

var (
	NoRowsAffectedErr = errors.New("no rows affected")
	ExistUserErr      = errors.New("there are already users with the taxnumber or email provided")
)

type UserDb struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *UserDb {
	return &UserDb{
		Db: db,
	}
}

func (us *UserDb) Create(u *user.User) (int, error) {
	insert := us.Db.Create(&u)
	if insert.RowsAffected == 0 {
		return 0, NoRowsAffectedErr
	}

	if insert.Error != nil {
		return 0, insert.Error
	}

	return 1, nil
}

func (us *UserDb) GetById(id uint) (*user.User, error) {
	var findUser *user.User

	err := us.Db.Where("id = ?", id).First(&findUser).Error
	if err != nil {
		return nil, err
	}

	return findUser, nil
}

func (us *UserDb) FindByTaxNumberOrEmail(taxNumber, email string) (int, error) {
	var users []user.User
	err := us.Db.Where("tax_number = ?", taxNumber).Or("email = ?", email).Find(&users).Error
	if err != nil {
		return 0, err
	}

	if len(users) > 0 {
		return 0, ExistUserErr
	}

	return 1, nil
}
