package Config

import (
	utils "PaymentService/Utils"
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func RegisterServiceWithConsul(config *Appconfig) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://localhost:8500" // Consul address

	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Error creating Consul client: %v", err)
	}

	serviceID := "payment-service-" + config.Server.GinPort

	reg := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "payment-service",
		Port:    utils.StringToInt(config.Server.GinPort),
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			HTTP: fmt.Sprintf("http://host.docker.internal:%s/health", config.Server.GinPort),
			Interval: "10s",
			Timeout: "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		log.Fatalf("Error registering service with Consul: %v", err)
	}

	fmt.Println("Service registered successfully with Consul!")
}