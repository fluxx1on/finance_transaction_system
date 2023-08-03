package config

import (
	"log"
	"net/url"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type DB struct {
	Hostname string `yaml:"hostname" env-default:"localhost"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:""`
	NameDB   string `yaml:"namedb" env-default:"db"`
}

func GetDB() *DB {
	pathDB := os.Getenv("DB_PATH")
	if pathDB == "" {
		log.Fatal("DB_PATH is not set")
	}

	if _, err := os.Stat(pathDB); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", pathDB)
	}

	var db DB

	if err := cleanenv.ReadConfig(pathDB, &db); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &db
}

func GetURI() string {
	db := GetDB()
	pg := url.URL{
		Scheme: "postgres",
		Host:   db.Hostname,
		User:   url.UserPassword(db.User, db.Password),
		Path:   db.NameDB,
	}
	query := pg.Query()
	query.Add("sslmode", "disable")

	pg.RawQuery = query.Encode()

	return pg.String()
}
