package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"meta-aggregator/internal/pkg/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IAPI interface {
	setupRoute(rg *gin.RouterGroup)
	setupRestrictedRoute(rg *gin.RouterGroup)
}

// API v1
func AddRouterV1(server *gin.Engine, config interface{}) {
	routerV1 := server.Group("/api/v1")
	addApi(newHealthApi(), "/health", routerV1)
}

func AddRestrictedRouterV1(server *gin.Engine, config interface{}) {

}

func addApi(api IAPI, path string, rg *gin.RouterGroup) {
	apiRg := rg.Group(path)
	api.setupRoute(apiRg)
}

func addRestrictedApi(api IAPI, path string, rg *gin.RouterGroup) {
	apiRg := rg.Group(path)
	api.setupRestrictedRoute(apiRg)
}

func getConfig(cfg interface{}) (api.RedisConfig, *gorm.DB) {
	c, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("Cannot marshal config")
	}

	var field api.Config
	err = json.Unmarshal(c, &field)
	if err != nil {
		log.Fatalf("Domain config is not properly config")
	}

	return field.Redis, getDB(field.DBConfig)
}

func getDB(config api.DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(config.DNS()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(config.LogLevel)),
		},
	)
	if err != nil {
		log.Fatalf("failed to open database connection, err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.db, err: %v", err)
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnLifeTime) * time.Second)

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database, err: %v", err)

	}

	fmt.Println("Successfully connect to MySQL database")

	return db
}

func GetRedisConnectionOption(redisAddresses []string, pwd string, masterName string) (asynq.RedisConnOpt, error) {
	if len(redisAddresses) == 0 {
		return nil, errors.New("redis host is empty")
	}
	if masterName != "" {
		return asynq.RedisFailoverClientOpt{
			SentinelAddrs: redisAddresses,
			MasterName:    masterName,
		}, nil

	}
	if len(redisAddresses) == 1 {
		return asynq.RedisClientOpt{
			Addr:     redisAddresses[0],
			Password: pwd,
		}, nil
	}
	return asynq.RedisClusterClientOpt{
		Addrs:    redisAddresses,
		Password: pwd,
	}, nil
}
