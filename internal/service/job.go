package service

import (
	"context"

	"job-port-api/internal/models"
)

func (s *Service) AddJob(ctx context.Context, jobData models.NewJob, cid uint64) (models.Job, error) {
	jobDetails:= models.Job{
		CompanyId:       uint64(cid),
		JobRole:         jobData.JobRole,
		Salary:          jobData.Salary,
		MinNotice:       jobData.MinNotice,
		MaxNotice:       jobData.MaxNotice,
		Budget:          jobData.Budget,
		JobLocations:    jobData.JobLocations,
		TechnologyStack: jobData.TechnologyStack,
		Description:     jobData.Description,
		MinExp:          jobData.MinExp,
		MaxExp:          jobData.MaxExp,
		Qualification:   jobData.Qualification,
	}

	jobDetails, err := s.UserRepo.AddJob(jobDetails)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJobByCompId(ctx context.Context, cid uint64) (models.Job, error) {
	jobData, err := s.UserRepo.Fetchjob(ctx, cid)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil

}

func (s *Service) FetchJobPosts(ctx context.Context) ([]models.Job, error) {
	jobData, err := s.UserRepo.FetchJobPosts(ctx)
	if err != nil {
		return nil, err
	}
	return jobData, nil

}

func (s *Service) FetchJobDetails(ctx context.Context, cid uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.FetchJobByCompanyId(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
