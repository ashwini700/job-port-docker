package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Jobs     []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyId"`
}
type NewCompany struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}

