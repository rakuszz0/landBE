package authdto

type LoginResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Role     string `gorm:"type: text" json:"role"`
}
type RegistResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Fullname string `gorm:"type: varchar(255)" json:"Fullname"`
	Role     string `gorm:"type: text" json:"role"`
}
type AuthResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Role     string `gorm:"type: text" json:"role"`
}
