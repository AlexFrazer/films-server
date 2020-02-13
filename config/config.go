package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "sqlite3",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "frazer",
			Password: "frazer!",
			Name:     "frazer",
			Charset:  "utf8",
		},
	}
}
