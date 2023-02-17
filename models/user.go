package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Username string `json:"username" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Fullname string `json:"fullname" gorm:"type:text"`
	Role     string `json:"role" gorm:"type:text"`
}

type UserRespon struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (UserRespon) TableName() string {
	return "users"
}
