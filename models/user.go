package models

import "time"

// User model struct
type User struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: text"`
	Msg string `json:"msg" gorm:"type:text"`
	// JourneyID int `json:"-"`
	Journey []JourneyInUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"journeys"`
	Bookmark []BookmarkResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"bookmarks"`
	Role      string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersProfileResponse struct {
	ID			int		`json:"id"`
	Fullname	string	`json:"fullname"`
	Email		string	`json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Msg string `json:"msg"`
	CreatedAt time.Time `json:"created_at"`
}


func (UsersProfileResponse) TableName() string {
	return "users"
}