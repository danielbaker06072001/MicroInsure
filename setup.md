# API Gateway Setup (Kong) - Demonstration Purpose Only

> Please refer to the official documentation of [Consul](https://www.consul.io/docs) and [Kong](https://docs.konghq.com/) for better setup. This setup is for demonstration purposes only.

## 📌 API Gateway Paths (After Setup)

| **Service Name**  | **API Gateway Path (Kong)** | **Backend Service Path** |
|-------------------|----------------------------|--------------------------|
| Claim Service    | `http://localhost:8000/claims` | `http://localhost:8080` |
| Payment Service  | `http://localhost:8000/payments` | `http://localhost:8081` |
| Policy Service   | `http://localhost:8000/policies` | `http://localhost:8082` |

## 🛠️ **Testing the Services**
After setup, you can test the API Gateway using `curl` or Postman:

## How to setup Consul 
docker run -d --name=consul \
  -p 8500:8500 \
  -p 8600:8600/udp \
  hashicorp/consul "agent" "-dev" "-client=0.0.0.0"


- Link to check services discovery: http://localhost:8500/v1/agent/services


## How to setup Kong API Gateway
```bash
docker network create kong-net

docker run -d --name kong-database \
  --network=kong-net \
  -p 5432:5432 \
  -e POSTGRES_USER=kong \
  -e POSTGRES_DB=kong \
  -e POSTGRES_PASSWORD=kong \
  postgres:latest

docker run -d --name kong \
  --network=kong-net \
  -e KONG_DATABASE=postgres \
  -e KONG_PG_HOST=kong-database \
  -e KONG_ADMIN_ACCESS_LOG=/dev/stdout \
  -e KONG_ADMIN_ERROR_LOG=/dev/stderr \
  -e KONG_PROXY_ACCESS_LOG=/dev/stdout \
  -e KONG_PROXY_ERROR_LOG=/dev/stderr \
  -e KONG_ADMIN_LISTEN=0.0.0.0:8001 \
  -e KONG_PROXY_LISTEN=0.0.0.0:8000 \
  -p 8000:8000 \
  -p 8001:8001 \
  kong:latestdo k

docker run --rm --network=kong-net -e KONG_DATABASE=postgres -e KONG_PG_HOST=host.docker.internal -e KONG_PG_PORT=5432 -e KONG_PG_USER=postgres -e KONG_PG_PASSWORD=Cel-365. -e KONG_PG_DATABASE=postgres kong:latest kong migrations bootstrap; docker stop kong; docker rm kong; docker run -d --name kong --network=kong-net -e KONG_DATABASE=postgres -e KONG_PG_HOST=host.docker.internal -e KONG_PG_PORT=5432 -e KONG_PG_USER=postgres -e KONG_PG_PASSWORD=Cel-365. -e KONG_PG_DATABASE=postgres -e KONG_ADMIN_ACCESS_LOG=/dev/stdout -e KONG_ADMIN_ERROR_LOG=/dev/stderr -e KONG_PROXY_ACCESS_LOG=/dev/stdout -e KONG_PROXY_ERROR_LOG=/dev/stderr -e KONG_ADMIN_LISTEN=0.0.0.0:8001 -e KONG_PROXY_LISTEN=0.0.0.0:8000 -p 8000:8000 -p 8001:8001 kong:latest
```


# Delete existing claim-service
```bash
1. Invoke-WebRequest -Uri "http://localhost:8001/services/claim-service" -Method Delete
2. Invoke-WebRequest -Uri "http://localhost:8001/services/payment-service" -Method Delete
3. Invoke-WebRequest -Uri "http://localhost:8001/services/policy-service" -Method Delete
```

# Create Service
```bash
1. Invoke-WebRequest -Uri "http://localhost:8001/services/" `
  -Method Post `
  -Body "name=claim-service&url=http://host.docker.internal:8080" `
  -ContentType "application/x-www-form-urlencoded"

2. Invoke-WebRequest -Uri "http://localhost:8001/services/" `
  -Method Post `
  -Body "name=payment-service&url=http://host.docker.internal:8081" `
  -ContentType "application/x-www-form-urlencoded"

3. Invoke-WebRequest -Uri "http://localhost:8001/services/" `
  -Method Post `
  -Body "name=policy-service&url=http://host.docker.internal:8082" `
  -ContentType "application/x-www-form-urlencoded"
```

# Create route for existing Services
```bash
1. Invoke-WebRequest -Uri "http://localhost:8001/services/claim-service/routes" `
  -Method Post `
  -Body "name=claim-route&paths[]=/claims" `
  -ContentType "application/x-www-form-urlencoded"

2. Invoke-WebRequest -Uri "http://localhost:8001/services/payment-service/routes" `
  -Method Post `
  -Body "name=payment-route&paths[]=/payments" `
  -ContentType "application/x-www-form-urlencoded"

3. Invoke-WebRequest -Uri "http://localhost:8001/services/policy-service/routes" `
  -Method Post `
  -Body "name=policy-route&paths[]=/policies" `
  -ContentType "application/x-www-form-urlencoded"
```
