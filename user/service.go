package user

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(fullname, taxnumber, email, password string, isshopkeeper bool) (int, error) {
	newUser, err := NewUser(fullname, taxnumber, email, password, isshopkeeper)
	if err != nil {
		return 0, err
	}

	res, err := s.repo.FindByTaxNumberOrEmail(taxnumber, email)
	if err != nil {
		return res, err
	}

	res, err = s.repo.Create(newUser)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetById(id uint) (*User, error) {
	u, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
