package usersdto

import "time"

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender"  gorm:"type: varchar(255)" `
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: text"`
	Msg      string `json:"msg" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
}
