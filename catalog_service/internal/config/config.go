package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	//AppPort  string
	Postgres PostgresConfig
	//Redis    RedisConfig
	//Kafka    KafkaConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host string
	Port string
}

type KafkaConfig struct {
	Brokers []string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := &Config{
		//AppPort: os.Getenv("APP_PORT"),

		Postgres: PostgresConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},

		//Redis: RedisConfig{
		//	Host: os.Getenv("REDIS_HOST"),
		//	Port: os.Getenv("REDIS_PORT"),
		//},
		//Kafka: KafkaConfig{
		//	Brokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		//},
	}
	return cfg, nil
}
