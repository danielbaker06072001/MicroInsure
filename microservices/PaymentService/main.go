package main

import (
	Config "PaymentService/AppConfig"
	application "PaymentService/Application"
	infrastructure "PaymentService/Infrastructure"
	interfaces "PaymentService/Interface"
	utils "PaymentService/Utils"
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
	db_mq, err := Config.ConnectRabbitMQ(config)
	if err != nil {
		log.Fatal("Error connecting to message broker")
	}
	
	repo := infrastructure.NewPaymentRepository(db_data, db_redis, db_mq)
	service := application.NewPaymentService(repo)
	handler := interfaces.NewPaymentHandler(service)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	
	router.POST("/payments", handler.CreatePayment)
	router.GET("/payments", handler.GetAllPayment)

	// * Listen to queue from Claims to process payment 
	err = utils.ListenToQueue(db_mq, "claim_validation", handler.ProcessPayment)
	if err != nil {
		log.Fatal("Error receiving message from queue")
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "WRONG API PATH"})
	})

	// * Register service in Consul
	Config.RegisterServiceWithConsul(config)

	if err := router.Run(":" + config.Server.GinPort); err != nil {
		log.Fatal("FAILED TO START SERVER", err)
	}
}