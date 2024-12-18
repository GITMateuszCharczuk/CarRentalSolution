package config

import (
	"blog-api/Infrastructure/config"

	"github.com/google/wire"
)

func NewPostgresConfigProvider(cfg *config.Config) *PostgresDatabase {
	return NewPostgresConfig(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDBName, cfg.PostgresHost, cfg.PostgresPort, cfg.RunPostgresMigration)
}

var WireSet = wire.NewSet(NewPostgresConfigProvider)
