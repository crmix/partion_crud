package main

import (
	"log"
	"net/http"
	"school/api/controller"
	"school/api/router"
	"school/config"
	"school/platform"
	"school/service"
	"school/storage"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)


func main() {
	db := platform.DBConn()
	cfg := config.Load()

    storage :=storage.NewPupilRepo(db)
	service :=service.NewPupilService(storage, *zap.NewStdLog(zap.L()))
	api := controller.NewPupilAPI(service)

	root := mux.NewRouter()

	router.Init(root, api)

	httpServer := http.Server{
		Addr: cfg.HTTPPort,
		Handler: root,
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("shut down server")
	}

}