package models

type Stasiun struct {
	ID   int    `json:"ID" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type:text"`
	Kota string `json:"kota" gorm:"type:text"`
}
