package user

import (
	"errors"
	"gorm.io/gorm"
)

var (
	FullNameNilErr  = errors.New("full_name is required")
	TaxNumberNilErr = errors.New("tax_number is required")
	EmailNilErr     = errors.New("email is required")
	PasswordNilErr  = errors.New("password is required")
)

type User struct {
	gorm.Model
	FullName     string `json:"full_name"`
	TaxNumber    string `json:"tax_number" gorm:"not null"`
	Email        string `json:"email" gorm:"not null"`
	Password     string `json:"password"`
	IsShopkeeper bool   `json:"is_shopkeeper"`
}

type Request struct {
	FullName     string `json:"full_name"`
	TaxNumber    string `json:"tax_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopkeeper bool   `json:"is_shopkeeper"`
}

func NewUser(fullname, taxnumber, email, password string, isshopkeeper bool) (*User, error) {
	err := validateUser(fullname, taxnumber, email, password)
	if err != nil {
		return nil, err
	}
	return &User{
		FullName:     fullname,
		TaxNumber:    taxnumber,
		Email:        email,
		Password:     password,
		IsShopkeeper: isshopkeeper,
	}, nil
}

func validateUser(fullname, taxnumber, email, password string) error {
	var errs []error
	if fullname == "" {
		errs = append(errs, FullNameNilErr)
	}
	if taxnumber == "" {
		errs = append(errs, TaxNumberNilErr)
	}
	if email == "" {
		errs = append(errs, EmailNilErr)
	}
	if password == "" {
		errs = append(errs, PasswordNilErr)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

type Reader interface {
	GetById(id uint) (*User, error)
	FindByTaxNumberOrEmail(taxNumber, email string) (int, error)
}

type Writer interface {
	Create(u *User) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetById(id uint) (*User, error)
	Create(fullname, taxnumber, email, password string, isshopkeeper bool) (int, error)
}
