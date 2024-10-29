// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"file-storage/API/routes"
	"file-storage/Application/commands"
	"file-storage/Application/handlers"
	"file-storage/Application/queries"
	"file-storage/Infrastructure/config"
	"file-storage/Infrastructure/db"
	"file-storage/Infrastructure/publisher"
	"file-storage/Infrastructure/queue"
	"file-storage/Infrastructure/repository"

	"github.com/google/wire"
)

func InitializeApplication() (*routes.Router, error) {
	wire.Build(
		config.WireSet,
		db.WireSet,
		queue.WireSet,
		publisher.WireSet,
		repository.WireSet,
		commands.WireSet,
		queries.WireSet,
		handlers.WireSet,
		routes.NewRouter,
	)
	return &routes.Router{}, nil
}
