package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xtracdev/rollreg/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/developers", handlers.DevsQueryHandler)
	r.HandleFunc("/v1/developers/{email}", handlers.GetDevHandler).Methods("GET")
	r.HandleFunc("/v1/developers/{email}", handlers.PutDevHandler).Methods("PUT")

	r.HandleFunc("/v1/applications", handlers.AppsHandler).Methods("GET")
	r.HandleFunc("/v1/applications", handlers.PostAppsHandler).Methods("POST")
	r.HandleFunc("/v1/applications/{client_id}", handlers.GetApplicationByIdHandler).Methods("GET")
	r.HandleFunc("/v1/applications/{client_id}", handlers.PutApplicationByIdHandler).Methods("PUT")

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Warn(err.Error())
	}
}
