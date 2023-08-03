package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/exp/slog"
)

// Logger
type Logger struct {
	Logfile   string     `yaml:"logfile"`
	LevelInfo slog.Level `yaml:"levelInfo"`
}

// RabbitClient configuration settings for RabbitMQ
type RabbitClient struct {
	Address                 string `yaml:"address"`
	User                    string `yaml:"user"`
	Password                string `yaml:"password"`
	WorkerByChannelAmount   int    `yaml:"workerByChannelAmount"`
	QueueAmount             int    `yaml:"queueAmount"`
	TransactionExchangeName string `yaml:"transactionExchangeName"`
	RoutingKey              string `yaml:"routingKey"`
}

// Config is a configuration struct that store enviromental variables
type Config struct {
	ServerAddress    string        `yaml:"serverAddress"`
	ListenerProtocol string        `yaml:"listenerProtocol"`
	PostgreSQL       string        `yaml:"postgreSQL"`
	Logger           *Logger       `yaml:"logger"`
	RabbitMQ         *RabbitClient `yaml:"rabbitMQ"`
}

func Setup() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	cfg.PostgreSQL = GetURI()

	return &cfg
}
