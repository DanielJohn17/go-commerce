package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             GetEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   GetEnv("PORT", "8080"),
		DBUser:                 GetEnv("DB_USER", "root"),
		DBPassword:             GetEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s%s", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", ":3306")),
		DBName:                 GetEnv("DB_Name", "go_commerce"),
		JWTExpirationInSeconds: GetEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              GetEnv("JWT_SECRET", "not-secret"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
