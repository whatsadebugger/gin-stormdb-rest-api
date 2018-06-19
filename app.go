package main

import (
	"github.com/asdine/storm"
)

type Application struct {
	Database *storm.DB
}

func NewApplication() *Application {
	db := createDatabase()
	return &Application{Database: db}
}
