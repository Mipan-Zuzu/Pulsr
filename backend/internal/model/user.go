package model

import (
	"gorm.io/gorm"
)


type AccsesTokenSupabase struct {
	gorm.Model
	Authorization string `validate:"required"`
}