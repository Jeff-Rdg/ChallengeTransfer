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
	u, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return u.MapUserToResponse(), nil
}
