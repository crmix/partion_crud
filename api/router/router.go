package router

import (
	"school/api/controller"

	"github.com/gorilla/mux"
)

func Init (r *mux.Router, api controller.PupilAPI) {

	r.HandleFunc("/pupils", api.GetPupilsName).Methods("GET")
}