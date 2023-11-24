package service

import (
	"context"
	"errors"
	"job-port-api/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

func (s *Service) ApplicantsFilter(ctx context.Context, applicantList []models.ApplicantsRequest) ([]models.ApplicantsResponse, error) {
	var response []models.ApplicantsResponse
	ch := make(chan models.ApplicantsResponse, len(applicantList))
	wg := &sync.WaitGroup{}

	for _, application := range applicantList {
		wg.Add(1)
		go func(application models.ApplicantsRequest) {
			defer wg.Done()
			jobData, err := s.UserRepo.ApplicantsFilter(application.JobId) //@ fetching job data
			if err != nil {
				log.Error().Err(err).Interface("Job ID", application.JobId).Send()
				return
			}
			if jobData.Budget < application.Budget {
				log.Error().Err(errors.New("Budget requirments not met")).Interface("applicant ID", application.JobRole).Send()
				return
			}
			if jobData.Experience <= application.Experience && application.Experience > jobData.MinExp {
				log.Error().Err(errors.New("Experience requirments not met")).Interface("applicant ID", application.JobRole).Send()
				return
			}
			if jobData.NoticePeriod <= application.NoticePeriod && application.NoticePeriod > jobData.MinNotice {
				log.Error().Err(errors.New("Notice period requirments not met")).Interface("applicant ID", application.JobRole).Send()
				return
			}
			for _, j := range jobData.Qualification {
				for _, a := range application.Qualification {
					if j.Model.ID != a {
						log.Error().Err(errors.New("Qualification requirments not met")).Interface("applicant ID", application.JobRole).Send()
						return
					}
				}
			}
			var available bool
			for _, j := range jobData.JobLocations {
				for _, a := range application.JobLocations {
					if j.Model.ID == a {
						available = true
					}
				}
			}
			if available == false {
				log.Error().Err(errors.New("Location requirments not met")).Interface("applicant ID", application.JobRole).Send()
				return
			}
			count := 0
			for _, j := range jobData.TechnologyStack {
				for _, a := range application.TechnologyStack {
					if j.Model.ID == a {
						count++
					}
				}
			}
			if count < (len(jobData.TechnologyStack) / 2) {
				log.Error().Err(errors.New("Techstack requirments not met")).Interface("applicant ID", application.JobRole).Send()
				return
			}
			respo := models.ApplicantsResponse{
				JobRole: application.JobRole,
				JobId:   application.JobId,
			}
			ch <- respo
		}(application)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		response = append(response, data)
	}
	return response, nil
}
