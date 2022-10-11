package repositories

import (
	"thejourney/models"

	"gorm.io/gorm"
)

type JourneyRepository interface {
	FindJourneys() ([]models.Journey, error)
	GetJourney(ID int) (models.Journey, error)
	CreateJourney(Journey models.Journey) (models.Journey, error)
	UpdateJourney(Journey models.Journey) (models.Journey, error)
	DeleteJourney(Journey models.Journey) (models.Journey, error)
}

func RepositoryJourney(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindJourneys() ([]models.Journey, error) {
	var journeys []models.Journey
	err := r.db.Preload("User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Find(&journeys).Error // add this code

	return journeys, err
}

func (r *repository) GetJourney(ID int) (models.Journey, error) {
	var journey models.Journey
	err := r.db.Preload("User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").First(&journey, ID).Error // add this code

	return journey, err
}

func (r *repository) CreateJourney(journey models.Journey) (models.Journey, error) {
	err := r.db.Preload("User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Create(&journey).Error

	return journey, err
}

func (r *repository) UpdateJourney(journey models.Journey) (models.Journey, error) {
	err := r.db.Preload("User").Save(&journey).Error

	return journey, err
}

func (r *repository) DeleteJourney(journey models.Journey) (models.Journey, error) {

	err := r.db.Preload("User").Delete(&journey).Error

	return journey, err
}
