package repository

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	"job-port-api/internal/models"
)

// CreateCompany creates a new company record in the database
func (r *Repo) CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error) {
	result := r.DB.Create(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not create the company")
	}
	return companyData, nil
}

// FetchCompany retrieves a company record by its unique ID.
// func (r *Repo) FetchCompany(ctx context.Context, cid uint64) (models.Company, error) {
// 	var companyData models.Company
// 	result := r.DB.Where("id = ?", cid).First(&companyData)
// 	if result.Error != nil {
// 		log.Info().Err(result.Error).Send()
// 		return models.Company{}, errors.New("Not find the company")
// 	}
// 	return companyData, nil
// }
func (r *Repo) FetchAllCompanies() ([]models.Company, error) {
	var companyDetails []models.Company
	result := r.DB.Find(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("Not find Any companies")
	}
	return companyDetails, nil
}
func (r *Repo) FetchCompByid(cid uint64) (models.Company, error) {
	var companyData models.Company
	result := r.DB.Where("id=?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("company data is not there with that id")
	}
	return companyData, nil
}
