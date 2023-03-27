package api

import (
	"fmt"
)

type DatabaseConfig struct {
	Host         string
	Port         int
	DBName       string
	User         string
	Password     string
	ConnLifeTime int
	ConnTimeOut  int
	MaxIdleConns int
	MaxOpenConns int
	LogLevel     int
}

type Config struct {
	Redis    RedisConfig
	DBConfig DatabaseConfig
}

type RedisConfig struct {
	RedisAddresses string
	Password       string
	MasterName     string
}

func (c *DatabaseConfig) DNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&timeout=%ds",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.ConnTimeOut,
	)
}
