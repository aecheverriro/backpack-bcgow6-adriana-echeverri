package users

type Service interface {
	GetAll(string, string, string, string, int, float64, bool, string) ([]User, error)
	GetId(string) (User, error)
	CreateUser(string, string, string, string, int, float64, bool, string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(id string, name string, lastName string, email string, age int, height float64, active bool, creation string) ([]User, error) {
	us, err := s.repository.GetAll(id, name, lastName, email, age, height, active, creation)
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) GetId(id string) (User, error) {
	us, err := s.repository.GetId(id)
	if err != nil {
		return User{}, err
	}

	return us, nil
}

func (s *service) CreateUser(id string, name string, lastName string, email string, age int, height float64, active bool, creation string) (User, error) {
	us, err := s.repository.CreateUser(id, name, lastName, email, age, height, active, creation)
	if err != nil {
		return User{}, err
	}

	return us, nil
}
