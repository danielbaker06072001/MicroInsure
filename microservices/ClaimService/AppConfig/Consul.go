package Config

import (
	utils "ClaimService/Utils"
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

	serviceID := "claim-service-" + config.Server.GinPort
	reg := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "claim-service",
		Port:    utils.StringToInt(config.Server.GinPort),
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			HTTP: fmt.Sprintf("http://host.docker.internal:%s/health", config.Server.GinPort),	// ! This is based on whether you are running the services in Docker or not
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

func DiscoverServiceWithConsul(service_name string) (string, error){
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://localhost:8500"

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return "", fmt.Errorf("error creating Consul client: %v", err)
	}

	// Look for the PaymentService in Consul
	service, _, err := client.Agent().Service(service_name, nil)
	if err != nil {
		return "", fmt.Errorf("error retrieving service: %v", err)
	}

	// Construct the service URL
	url := fmt.Sprintf("http://%s:%d", service.Address, service.Port)
	return url, nil
}