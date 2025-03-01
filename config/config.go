package config

import (
	"os"
)

type Config struct {
	PublicIP string
	dbHost   string
}

func LoadConfig() Config {
	return Config{
		PublicIP: os.Getenv("PUBLIC_IP"),
		dbHost:   os.Getenv("DB_HOST"),
	}
}
