package api

import "github.com/jinzhu/gorm"

type dbRepository struct {
	db *gorm.DB
}

func NewDBResponsitory(db *gorm.DB) *dbRepository {
	return &dbRepository{db: db}
}
