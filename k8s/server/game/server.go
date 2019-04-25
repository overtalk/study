package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	key, isExist := os.LookupEnv("GAME")
	if !isExist {
		log.Fatal("env `GAME` is absent")
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "GAME PONG, key = "+key)
	})

	go func() {
		router.Run(":8080")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
}
