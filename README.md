# Payment Processor

A high-performance payment processing service built in Go, designed for the [Rinha de Backend 2025](https://github.com/zanfranceschi/rinha-de-backend-2025/blob/main/README.md).  
The system handles payment requests, balance management, and transaction summaries, optimized for concurrency and scalability.

## Features

- **Payment Creation** — Handles debit and credit operations.
- **Balance Management** — Maintains an in-memory and persistent balance store.
- **Transaction Summary** — Retrieves payment statistics for a given date range.
- **MongoDB Integration** — Stores transactions with indexes for fast queries.
- **Redis Integration** — Optional cache layer for performance optimization.
- **Retry Pattern** — Uses NATS to retry failed payment processing automatically.

## Tech Stack

- **Language**: Go
- **Database**: MongoDB
- **Cache**: Redis
- **Messaging**: NATS
- **Containerization**: Docker & Docker Compose

## Project Structure

```bash
.
├── cmd/
│   ├── api/           # REST API entrypoint
│   ├── worker/        # Worker that processes payments
├── configs/           # JSON config files for consumers and services
├── internal/
│   ├── entities/      # Domain entities
│   ├── repositories/  # MongoDB & Redis repositories
│   ├── services/      # Business logic
│   ├── usecases/      # Application use cases
│   ├── queue/         # NATS queue integration
├── test/              # Unit and integration tests
└── docker-compose.yml

## Installation & Setup

Follow the steps below to run the project locally:

### 1. Clone the repository
```bash
git clone <REPOSITORY_URL>
cd payment-processor

### 2. Build the application
docker compose up --build

The application will be available at:
http://localhost:9999

## Available Endpoints:

curl --request POST \
  --url http://localhost:9999/payments \
  --header 'Content-Type: application/json' \
  --data '{
    "correlationId": string,
    "amount": float
  }'

curl --request GET \
  --url http://localhost:9999/payments-summary