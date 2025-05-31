package config

type Config struct {
	DBService DBService
}

type DBService struct {
	URL string
}

func New() *Config {
	return &Config{DBService: DBService{URL: "http://db-service:25432"}}
}
