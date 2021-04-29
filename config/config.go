package config

import (
	"fmt"
	"os"
)

type Config struct {
	dbUser     string
	dbPass     string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
}

func Get() *Config {
	conf := &Config{
		dbUser:     os.Getenv("POSTGRES_USER"),
		dbPass:     "password",
		dbHost:     os.Getenv("POSTGRES_HOST"),
		dbPort:     os.Getenv("POSTGRES_PORT"),
		dbName:     os.Getenv("POSTGRES_DB"),
		testDBHost: os.Getenv("TEST_DB_HOST"),
		testDBName: os.Getenv("TEST_DB_NAME"),
	}

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBHost, c.testDBName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPass,
		dbhost,
		c.dbPort,
		dbname,
	)
}
