package inventory

import "errors"

// TODO: this is service/business logic implementation

type service struct {
	repo Repository
}

type Service interface {
	GetAll(page int, limit int) ([]Inventory, error)
	GetByCode(code string) (Inventory, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(page int, limit int) ([]Inventory, error) {
	return s.repo.GetAll(page, limit)
}

func (s *service) GetByCode(code string) (Inventory, error) {
	return s.repo.GetByCode(code)
}

func (s *service) ProcessPayment(method string, amount float64) error {
	switch method {
	case "creditcard":
		// do something
	case "paypal":
		// do something
	default:
		return errors.New("payment cannot be processed")
	}

	return nil
}
