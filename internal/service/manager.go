package service

import (
	driven "github.com/madkingxxx/backend-test/internal/driven"
	skinportservice "github.com/madkingxxx/backend-test/internal/service/skinport"
	userservice "github.com/madkingxxx/backend-test/internal/service/user"
)

type Manager struct {
	Skinport *skinportservice.Service
	User     *userservice.Service
}

func New(drivenCfg *driven.Config) *Manager {
	return &Manager{
		Skinport: skinportservice.New(drivenCfg.Inmem.Skinport, drivenCfg.ExtHTTP.Skinport),
		User:     userservice.New(drivenCfg.PostgreSQL.User),
	}
}
