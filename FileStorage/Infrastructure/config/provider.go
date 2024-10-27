package config

import "github.com/google/wire"

func ProvideConfig() *Config {
	cfg, err := NewConfig("../../.env")
	if err != nil {
		panic("Failed to load configuration")
	}
	return cfg
}

var WireSet = wire.NewSet(ProvideConfig)
