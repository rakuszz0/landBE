package traindto

type AddTrain struct {
	Name  string `json:"name" form:"name" gorm:"type:text"`
	Kelas string `json:"kelas" form:"kelas" gorm:"type:text"`
}

type UpdateTrain struct {
	Name  string `json:"name" form:"name" gorm:"type:text"`
	Kelas string `json:"kelas" form:"kelas" gorm:"type:text"`
}
