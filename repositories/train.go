package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TrainRepository interface {
	AddTrain(train models.Train) (models.Train, error)
	FindAllTrain() ([]models.Train, error)
	FindTransID(ID int) (models.Train, error)
	UpdateTrain(train models.Train) (models.Train, error)
	DeleteTrain(train models.Train) (models.Train, error)
}

func RepositoryTrain(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddTrain(train models.Train) (models.Train, error) {
	err := r.db.Create(&train).Error

	return train, err
}
func (r *repository) FindAllTrain() ([]models.Train, error) {
	var train []models.Train
	err := r.db.Find(&train).Error
	return train, err
}
func (r *repository) FindTransID(ID int) (models.Train, error) {
	var train models.Train
	err := r.db.First(&train, ID).Error
	return train, err
}

func (r *repository) UpdateTrain(train models.Train) (models.Train, error) {
	err := r.db.Save(&train).Error

	return train, err
}

func (r *repository) DeleteTrain(train models.Train) (models.Train, error) {
	err := r.db.Delete(&train).Error
	return train, err
}
