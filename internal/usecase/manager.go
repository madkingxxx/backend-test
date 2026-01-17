package usecase

import (
	"github.com/madkingxxx/backend-test/internal/service"
	skinport "github.com/madkingxxx/backend-test/internal/usecase/skinport"
	user "github.com/madkingxxx/backend-test/internal/usecase/user"
)

type Manager struct {
	Skinport *skinport.UseCase
	User     *user.UseCase
}

func New(serviceManager *service.Manager) *Manager {
	return &Manager{
		Skinport: skinport.New(serviceManager.Skinport),
		User:     user.New(serviceManager.User, serviceManager.Skinport),
	}
}
