package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Status    string    `json:"status" gorm:"type:text"`
	TiketID   int       `json:"tiket_id"`
	Tiket     Tiket     `json:"tiket" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty       int       `json:"qty" gorm:"type:int"`
	Total     int       `json:"total" gorm:"type:int"`
	DateTrans time.Time `json:"datetrans"`
}
