package config

import (
	"log"
	"net/url"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Logger
type Logger struct {
	Logfile   string `yaml:"logfile"`
	LevelInfo string `yaml:"levelInfo"`
}

// RabbitConfig configuration settings for RabbitMQ
type RabbitConfig struct {
	URL                   string
	Host                  string `yaml:"host"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	WorkerByChannelAmount int    `yaml:"workerByChannelAmount"`
	QueueAmount           int    `yaml:"queueAmount"`
	ExchangeName          string `yaml:"exchangeName"`
	QueueName             string `yaml:"queueName"`
	RoutingKey            string `yaml:"routingKey"`
}

func (rmq *RabbitConfig) SetURI() {
	mqURL := &url.URL{
		Scheme: "amqp",
		Host:   rmq.Host,
		User:   url.UserPassword(rmq.User, rmq.Password),
	}

	rmq.URL = mqURL.String()
}

// Config is a configuration struct that store environmental variables
type Config struct {
	ServerAddress    string        `yaml:"serverAddress"`
	ListenerProtocol string        `yaml:"listenerProtocol"`
	Logger           *Logger       `yaml:"logger"`
	RabbitMQ         *RabbitConfig `yaml:"rabbitMQ"`
	PostgreSQL       *PostgresConfig
	Docker           *DockerConfig
}

func (cfg *Config) GetAlt() {
	cfg.ServerAddress = cfg.Docker.Hosts.ListenerHost

	cfg.RabbitMQ.Host = cfg.Docker.Hosts.RabbimqHost
	cfg.PostgreSQL.Host = cfg.Docker.Hosts.PostgresqlHost

	cfg.RabbitMQ.SetURI()
	cfg.PostgreSQL.SetURI()
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
		log.Fatalf("can't read config: %s", err)
	}

	cfg.RabbitMQ.SetURI()

	cfg.PostgreSQL = NewDB()

	if os.Getenv("DOCKER_PATH") != "" {
		cfg.Docker = GetDocker()
		cfg.GetAlt()
	}

	return &cfg
}
