package storage

import (
	"database/sql"
	"school/models"
	"school/storage/postgres"
)

type Pupil1 interface {
	GetPupilsName(page, limit int) ([]models.School, error)
}

func NewPupilRepo(db *sql.DB) Pupil1 {
	return postgres.NewPupilRepo(db)
}