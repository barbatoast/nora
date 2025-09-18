package config

import (
	"os"
	"strings"
	"fmt"
)

type Config struct {
	Port       string
	JWTSecret  string
	DSN        string
	CORSOrigins []string
}

func Env() Config {
	port := getenv("PORT", "8080")
	secret := getenv("JWT_SECRET", "dev-secret")
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			getenv("PGHOST","localhost"), getenv("PGPORT","5432"), getenv("PGUSER","postgres"), getenv("PGPASSWORD","postgres"), getenv("PGDATABASE","vapi"))
	}
	cors := strings.Split(getenv("CORS_ORIGINS", "*"), ",")
	return Config{Port: port, JWTSecret: secret, DSN: dsn, CORSOrigins: cors}
}

func getenv(k, def string) string { if v := os.Getenv(k); v != "" { return v }; return def }
