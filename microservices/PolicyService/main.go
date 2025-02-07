package main

import (
	Config "PolicyService/AppConfig"
	application "PolicyService/Application"
	infrastructure "PolicyService/Infrastructure"
	interfaces "PolicyService/Interface"
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

	Config.SetEnvironment(env)
	config, err := Config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_data, err := Config.Connect(config)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db_redis, err  := Config.ConnectRedis(config)
	if err != nil {
		log.Fatal("Error connecting to redis")
	}
	
	repo := infrastructure.NewPolicyRepository(db_data, db_redis)
	service := application.NewPolicyService(repo)
	handler := interfaces.NewPolicyHandler(service)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	
	router.POST("/policys", handler.CreatePolicy)
	router.GET("/policys", handler.GetAllPolicy)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "WRONG API PATH"})
	})

	// * Register service in Consul
	Config.RegisterServiceWithConsul(config)

	if err := router.Run(":" + config.Server.GinPort); err != nil {
		log.Fatal("FAILED TO START SERVER", err)
	}
}