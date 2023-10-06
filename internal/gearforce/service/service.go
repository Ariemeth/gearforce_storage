package service

import (
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
)

type Hello interface {
	Hello(s string) (string, error)
}

type SaveRoster interface {
	Store(r models.Roster) (models.Roster, error)
}
