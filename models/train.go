package models

type Train struct {
	ID    int    `json:"ID" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type:text"`
	Kelas string `json:"kelas" gorm:"type:text"`
}
