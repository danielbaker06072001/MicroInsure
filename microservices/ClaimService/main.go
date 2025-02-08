package main

import (
	Config "ClaimService/AppConfig"
	application "ClaimService/Application"
	infrastructure "ClaimService/Infrastructure"
	interfaces "ClaimService/Interface"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	// * Connect to database
	Config.SetEnvironment(env)
	config, err := Config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_data, err := Config.Connect(config)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db_redis, err := Config.ConnectRedis(config)
	if err != nil {
		log.Fatal("Error connecting to redis")
	}
	db_mq, err := Config.ConnectRabbitMQ(config)
	if err != nil {
		log.Fatal("Error connecting to message broker")
	}

	repo := infrastructure.NewClaimRepository(db_data, db_redis, db_mq)
	service := application.NewClaimService(repo)
	handler := interfaces.NewClaimHandler(service)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	router.POST("/claims", handler.CreateClaim)
	router.GET("/claims", handler.GetAllClaim)
	router.POST("/claims/validate", handler.ValidateClaim)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "WRONG API PATH"})
	})

	// * Register service in Consul
	Config.RegisterServiceWithConsul(config)

	// Start the Gin server
	if err := router.Run(":" + config.Server.GinPort); err != nil {
		log.Fatal("FAILED TO START SERVER", err)
	}
}