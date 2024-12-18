package cqrs

import (
	postgres_db "blog-api/Infrastructure/databases/postgres/config"

	"github.com/google/wire"
)

func ProvideUnitOfWork(postgresDatabase *postgres_db.PostgresDatabase) UnitOfWork {
	return NewUnitOfWork(postgresDatabase.DB)
}

var WireSet = wire.NewSet(
	ProvideUnitOfWork,
)
