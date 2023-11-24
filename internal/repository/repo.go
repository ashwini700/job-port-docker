package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"job-port-api/internal/models"
)

//go:generate mockgen -source=repo.go -destination=repository_mock.go -package=repository

type Repo struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)
	//
	CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error)
	FetchCompByid(cid uint64) (models.Company, error)
	FetchAllCompanies() ([]models.Company, error)
	//
	AddJob(jobData models.Job) (models.Job, error)
	FetchJobByCompanyId(ctx context.Context, cid uint64) ([]models.Job, error)
	FetchJobPosts(ctx context.Context) ([]models.Job, error)
	Fetchjob(ctx context.Context, cid uint64) (models.Job, error)

	ApplicantsFilter(jobId uint) (*models.Job, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
