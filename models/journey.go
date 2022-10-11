package models

import "time"

type Journey struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	User        UsersProfileResponse `json:"user"`
	UserID      int                  `json:"user_id" gorm:"int"`
	Image       string               `json:"image" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: text"`
	CreatedAt time.Time 			 `json:"created_at"`
	UpdatedAt time.Time				 `json:"updated_at"`
	Bookmark	[]BookmarkResponse	`json:"bookmark" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message      string				 `json:"message" gorm:"type:text"`
}

type JourneyInUser struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	UserID      int                  `json:"user_id" gorm:"int"`
	User        UsersProfileResponse `json:"user"`
	Image       string               `json:"image" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: text"`
	CreatedAt time.Time 			 `json:"created_at"`
	UpdatedAt time.Time				 `json:"updated_at"`
	Message      string				 `json:"message" gorm:"type:text"`
}

func (JourneyInUser) TableName() string {
	return "journeys"
}