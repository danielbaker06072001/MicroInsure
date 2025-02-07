

# Microservices Project and Kubernetes : MicroInsure

## Project Introduction: 

- MicroInsure is a Golang based microservices Insurance portal application, featuring a scalable distributed and high performance concurrentcy along with CI/CD implementation. Each services will have its own database system, along with Kafka or RabbitMQ for reliable message queue. 

## Tech used 
- Framework and Programming Language: Golang
- Tool and Testing: PostgreSQL, RedisCache, RabbitMQ, Kafka, API Gateway, K6 for load testing

### Sprint 1: Design Monolithic system and improved to Microservices - discuss their differences

#### Monolithic System - Code base under /monolithic
- A single unified codebase.
- All components are interconnected and interdependent.
- Easier to develop initially but harder to scale.
- Deployment involves the entire system.
- If one service have high load then other services will be impacted as well.

![alt text](assets/monolithic_design.png)

#### Microservices System - Code base under /microservices
- Composed of small, independent services.
- Each service can be developed, deployed, and scaled independently.
- More complex to develop initially but easier to scale and maintain.
- Deployment involves individual services.
- In case one services fail, the other services will not be down.
- If any service experience high load, other services will not be affected, meanwhile, we can set different resources for each services based on usage.

![alt text](assets/microservice_design.png)

#### Differences
- **Scalability**: Monolithic systems are harder to scale, while microservices can be scaled independently.
- **Development**: Monolithic systems are easier to develop initially, whereas microservices require more effort.
- **Deployment**: Monolithic systems require full deployment, while microservices allow for independent deployment.
- **Maintenance**: Microservices are easier to maintain due to their modular nature, while monolithic systems can become complex over time.

##### Running Services at local: 

| Services  | Port |
| ------------- | ------------- |
| Claim Service  | :8080  |
| Payment Service  | :8081  |
| Policy Service | :8082  |
| Claim DB/Cache  | :5432 / :6379  |
| Payment DB/Cache  | :5433 / :6380  |
| Policy DB/Cache  | :5434 / :6381  |