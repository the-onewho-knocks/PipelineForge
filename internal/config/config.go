package config

import (
	"log"
	"os"
)

type Config struct {
	// RabbitMQ
	RabbitURL string
	QueueName string

	// Postgres
	PostgresDSN string
}

func Load() *Config {
	cfg := &Config{
		RabbitURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		QueueName: getEnv("RABBITMQ_QUEUE", "trending.repos"),
		PostgresDSN: getEnv(
			"POSTGRES_DSN",
			"postgres://postgres:root@localhost:5432/pipelineforge?sslmode=disable",
		),
	}

	log.Println("Loaded POSTGRES_DSN:", cfg.PostgresDSN)

	validate(cfg)
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func validate(cfg *Config) {
	if cfg.RabbitURL == "" {
		log.Fatal("RABBITMQ_URL is required")
	}
	if cfg.QueueName == "" {
		log.Fatal("RABBITMQ_QUEUE is required")
	}
	if cfg.PostgresDSN == "" {
		log.Fatal("POSTGRES_DSN is required")
	}
}
