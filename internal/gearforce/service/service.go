package service

import (
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/google/uuid"
)

type Hello interface {
	Hello(s string) (string, error)
}

type SaveRoster interface {
	SaveRoster(r models.Roster) (uuid.UUID, error)
}

type GetRoster interface {
	GetRoster(uuid.UUID) (models.Roster, error)
}
