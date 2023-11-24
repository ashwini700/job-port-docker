package repository

import (
	"job-port-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) ApplicantsFilter(jobId uint) (*models.Job, error) {
	var jobData models.Job
	err := r.DB.Preload("JobLocations").Preload("TechnologyStack").Preload("Qualification").Where("ID = ?", jobId).Find(&jobData).Error
	if err != nil {
		log.Error().Err(err).Msg("Problem in fetching joba data")
		return nil, err
	}
	return &jobData, nil
}