package service

import (
	"context"

	"job-port-api/internal/models"
)

func (s *Service) AddCompany(ctx context.Context, companyData models.Company) (models.Company, error) {
	companyData, err := s.UserRepo.CreateCompany(ctx, companyData)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
func (s *Service) FetchCompByid(cid uint64) (models.Company, error) {
	companyData, err := s.UserRepo.FetchCompByid(cid)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
func (s *Service) FetchAllCompanies() ([]models.Company, error) {
	companyDetails, err := s.UserRepo.FetchAllCompanies()
	if err != nil {
		return nil, err
	}
	return companyDetails, nil
}
