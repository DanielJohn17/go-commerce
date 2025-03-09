package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: GetEnv("PUBLIC_HOST", "http://localhost"),
		Port:       GetEnv("PORT", "8080"),
		DBUser:     GetEnv("DB_USER", "root"),
		DBPassword: GetEnv("DB_PASSWORD", "mypassword"),
		DBAddress:  fmt.Sprintf("%s%s", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", ":3306")),
		DBName:     GetEnv("DB_Name", "go_commerce"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
