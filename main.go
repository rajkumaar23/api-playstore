package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rollbar/rollbar-go"
)

var ctx context.Context
var rdb *redis.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	if os.Getenv("ROLLBAR_TOKEN") != "" {
		rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
		rollbar.SetEnvironment(os.Getenv("GIN_MODE"))
		rollbar.SetCodeVersion(getCurrentGitHeadHash())
		defer rollbar.Close()
	}

	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})
	defer rdb.Close()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		rollbar.Critical(fmt.Sprintf("redis connection was refused; addr = %s", rdb.Options().Addr))
		panic("redis connection failed")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	router.GET("/", getREADME)
	router.GET("/json", getAllData)
	router.GET("/:key", getDataByKey)
	
	err = endless.ListenAndServe(fmt.Sprintf("localhost:%s", os.Getenv("SERVER_PORT")), router)
	if err != nil {
		rollbar.Critical(fmt.Sprintf("http server failed to start - %v\n", err.Error()))
		panic("http server failed to start")
	}
}
