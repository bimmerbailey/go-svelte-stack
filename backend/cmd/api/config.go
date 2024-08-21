package main

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	hostname string
	port     string
	password string
	name     string
	username string
}

func (config *DatabaseConfig) DatabaseURL() string {
	dbEnv := os.Getenv("DATABASE_URL")
	if len(dbEnv) < 0 {
		return fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			config.username,
			config.password,
			config.hostname,
			config.port,
			config.name,
		)
	}
	return dbEnv
}

func DatabaseFromEnvs() *DatabaseConfig {
	dbHostname := os.Getenv("DATABASE_HOSTNAME")
	if len(dbHostname) == 0 {
		dbHostname = "mongo"
	}
	dbPort := os.Getenv("DATABASE_PORT")
	if len(dbPort) == 0 {
		dbPort = "27017"
	}
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	if len(dbPassword) == 0 {
		dbPassword = "password"
	}
	dbName := os.Getenv("DATABASE_NAME")
	if len(dbName) == 0 {
		dbName = "your_name"
	}
	dbUsername := os.Getenv("DATABASE_USERNAME")
	if len(dbUsername) == 0 {
		dbUsername = "app_user"
	}
	return &DatabaseConfig{
		hostname: dbHostname,
		port:     dbPort,
		password: dbPassword,
		name:     dbName,
		username: dbUsername,
	}
}
