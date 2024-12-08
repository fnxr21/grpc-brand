package repository

import (
	"github.com/fnxr21/brand/internal/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func Repositorybrand(db *gorm.DB) *repository {
	return &repository{db}
}

type BrandService interface {
	Create(brand entities.Brand) (entities.Brand, error)
}

// Create implements BrandService.
func (r *repository) Create(brand entities.Brand) (entities.Brand, error) {
	err := r.db.Create(&brand).Error
	return brand, err
}
