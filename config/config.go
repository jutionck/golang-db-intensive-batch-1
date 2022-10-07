package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Db *sql.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err)
	}
	c.Db = db
}

func (c *Config) DbConn() *sql.DB {
	return c.Db
}

func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
