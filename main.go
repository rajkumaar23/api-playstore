package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
