// main/wire.go

//go:build wireinject
// +build wireinject

package main

import (
	"file-storage/API/handlers"
	"file-storage/API/routes"
	"file-storage/Application/commands"
	"file-storage/Application/queries"
	"file-storage/Domain/repository"
	"file-storage/Infrastructure/config"
	"file-storage/Infrastructure/db"
	"file-storage/Infrastructure/queue"

	"github.com/google/wire"
)

func InitializeApplication() (*routes.Router, error) {
	wire.Build(
		config.WireSet,
		db.WireSet,
		queue.WireSet,
		repository.WireSet,
		commands.WireSet,
		queries.WireSet,
		handlers.NewFileHandler,
		routes.NewRouter,
	)
	return &routes.Router{}, nil
}
