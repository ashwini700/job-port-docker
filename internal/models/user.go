package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type UserSignup struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Applicants struct {
	gorm.Model
	JobId           uint   `json:"jobid"`
	JobRole         string `json:"job_role"`
	NoticePeriod	uint	`json:"notice"`
	Experience      uint   `json:"experience"`
	Salary          uint   `json:"salary"`
	MinNotice       uint   `json:"minnotice"`
	MaxNotice       uint   `json:"maxnotice"`
	Budget          uint   `json:"budget"`
	JobLocations    []uint `gorm:"many2many:job_location;"`
	TechnologyStack []uint `gorm:"many2many:tech_stack;"`
	Description     string `json:"desc"`
	MinExp          uint   `json:"min_exp"`
	MaxExp          uint   `json:"max_exp"`
	Qualification   []uint `gorm:"many2many:qualification;"`
}

type ApplicantsRequest struct {
	JobId           uint
	JobRole         string
	Experience      uint
	NoticePeriod	uint
	Salary          uint
	MinNotice       uint
	MaxNotice       uint
	Budget          uint
	JobLocations    []uint
	TechnologyStack []uint
	Description     string
	MinExp          uint
	MaxExp          uint
	Qualification   []uint
}

type ApplicantsResponse struct {
	JobRole string
	JobId   uint
}
