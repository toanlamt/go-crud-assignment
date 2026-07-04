package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "goapp",
		DBPassword: "123456",
		DBName:     "productdb",
		DBSSLMode:  "disable",
	}
}