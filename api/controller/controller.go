package controller

import (
	"fmt"
	"log"
	"net/http"
	"school/service"
	"strconv"
)



type pupilApi struct {
	service service.PupilService
}

func NewPupilAPI(service service.PupilService) PupilAPI {
	return &pupilApi{
		service: service,
	}
}

type PupilAPI interface {
	GetPupilsName(w http.ResponseWriter, r *http.Request)
}

func (api *pupilApi) GetPupilsName(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, err := strconv.Atoi(q.Get("page"))
	if err != nil {
		fmt.Print(err)
	}
	limit, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
		fmt.Print(err)
	}
	
	res, err := api.service.GetPupilsName(page, limit)
	if err != nil {
		log.Println(err)
		return
	}

	WriteSuccess(w, res)
 }