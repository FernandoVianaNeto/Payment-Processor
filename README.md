# 💸 Smart Payment Router – High Performance Backend

This project is a **high-performance backend service** designed to receive payment requests and intelligently route them to one of two external payment processors. It ensures **cost efficiency**, **resilience**, and **speed**, even under unstable external conditions or heavy load.

This implementation demonstrates architecture design, fault tolerance, performance optimization, and real-world backend problem-solving — making it ideal as a portfolio project.

---

## ✨ Features

- 🚦 **Smart routing** – Selects the processor with the lowest transaction fee  
- 🔁 **Failover strategy** – Detects instability and reroutes to the available processor  
- ⚡ **High throughput** – Optimized to handle thousands of concurrent requests  
- 🔍 **Observability** – Health checks, structured logging, and metrics included  
- 🧪 **Tested and verifiable** – Automated tests ensure system correctness

---

## ⚙️ Tech Stack

- **Golang** for fast, concurrent backend logic
- **PostgreSQL** or equivalent as the transactional database
- **Docker** and **Docker Compose** for containerization
- **CI/CD** via GitHub Actions (or similar)
- **Optional Frontend** built with **Vite** and deployed to Vercel

---

## 📦 Running Locally

Clone the repo and run:

```bash
docker-compose up --build
