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

func TestService_UserSignup(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.UserSignup
	}
	tests := []struct {
		name             string
		args             args
		want             models.User
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		// TODO: Add test cases.
		{
			name: "error from the database",
			args: args{
				ctx: context.Background(),
				userData: models.UserSignup{
					Name:     "ashwini",
					Email:    "ash@gmail.com",
					Password: "123",
				},
			},
			want:    models.User{}, // Change the expected result to an empty User since an error is expected.
			wantErr: true,          // Set wantErr to true since an error is expected.
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("error while hashing the password")
			},
		},
		{
			name: "success from the database",
			args: args{
				ctx: context.Background(),
				userData: models.UserSignup{
					Name:     "ashwini",
					Email:    "ash@gmail.com",
					Password: "123",
				},
			},
			want: models.User{
				Name:         "ashwini",
				Email:        "ash@gmail.com",
				PasswordHash: "hashed password",
			}, // Change the expected result to an empty User since an error is expected.
			wantErr: false, // Set wantErr to true since an error is expected.
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Name:         "ashwini",
					Email:        "ash@gmail.com",
					PasswordHash: "hashed password",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error in initializing the repo layer")
				return
			}
			got, err := s.UserSignup(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestService_UserLogin(t *testing.T) {
// 	type args struct {
// 		ctx      context.Context
// 		userData models.UserLogin
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             string
// 		wantErr          bool
// 		claims           jwt.RegisteredClaims
// 		mockResponse     func() (models.User, error)
// 		mockAuthResponse func() (string, error)
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "wrong email",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserLogin{
// 					Email:    "ashw@gmail.com",
// 					Password: "123",
// 				},
// 			},
// 			want:    "",
// 			wantErr: true,
// 			mockResponse: func() (models.User, error) {
// 				return models.User{}, errors.New("test error from the mock function")
// 			},
// 			mockAuthResponse: func() (string, error) {
// 				return "", errors.New("test error from the mock function")
// 			},
// 		},
// 		{
// 			name: "token generation failed",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserLogin{
// 					Email:    "ash@gmail.com",
// 					Password: "123",
// 				},
// 			},
// 			want:    "jwt test string",
// 			wantErr: false,
// 			mockResponse: func() (models.User, error) {
// 				return models.User{
// 					Email:        "ash@gmail.com",
// 					PasswordHash: "$2a$10$uS/GmX48bxvhGPS.IrujaefuktoqGuKz3HBeOOMH6MGrnDT1H4TEy",
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 				}, nil
// 			},
// 			claims: jwt.RegisteredClaims{
// 				Issuer:  "job portal project",
// 				Subject: "1",
// 				Audience: jwt.ClaimStrings{
// 					"users",
// 				},
// 				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
// 				IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			},
// 			mockAuthResponse: func() (string, error) {
// 				return "jwt test string", nil
// 			},
// 		},
// 		{
// 			name: "success generate token",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserLogin{
// 					Email:    "ash@gmail.com",
// 					Password: "123",
// 				},
// 			},
// 			want:    "",
// 			wantErr: true,
// 			mockResponse: func() (models.User, error) {
// 				return models.User{
// 					Email:        "ash@gmail.com",
// 					PasswordHash: "$2a$10$uS/GmX48bxvhGPS.IrujaefuktoqGuKz3HBeOOMH6MGrnDT1H4TEy",
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 				}, nil
// 			},
// 			claims: jwt.RegisteredClaims{
// 				Issuer:  "job portal project",
// 				Subject: "1",
// 				Audience: jwt.ClaimStrings{
// 					"users",
// 				},
// 				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
// 				IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			},
// 			mockAuthResponse: func() (string, error) {
// 				return "", errors.New("test error from mock function")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			mockAuth := .NewMockAuthentication(mc)

// 			mockRepo.EXPECT().CheckEmail(tt.args.ctx, tt.args.userData.Email).Return(tt.mockResponse()).AnyTimes()

// 			mockAuth.EXPECT().GenerateAuthToken(tt.claims).Return(tt.mockAuthResponse()).AnyTimes()

// 			s, err := NewService(mockRepo, mockAuth)
// 			if err != nil {
// 				t.Errorf("error is initializing the repo layer")
// 				return
// 			}
// 			got, err := s.UserLogin(tt.args.ctx, tt.args.userData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Service.UserLogin() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
