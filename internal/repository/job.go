package repository

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	"job-port-api/internal/models"
)

func (r *Repo) Fetchjob(ctx context.Context, cid uint64) (models.Job, error) {
	var jobData models.Job
	result := r.DB.Where("id = ?", cid).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not find the job id")
	}
	return jobData, nil

}

func (r *Repo) FetchJobPosts(ctx context.Context) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find jobs")
	}
	return jobDetails, nil

}

func (r *Repo) FetchJobByCompanyId(ctx context.Context, cid uint64) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Where("company_id = ?", cid).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job for the cid")
	}
	return jobDetails, nil

}

func (r *Repo) AddJob(jobData models.Job) (models.Job, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create the job")
	}
	return jobData, nil
}
