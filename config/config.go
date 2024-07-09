package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type AppConfig struct {
	AppPort string
}
type JwtConfig struct {
	Key string
	Durasi time.Duration
	Issues string
}

type Config struct {
	DbConfig
	AppConfig
	JwtConfig
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.readConfig(); err != nil {
		return nil, err
	}
	return config, nil
}


func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	lifeTime, err := strconv.Atoi(os.Getenv("JWT_LIFE_TIME"))
	if err != nil {
		return err
	}
	c.JwtConfig = JwtConfig{
		Key:    os.Getenv("JWT_KEY"),
		Durasi: time.Duration(lifeTime),
		Issues: os.Getenv("JWT_ISSUER_NAME"),
	}

	c.DbConfig = DbConfig{
		Host :     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		}

		c.AppConfig = AppConfig{
			AppPort: os.Getenv("SERVER_PORT"),
		}
	return nil
}