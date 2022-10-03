package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" from:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Msg      string `gorm:"type: varchar(255)" json:"msg"`
}