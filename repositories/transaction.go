package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TransRepository interface {
	FindAllTransaction() ([]models.Transaction, error)
	FindTransUser(userId int) ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	TransPending(id int) ([]models.Transaction, error)
	CancelTrans(ID int) (models.Transaction, error)
	UpdatePayment(status string, ID int) error
}

func RepositoryTrans(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransUser(userId int) ([]models.Transaction, error) {
	var trans []models.Transaction
	err := r.db.Preload("Tiket").Preload("Tiket.Train").Preload("Tiket.StasiunAsal").Preload("Tiket.StasiunTujuan").Preload("User").Where("user_id =?", userId).Order("id DESC").Find(&trans).Error

	return trans, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Tiket").Preload("Tiket.Train").Preload("Tiket.StasiunAsal").Preload("Tiket.StasiunTujuan").Preload("User").First(&transaction, ID).Error

	return transaction, err
}
func (r *repository) TransPending(id int) ([]models.Transaction, error) {
	var trans []models.Transaction
	err := r.db.Where("status =? and user_id=?", "pending", id).Find(&trans).Error

	return trans, err
}
func (r *repository) CancelTrans(ID int) (models.Transaction, error) {
	var trans models.Transaction
	r.db.First(&trans, ID)

	trans.Status = "cancel"
	err := r.db.Save(&trans).Error
	return trans, err
}

func (r *repository) FindAllTransaction() ([]models.Transaction, error) {
	var trans []models.Transaction
	err := r.db.Preload("Tiket").Preload("Tiket.Train").Preload("Tiket.StasiunAsal").Preload("Tiket.StasiunTujuan").Preload("User").Find(&trans).Error
	return trans, err
}
func (r *repository) UpdatePayment(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("Tiket").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var tiket models.Tiket
		r.db.First(&tiket, transaction.TiketID)
		tiket.Stock = tiket.Stock - transaction.Qty
		r.db.Save(&tiket)
	}

	transaction.Status = status
	err := r.db.Debug().Save(&transaction).Error
	return err
}
