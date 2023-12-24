package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

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

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := srv.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpectedly")
		}
	}

	log.Println("Server exiting")
}
