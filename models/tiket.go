package models

import "time"

type Tiket struct {
	ID     int       `json:"id" gorm:"primary_key:auto_increment"`
	Jadwal time.Time `json:"jadwal"`
	//========================
	TrainID int   `json:"train_id"`
	Train   Train `json:"train" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//========================
	StasiunAsalID  int       `json:"stasiun_asal"`
	StasiunAsal    Stasiun   `json:"stasiunasal" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WaktuBerangkat time.Time `json:"waktu_berangkat"`
	//========================
	StasiunTujuanID int       `json:"stasiun_tujuan"`
	StasiunTujuan   Stasiun   `json:"stasiuntujuan" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WaktuTiba       time.Time `json:"waktu_tiba"`
	//========================
	Harga int    `json:"harga" gorm:"type:int"`
	Stock int    `json:"stock" gorm:"type:int"`
	Kode  string `json:"kode" gorm:"type:text"`
}
