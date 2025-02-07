

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

### Sprint 2: Implement API Gateway and Service Registry

##### Running Services at local: 

| Services  | Port |
| ------------- | ------------- |
| Claim Service  | :8080  |
| Payment Service  | :8081  |
| Policy Service | :8082  |
| Claim DB/Cache  | :5432 / :6379  |
| Payment DB/Cache  | :5433 / :6380  |
| Policy DB/Cache  | :5434 / :6381  |

#### Result after adding API Gateway 
![alt text](assets/api_gateway.png)

- How the service registry work between multiple services
![alt text](assets/registry.png)

### Sprint 3: Decoupling problem 

#### Why Do We Need Decoupling Between Two Services?
Decoupling services allows systems to be **scalable, resilient, and maintainable** by reducing dependencies between components. In a **tightly coupled system**, one service directly depends on another, making failures or changes more disruptive. By introducing **decoupling**, services can **communicate asynchronously, handle failures gracefully, and scale independently**.

#### Key Benefits of Decoupling
- **Scalability**: Each service can scale independently based on its workload.
- **Fault Tolerance**: If one service fails, it does not bring down the entire system.
- **Maintainability**: Easier to modify, replace, or update services without affecting others.
- **Asynchronous Processing**: Improves performance by allowing tasks to be queued and processed later.
- **Better Resource Utilization**: Services can process tasks at their own speed, avoiding bottlenecks.

#### Example: Payment and Claim
Consider a scenario where 
- **Claim Service** (Handles customer claim)
- **Payment Service** (Processes payments)

#### Without Decoupling (Synchronous Request-Response)
- The **Claim Service** calls the **Payment Service** immediately.
- If the **Payment Service** is down, claim fail.
- The **Claim Service** is blocked, leading to poor performance and possibly data lost.

#### With Decoupling (Using a Message Queue)
- The **Claim Service** places a **payment request** in a **message queue**.
- The **Payment Service** picks up the request **asynchronously** and processes it.
- If the **Payment Service** is down, the request remains in the queue and is retried later.

**Result**: Claim are **not blocked**, failures are **handled**, and the system **scales better**.
![alt text](assets/message.png)