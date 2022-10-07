package config

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Db *sql.DB
}

func (c *Config) initDb() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "P@ssw0rd"
	dbName := "db_enigma"
	dbDriver := "postgres"

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
