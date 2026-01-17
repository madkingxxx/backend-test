package inmem

import (
	skinport "github.com/madkingxxx/backend-test/internal/driven/inmem/skinport/adapter"
)

type Config struct {
	Skinport *skinport.Inmem
}

func New() *Config {
	return &Config{
		Skinport: skinport.New(),
	}
}
