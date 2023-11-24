package service

import (
	"context"
	"errors"
	"job-port-api/internal/models"
	"job-port-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestApplicantsFilter(t *testing.T) {
	type args struct {
		ctx           context.Context
		applicantList []models.ApplicantsRequest
	}
	tests := []struct {
		name             string
		mockRepoResponse func() ([]models.ApplicantsResponse, error)
		args             args
		want             []models.ApplicantsResponse
		wantErr          bool
	}{
		{
			name: "error in db",
			args: args{
				ctx: context.Background(),
				applicantList: []models.ApplicantsRequest{
					{
						JobId:           1,
						JobRole:         "Software Engineer",
						Experience:      3,
						Salary:          80000,
						MinNotice:       30,
						MaxNotice:       60,
						Budget:          100000,
						JobLocations:    []uint{14},
						TechnologyStack: []uint{14},
						Description:     "Looking for a skilled software engineer",
						MinExp:          2,
						MaxExp:          5,
						Qualification:   []uint{10},
					},
					// Add more applicants as needed
				},
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.ApplicantsResponse, error) {
				return nil, errors.New("test error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new instance of Service
			s := &Service{}

			// Create a mock controller
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ApplicantsFilter(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			// Set the mock repository in the Service instance
			s.UserRepo = mockRepo

			// Call the method under test
			got, err := s.ApplicantsFilter(tt.args.ctx, tt.args.applicantList)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplicantsFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplicantsFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
