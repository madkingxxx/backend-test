package driver

import (
	skinporthttp "github.com/madkingxxx/backend-test/internal/driver/http/skinport/adapter"
	userhttp "github.com/madkingxxx/backend-test/internal/driver/http/user/adapter"
	"github.com/madkingxxx/backend-test/internal/usecase"
)

type Config struct {
	Skinport *skinporthttp.Controller
	User     *userhttp.Controller
}

func New(usecaseManager *usecase.Manager) *Config {
	return &Config{
		Skinport: skinporthttp.New(usecaseManager.Skinport),
		User:     userhttp.New(usecaseManager.User),
	}
}
