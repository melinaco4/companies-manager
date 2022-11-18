package company

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingCompany = errors.New("failed to fetch company by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// representation of the company structure
type Company struct {
	ID                string
	Name              string
	Description       string
	AmountofEmployees string
	Registered        bool
	Type              string
}

type Store interface {
	GetCompany(context.Context, string) (Company, error)
}

type Service struct {
	Store Store
}

// returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetCompany(ctx context.Context, id string) (Company, error) {

	fmt.Println("Retrieving Company...")
	cmpn, err := s.Store.GetCompany(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Company{}, ErrFetchingCompany
	}

	return cmpn, nil
}

func (s *Service) UpdateCompany(ctx context.Context, cmmpn Company) error {
	return ErrNotImplemented
}

func (s *Service) DeleteCompany(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateCompany(ctx context.Context, cmpn Company) (Company, error) {
	return Company{}, ErrNotImplemented
}
