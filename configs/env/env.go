package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	JWTTokenName           string
	JWTExpirationInSeconds int
	JWTSecret              string
	GoogleOAuthURL         string
	DBHost                 string
	DBPort                 int
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBSSLMode              string
}

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valAsInt
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		Port:                   GetString("PORT", ":3000"),
		JWTTokenName:           GetString("JWT_TOKEN_NAME", "token"),
		JWTExpirationInSeconds: GetInt("JWT_EXP", 3600),
		JWTSecret:              GetString("JWT_SECRET", ""),
		GoogleOAuthURL:         GetString("GOOGLE_OAUTH_URL", "https://www.googleapis.com/oauth2/v3/userinfo"),
		DBHost:                 GetString("DB_HOST", "localhost"),
		DBPort:                 GetInt("DB_HOST", 5432),
		DBUser:                 GetString("DB_USER", "root"),
		DBPassword:             GetString("DB_PASSWORD", ""),
		DBName:                 GetString("DB_NAME", "db"),
		DBSSLMode:              GetString("DB_SSLMODE", "disable"),
	}
}

var Envs = initConfig()
