package service

import (
	"fmt"
	"log"
	"school/models"
	"school/storage"
)

type pupilServiceImpl struct {
	storage storage.Pupil1
	//logger  log.Logger
}

func NewPupilService(pupil1 storage.Pupil1, logger log.Logger) PupilService {
	return &pupilServiceImpl{
		storage: pupil1,
		//logger: logger,
	}
}

type PupilService interface {
	GetPupilsName(page, limit int) ([]models.School, error)
}


func (c *pupilServiceImpl) GetPupilsName(page, limit int) ([]models.School, error) {
	res, err :=c.storage.GetPupilsName(page, limit)
	if err !=nil {
		log.Println(err)
		return nil, fmt.Errorf("error in getpupilname %w", err) 
	}

	return res, nil

}