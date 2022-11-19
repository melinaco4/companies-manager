package company

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingCompany = errors.New("failed to fetch Company by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Company Represention
// structure for the service
type Company struct {
	ID                string
	Name              string
	Description       string
	AmountofEmployees int
	Registered        bool
	Type              string
}

// Store - this interface defines all of the methods
// that the service needs in order to operate
type Store interface {
	GetCompany(context.Context, string) (Company, error)
	PostCompany(context.Context, Company) (Company, error)
	DeleteCompany(context.Context, string) error
	UpdateCompany(context.Context, string, Company) (Company, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetCompany(ctx context.Context, id string) (Company, error) {
	fmt.Printf("Getting a Company with id: %s", id)

	cmt, err := s.Store.GetCompany(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Company{}, ErrFetchingCompany
	}

	return cmt, nil
}

func (s *Service) UpdateCompany(
	ctx context.Context,
	ID string,
	updatedCmpn Company,
) (Company, error) {
	cmpn, err := s.Store.UpdateCompany(ctx, ID, updatedCmpn)
	if err != nil {
		fmt.Println("error updating Company")
		return Company{}, err
	}
	fmt.Printf("Updating Company with id: %s", ID)
	return cmpn, nil
}

func (s *Service) DeleteCompany(ctx context.Context, id string) error {
	fmt.Printf("Deleting Company with id: %s", id)
	return s.Store.DeleteCompany(ctx, id)
}

func (s *Service) PostCompany(ctx context.Context, cmpn Company) (Company, error) {
	insertedCmpn, err := s.Store.PostCompany(ctx, cmpn)
	if err != nil {
		return Company{}, err
	}
	fmt.Println("Creating a Company")
	return insertedCmpn, nil
}
