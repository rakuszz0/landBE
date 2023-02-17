package repositories

import (
	"landtick/models"
	"time"

	"gorm.io/gorm"
)

type TiketRepository interface {
	CreateTiket(tiket models.Tiket) (models.Tiket, error)
	FilterTiket(asal int, tujuan int, jadwal time.Time) ([]models.Tiket, error)
	FindTiket() ([]models.Tiket, error)
	GetTiket(ID int) (models.Tiket, error)
	Deletetiket(tiket models.Tiket) (models.Tiket, error)
	//========STASIUN===============
	FilterKotaStasiun(kota string) (models.Stasiun, error)
	// ======TRANSACTION============
	CreateTrans(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTiket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTiket(tiket models.Tiket) (models.Tiket, error) {
	err := r.db.Create(&tiket).Error
	return tiket, err
}
func (r *repository) FindTiket() ([]models.Tiket, error) {
	var tiket []models.Tiket
	err := r.db.Preload("Train").Preload("StasiunAsal").Preload("StasiunTujuan").Order("id DESC").Find(&tiket).Error

	return tiket, err
}
func (r *repository) FilterTiket(asal int, tujuan int, jadwal time.Time) ([]models.Tiket, error) {
	var tiket []models.Tiket
	err := r.db.Debug().Preload("Train").Preload("StasiunAsal").Preload("StasiunTujuan").Where("stasiun_asal_id = ? AND stasiun_tujuan_id = ? AND jadwal = ?", asal, tujuan, jadwal).Find(&tiket).Error

	return tiket, err
}
func (r *repository) GetTiket(ID int) (models.Tiket, error) {
	var tiket models.Tiket
	err := r.db.First(&tiket, ID).Error

	return tiket, err
}
func (r *repository) Deletetiket(tiket models.Tiket) (models.Tiket, error) {
	err := r.db.Delete(&tiket).Error
	return tiket, err
}

// =============STASIUN==============
func (r *repository) FilterKotaStasiun(kota string) (models.Stasiun, error) {
	var stasiun models.Stasiun
	err := r.db.Where("kota = ?", kota).First(&stasiun).Error
	return stasiun, err
}

// ==========TRANSACTION=============
func (r *repository) CreateTrans(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Debug().Create(&transaction).Error
	return transaction, err
}

func (r *repository) HistoryTransTiket(user int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Tiket").Preload("Tiket.Train").Preload("User").Where("status =? AND user_id = ?", "pending", user).Find(&transaction).Error
	return transaction, err
}
