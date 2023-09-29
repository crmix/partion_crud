package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"school/models"
	"school/query"
)

type pupilRepo struct {
	db *sql.DB
}

func NewPupilRepo(db *sql.DB) *pupilRepo {
	return &pupilRepo{
		db: db,
	}
}

func (s *pupilRepo) GetPupilsName(page, limit int) ([]models.School, error) {
	offset := (page - 1) * limit

	row, err := s.db.Query(query.SCHOOL_SQL, offset, limit)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("errorf in GetPupilsName: %w", err)
	}
	var schools []models.School

	for row.Next() {
		var school models.School
		row.Scan(
			&school.Id,
			&school.Name)

		schools = append(schools, school)
	}
	for i, school := range schools {
		rows, err := s.db.Query(query.CLASS_SQL, school.Id)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("error in GetClassesName: %w", err)
		}
		var classes []models.Classes

		for rows.Next() {
			var class models.Classes
			rows.Scan(
				&class.Id,
				&class.Name)
			schools[i].Class = append(schools[i].Class, class)

		}

		for k, class := range schools[i].Class {
			rows, err := s.db.Query(query.PUPIL_SQL, class.Id)
			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("error in GetPupilsName: %w", err)
			}
			for rows.Next() {
				var pupil models.Pupils
				rows.Scan(
					&pupil.Id,
					&pupil.Name)
				schools[i].Class[k].Pupil = append(schools[i].Class[k].Pupil, pupil)
			}
			fmt.Println(classes)
		}
	}
	fmt.Println(schools)
	return schools, nil
}
