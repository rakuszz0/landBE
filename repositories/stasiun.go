package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type StasiunRepository interface {
	AddStasiun(stasiun models.Stasiun) (models.Stasiun, error)
	FindAllStasiun() ([]models.Stasiun, error)
	FindStasiunID(ID int) (models.Stasiun, error)
	UpdateStasiun(stasiun models.Stasiun) (models.Stasiun, error)
	DeleteStasiun(stasiun models.Stasiun) (models.Stasiun, error)
}

func RepositoryStasiun(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddStasiun(stasiun models.Stasiun) (models.Stasiun, error) {
	err := r.db.Create(&stasiun).Error

	return stasiun, err
}
func (r *repository) FindAllStasiun() ([]models.Stasiun, error) {
	var stasiun []models.Stasiun
	err := r.db.Order("kota ASC").Find(&stasiun).Error
	return stasiun, err
}
func (r *repository) FindStasiunID(ID int) (models.Stasiun, error) {
	var stasiun models.Stasiun
	err := r.db.First(&stasiun, ID).Error
	return stasiun, err
}

func (r *repository) UpdateStasiun(stasiun models.Stasiun) (models.Stasiun, error) {
	err := r.db.Save(&stasiun).Error

	return stasiun, err
}

func (r *repository) DeleteStasiun(stasiun models.Stasiun) (models.Stasiun, error) {
	err := r.db.Delete(&stasiun).Error
	return stasiun, err
}
