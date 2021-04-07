package config

import (
	"fmt"
	"os"
)

type Config struct {
	ConnectionString string
}

func getConfigValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}

func getConnectionString() (connectionString string) {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")
	connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDb)
	return connectionString
}
func NewConfig() *Config {

	return &Config{
		ConnectionString: getConnectionString(),
	}
}
