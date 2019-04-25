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
	key, isExist := os.LookupEnv("PVP")
	if !isExist {
		log.Fatal("env `PVP` is absent")
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PVP PONG, key = "+key)
	})

	go func() {
		router.Run(":8081")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
}
