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

// model of the company, will be implemented in our service
// trying to accept interfaces and return structs
type Company struct {
	ID                string
	Name              string
	Description       string
	AmountofEmployees int
	Registered        bool
	Type              string
}

// this is where all the methods of the service are stored
type Store interface {
	GetCompany(context.Context, string) (Company, error)
	PostCompany(context.Context, Company) (Company, error)
	DeleteCompany(context.Context, string) error
	UpdateCompany(context.Context, string, Company) (Company, error)
}

// Service is the main struct
type Service struct {
	Store Store
}

// NewService - returns a pointer of a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// retrieves a company from the database or returns error
func (s *Service) GetCompany(ctx context.Context, id string) (Company, error) {
	fmt.Printf("Getting a Company with id: %s\n", id)

	cmpn, err := s.Store.GetCompany(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Company{}, ErrFetchingCompany
	}

	return cmpn, nil
}

// updating by a company in the database or returns an error
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
	fmt.Printf("Updating Company with id: %s\n", ID)
	return cmpn, nil
}

// deleting company by id in the database
func (s *Service) DeleteCompany(ctx context.Context, id string) error {
	fmt.Printf("Deleting Company with id: %s\n", id)
	return s.Store.DeleteCompany(ctx, id)
}

// creating new company record in the database
func (s *Service) PostCompany(ctx context.Context, cmpn Company) (Company, error) {
	insertedCmpn, err := s.Store.PostCompany(ctx, cmpn)
	if err != nil {
		return Company{}, err
	}
	fmt.Println("Creating a Company")
	return insertedCmpn, nil
}
