package wallet

import (
	"errors"
)

var (
	InvalidUserIdErr = errors.New("user_id already has wallet")
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(request Request) (int, error) {
	// verificar se o userId ja possui carteira
	haveWallet, _ := s.repo.GetWalletByUserId(request.UserID)
	if haveWallet != nil {
		return 0, InvalidUserIdErr
	}

	newWallet, err := NewWallet(request)
	if err != nil {
		return 0, err
	}

	res, err := s.repo.Create(newWallet)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetById(id uint) (*Wallet, error) {
	w, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return w, nil
}
