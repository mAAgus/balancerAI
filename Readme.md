# 🤖 AI Balancer

**Simple AI load balancer** that routes requests to different LLM models depending on request type.

---

## 🗂 Supported Types
- 💬 Chat
- 🖼️ Image generation
- 🔊 Audio generation

The service distributes requests between multiple models using **round-robin balancing** with **fallback support**.

---

## 🏗 Architecture


client
│
▼
REST API (Fiber)
│
▼
Service layer
│
▼
AI Balancer
│
├── Chat Models 💬
├── Image Models 🖼️
└── Audio Models 🔊


The balancer routes requests to appropriate model groups.

---

## ✨ Features
- ✔ REST API
- ✔ Round-robin model balancing
- ✔ Model fallback on failure
- ✔ Configuration via TOML + ENV
- ✔ Observability stack:
  - 📊 Prometheus metrics
  - 📝 Loki logs
  - 🕵️ Jaeger tracing
  - 📈 Grafana dashboards
- ✔ Unit tests
- ✔ Integration tests
- ✔ Docker multi-stage build
- ✔ CI with GitHub Actions

---

## 🛠 Tech Stack
- Go
- Fiber
- Prometheus
- Grafana
- Loki
- Jaeger
- Docker
- OpenTelemetry

---

## ⚙️ Configuration

Configuration is loaded from:


configs/config.toml


Environment variables override config values.

Example:

```bash
AI_BALANCER_SERVER_BIND_ADDR=:9080
```
### Example Config
```
[server]
bind_addr=":9080"

[logger]
level="debug"

[jaeger]
url="http://jaeger:14268/api/traces"

[models.chat]
endpoint="http://chat-model:8080/generate"
api_key="test-key"

[models.image]
endpoint="http://image-model:8080/generate"
api_key="test-key"

[models.audio]
endpoint="http://audio-model:8080/generate"
api_key="test-key"

```

## 🧩 API

### Generate
**POST** `/api/v1/generate`

**Example request:**

```json
{
  "type": "chat",
  "payload": {
     "message": "hello"
  }
}
```
### Example response:
```json
{
  "result": "hello world"
}
```
### Health Check

GET /api/v1/health

## Metrics

GET /metrics – Prometheus endpoint
### 👀 Observability

The project includes full observability stack:

---
- Tool	Purpose
- Prometheus	Metrics
- Grafana	Dashboards
- Loki	Logs
- Jaeger	Tracing


Access URLs:
---
- Grafana: http://localhost:3000

- Prometheus: http://localhost:9090

- Jaeger: http://localhost:16686

---

## 📊 Metrics

Available metrics:
---
- ai_balancer_requests_total

- ai_balancer_requests_errors_total

- ai_balancer_requests_duration_second

### Labels:

type = chat | image | audio

---
## 🚀 Running the Project

### Start Full Stack
- task docker

or

- docker compose up --build
## Development

### Run locally:

- task run

### Format code:

- task format

### Run linter:

- task lint

### Run unit tests:

- task test

### Run integration tests:

- task test-integration

## ⚙️ CI

GitHub Actions pipeline runs on Pull Requests.

Pipeline stages:

---

- lint

- format

- build

- unit tests

- integration tests

---

Integration tests use docker-compose environment.

## 🧪 Testing

### Unit Tests

- go test ./...

### Integration Tests

- docker compose -f docker-compose-test.yml up

## 📝 Git Commit Convention


Project follows Conventional Commits:

---

- feat: add audio model support

- fix: handle model fallback error

- test: add balancer concurrency test

- refactor: simplify service logic

---

## 🔮 Future Improvements

### Possible improvements:

---

- Weighted round-robin

- Circuit breaker for models

- Model health checks

- Rate limiting

- Request queue

- Dynamic model discovery

### Also 
A client_for_test is provided here
to verify the operation of balancerAI when interacting with external models and to check tracing functionality in Jaeger.
## 👤 Author

AI Balancer test project
