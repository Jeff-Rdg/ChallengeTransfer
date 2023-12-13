package user

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(request Request) (int, error) {
	newUser, err := NewUser(request)
	if err != nil {
		return 0, err
	}

	res, err := s.repo.FindByTaxNumberOrEmail(request.TaxNumber, request.Email)
	if err != nil {
		return res, err
	}

	res, err = s.repo.Create(newUser)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetById(id uint) (*Response, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user.MapUserToResponse(), nil
}

func (s *Service) AddMoney(id uint, value float64) error {
	user, err := s.repo.GetById(id)
	if err != nil {
		return err
	}
	if user.IsShopkeeper {
		return InvalidAddBalanceErr
	}
	user.Balance += value
	_, err = s.repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

// todo: implementar serviço para realizar transferência
