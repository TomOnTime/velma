package db

import "github.com/TomOnTime/velma/models"

type Db interface {
	GetPasswordHash(user string) string
	GetAllLocations() []models.Location
	UpdateLocation(*models.Location) error
}
