package skinport

import skinport "github.com/madkingxxx/backend-test/internal/driven/ext_http/skinport/adapter"

type Config struct {
	Skinport *skinport.Sender
}

func New(baseURL string) *Config {
	return &Config{
		Skinport: skinport.New(baseURL),
	}
}
