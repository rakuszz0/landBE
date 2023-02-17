package stasiundto

type AddStasiun struct {
	Name string `json:"name" form:"name" gorm:"type:text"`
	Kota string `json:"kota" form:"kota" gorm:"type:text"`
}
type UpdateStasiun struct {
	Name string `json:"name" form:"name" gorm:"type:text"`
	Kota string `json:"kota" form:"kota" gorm:"type:text"`
}
