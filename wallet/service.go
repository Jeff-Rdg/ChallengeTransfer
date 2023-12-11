package wallet

import (
	"ChallengeBackEndPP/user"
	"errors"
)

var (
	InvalidUserIdErr = errors.New("user_id already has wallet")
)

type Service struct {
	repo     Repository
	userRepo user.Repository
}

func NewService(r Repository, ur user.Repository) *Service {
	return &Service{repo: r, userRepo: ur}
}

func (s *Service) Create(request Request) (int, error) {
	_, err := s.userRepo.GetById(request.UserID)
	if err != nil {
		return 0, err
	}

	// verificar se o userId ja possui carteira
	haveWallet, _ := s.repo.GetWalletByUserId(request.UserID)
	if haveWallet != nil {
		return 0, InvalidUserIdErr
	}

	newWaller, err := NewWallet(request)
	if err != nil {
		return 0, err
	}

	res, err := s.repo.Create(newWaller)
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
