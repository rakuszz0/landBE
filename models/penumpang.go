package models

type Penumpang struct {
	NIK    int    `json:"nik" gorm:"primary_key"`
	Name   string `json:"name" gorm:"type:text"`
	Alamat string `json:"alamat" gorm:"type:text"`
	Phone  int    `json:"phone" gorm:"type:varchar(100)"`
	Photo  string `json:"photo" gorm:"type:text"`
	UserID int    `json:"user_id"`
	User   User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
