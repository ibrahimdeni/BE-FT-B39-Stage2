package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" from:"email"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
	Msg      string `gorm:"type: varchar(255)" json:"msg"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
}
