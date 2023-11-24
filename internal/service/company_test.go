package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"job-port-api/internal/auth"
	"job-port-api/internal/models"
	"job-port-api/internal/repository"
)

func TestService_AddCompany(t *testing.T) {
	type args struct {
		ctx         context.Context
		companyData models.Company
	}
	tests := []struct {
		name             string
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		// TODO: Add test cases.
		{
			name: "error in db",
			args: args{
				ctx: context.Background(),
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("test error")
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				companyData: models.Company{
					Name:     "Microsoft",
					Location: "bangalore",
				},
			},
			want: models.Company{
				Name:     "Microsoft",
				Location: "bangalore",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Microsoft",
					Location: "bangalore",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		mc := gomock.NewController(t)
		mockRepo := repository.NewMockUserRepo(mc)
		if tt.mockRepoResponse != nil {
			mockRepo.EXPECT().CreateCompany(tt.args.ctx, tt.args.companyData).Return(tt.mockRepoResponse()).AnyTimes()
		}
		s, _ := NewService(mockRepo, &auth.Auth{})
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AddCompany(tt.args.ctx, tt.args.companyData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FetchCompanyDetails(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
			name: "error in db",
			args: args{
				ctx: context.Background(),
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("test error")
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: models.Company{
				Name:     "Teksystem",
				Location: "bangalore",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Teksystem",
					Location: "bangalore",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().FetchAllCompanies().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})

			got, err := s.FetchCompByid(tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FetchAllCompanies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		args             args
		want             []models.Company
		wantErr          bool
		mockRepoResponse func() ([]models.Company, error)
	}{
		{
			name: "error in db",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Company, error) {
				return nil, errors.New("test error")
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Company{
				{
					Name:     "Teksystem",
					Location: "bangalore",
				},
				{
					Name:     "IBM",
					Location: "bangalore",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Company, error) {
				return []models.Company{
					{
						Name:     "Teksystem",
						Location: "bangalore",
					},
					{
						Name:     "IBM",
						Location: "bangalore",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().FetchAllCompanies().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})

			got, err := s.FetchAllCompanies()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}
