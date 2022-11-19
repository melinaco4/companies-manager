package company

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
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
	gorm.Model
	ID                string
	Name              string
	Description       string
	AmountofEmployees string
	Registered        bool
	Type              string
}

type CompanyService interface {
	GetCompany(ID uint) (Company, error)
	CreateCompany(company Company) (Company, error)
	UpdateCompany(ID uint, newCompany Company) (Company, error)
	DeleteCompany(ID uint) error
	//	Ping(context.Context) error
}

type Service struct {
	DB *gorm.DB
}

// returns a pointer to a new service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetCompany(ID uint) (Company, error) {
	var cmpn Company
	fmt.Println("Retrieving Company...")
	if result := s.DB.First(&cmpn, ID); result.Error != nil {
		return Company{}, result.Error
	}
	return cmpn, nil
}

func (s *Service) CreateCompany(company Company) (Company, error) {
	if result := s.DB.Save(&company); result.Error != nil {
		return Company{}, result.Error
	}
	return company, nil
}

func (s *Service) UpdateCompany(ID uint, newCompany Company) (Company, error) {
	company, err := s.GetCompany(ID)
	if err != nil {
		return Company{}, err
	}

	if result := s.DB.Model(&company).Updates(newCompany); result.Error != nil {
		return Company{}, result.Error
	}

	return company, nil
}

func (s *Service) DeleteCompany(ID uint) error {
	if result := s.DB.Delete(&Company{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

/*
func (s *Service) ReadyCheck(ctx context.Context) error {
	//log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}

*/
