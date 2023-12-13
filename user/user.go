package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

var (
	FullNameNilErr       = errors.New("full_name is required")
	TaxNumberNilErr      = errors.New("tax_number is required")
	EmailNilErr          = errors.New("email is required")
	InvalidEmailErr      = errors.New("email invalid")
	PasswordNilErr       = errors.New("password is required")
	FullNameInvalidErr   = errors.New("full_name must not contain special characters or numbers")
	TaxNumberInvalidErr  = errors.New("tax_number invalid")
	InvalidAddBalanceErr = errors.New("shopkeeper is not allowed to add money")
)

type Reader interface {
	GetById(id uint) (*User, error)
	FindByTaxNumberOrEmail(taxNumber, email string) (int, error)
}

type Writer interface {
	Create(u *User) (int, error)
	Update(u *User) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetById(id uint) (*Response, error)
	Create(request Request) (int, error)
	AddMoney(id uint, value float64) error
}

type User struct {
	gorm.Model
	FullName     string `json:"full_name"`
	TaxNumber    string `json:"tax_number" gorm:"not null"`
	Email        string `json:"email" gorm:"not null"`
	Password     string `json:"password"`
	IsShopkeeper bool   `json:"is_shopkeeper"`
	Wallet       `json:"wallet"`
}

type Wallet struct {
	Balance float64 `json:"balance"`
}

type Request struct {
	FullName     string `json:"full_name"`
	TaxNumber    string `json:"tax_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopkeeper bool   `json:"is_shopkeeper"`
}

type Response struct {
	gorm.Model
	FullName     string `json:"full_name"`
	TaxNumber    string `json:"tax_number"`
	Email        string `json:"email"`
	IsShopkeeper bool   `json:"is_shopkeeper"`
	Wallet       Wallet `json:"wallet"`
}

type AddValueRequest struct {
	Value float64 `json:"value"`
}

func NewUser(request Request) (*User, error) {
	err := validateUser(request.FullName, request.TaxNumber, request.Email, request.Password, request.IsShopkeeper)
	if err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FullName:     request.FullName,
		TaxNumber:    request.TaxNumber,
		Email:        request.Email,
		Password:     string(hash),
		IsShopkeeper: request.IsShopkeeper,
		Wallet:       Wallet{},
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func validateUser(fullname, taxnumber, email, password string, isshopkeeper bool) error {
	var errs []error
	if fullname == "" {
		errs = append(errs, FullNameNilErr)
	}
	if !isValidName(fullname) {
		errs = append(errs, FullNameInvalidErr)
	}

	if taxnumber == "" {
		errs = append(errs, TaxNumberNilErr)
	}

	if isshopkeeper {
		if !isValidCNPJ(taxnumber) {
			errs = append(errs, TaxNumberInvalidErr)
		}
	} else {
		if !isValidCPF(taxnumber) {
			errs = append(errs, TaxNumberInvalidErr)
		}
	}

	if email == "" {
		errs = append(errs, EmailNilErr)
	}
	if !isValidEmail(email) {
		errs = append(errs, InvalidEmailErr)
	}

	if password == "" {
		errs = append(errs, PasswordNilErr)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func isValidName(name string) bool {
	regexName := regexp.MustCompile("^[a-zA-ZÀ-ÖØ-öø-ÿ\\s]+$")

	return regexName.MatchString(name)
}

func isValidCNPJ(taxnumber string) bool {
	taxnumber = regexp.MustCompile("\\D").ReplaceAllString(taxnumber, "")

	if len(taxnumber) != 14 {
		return false
	}

	sum := 0
	weight := 5
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(taxnumber[i]))
		sum += digit * weight
		weight--
		if weight == 1 {
			weight = 9
		}
	}

	rest := sum % 11
	verifyDigit1 := 11 - rest
	if verifyDigit1 == 10 || verifyDigit1 == 11 {
		verifyDigit1 = 0
	}
	if verifyDigit1 != int(taxnumber[12]-'0') {
		return false
	}

	sum = 0
	weight = 6
	for i := 0; i < 13; i++ {
		digit, _ := strconv.Atoi(string(taxnumber[i]))
		sum += digit * weight
		weight--
		if weight == 1 {
			weight = 9
		}
	}
	rest = sum % 11
	verifyDigit2 := 11 - rest
	if verifyDigit2 == 10 || verifyDigit2 == 11 {
		verifyDigit2 = 0
	}
	if verifyDigit2 != int(taxnumber[13]-'0') {
		return false
	}

	return true

}

func isValidCPF(taxnumber string) bool {
	taxnumber = regexp.MustCompile("\\D").ReplaceAllString(taxnumber, "")

	if len(taxnumber) != 11 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(taxnumber[i]))
		sum += digit * (10 - i)
	}
	rest := sum % 11
	verifyDigit1 := 11 - rest
	if verifyDigit1 == 10 || verifyDigit1 == 11 {
		verifyDigit1 = 0
	}
	if verifyDigit1 != int(taxnumber[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(taxnumber[i]))
		sum += digit * (11 - i)
	}
	rest = sum % 11
	verifyDigit2 := 11 - rest
	if verifyDigit2 == 10 || verifyDigit2 == 11 {
		verifyDigit2 = 0
	}
	if verifyDigit2 != int(taxnumber[10]-'0') {
		return false
	}

	return true
}

func isValidEmail(email string) bool {
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return regexpEmail.MatchString(email)
}

func (u *User) MapUserToResponse() *Response {
	return &Response{
		Model: gorm.Model{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		},
		FullName:     u.FullName,
		TaxNumber:    u.TaxNumber,
		Email:        u.Email,
		IsShopkeeper: u.IsShopkeeper,
		Wallet:       u.Wallet,
	}
}
