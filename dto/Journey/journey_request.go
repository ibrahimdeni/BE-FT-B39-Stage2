package journeydto

import "time"

type JourneyRequest struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	UserID      int    `json:"user_id" form:"user_id" gorm:"int"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description" form:"description" gorm:"type: text"`
	Message      string `json:"message" gorm:"type:text"`
}

type UpdateJourneyRequest struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	UserID      int    `json:"user_id" form:"user_id" gorm:"int"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Description string `json:"description" form:"description" gorm:"type: text"`
	Message      string `json:"message" gorm:"type:text"`
}
