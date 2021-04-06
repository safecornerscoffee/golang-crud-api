package config

import (
	"fmt"
	"os"
)

type Config struct {
	Environment      string
	ConnectionString string
}

func getConfigValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}

func getConnectionString() (pqInfo string) {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")
	pqInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDb)
	fmt.Println(pqInfo)
	return pqInfo
}
func NewConfig() *Config {

	return &Config{
		Environment:      getConfigValue("ENV", "local"),
		ConnectionString: getConnectionString(),
	}
}
