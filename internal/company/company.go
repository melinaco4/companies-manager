package company

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingCompany = errors.New("failed to fetch company by id")
	ErrNotImplemented  = errors.New("not implemented")
	ErrUpdatingCompany = errors.New("could not update company")
	ErrNoCompanyFound  = errors.New("no company found")
	ErrDeletingCompany = errors.New("could not delete company")
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

// the interface we need the company storage to implement
type CompanyStore interface {
	GetCompany(context.Context, string) (Company, error)
	CreateCompany(context.Context, Company) (Company, error)
	UpdateCompany(context.Context, string, Company) (Company, error)
	DeleteCompany(context.Context, string) error
	Ping(context.Context) error
}

type Service struct {
	Store CompanyStore
}

// returns a pointer to a new service
func NewService(store CompanyStore) *Service {
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

func (s *Service) CreateCompany(ctx context.Context, cmmpn Company) (Company, error) {
	cmmpn, err := s.Store.CreateCompany(ctx, cmmpn)
	if err != nil {
		//log.Errorf("an error occurred adding the company: %s", err.Error())
		fmt.Println("an error occurred adding the company")
	}
	return cmmpn, nil
}

func (s *Service) UpdateCompany(ctx context.Context, ID string, newCompany Company,
) (Company, error) {
	cmt, err := s.Store.UpdateCompany(ctx, ID, newCompany)
	if err != nil {
		//log.Errorf("an error occurred updating the company: %s", err.Error())
		fmt.Println("an error occurred updating the company")
	}
	return cmt, nil
}

func (s *Service) DeleteCompany(ctx context.Context, ID string) error {
	return s.Store.DeleteCompany(ctx, ID)
}

func (s *Service) ReadyCheck(ctx context.Context) error {
	//log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}
