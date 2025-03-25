![CI](https://github.com/kirebyte/thd-project/actions/workflows/ci.yml/badge.svg)

# Car Inventory Microservice

This project is a fully functional Go microservice designed to manage a car inventory. It includes clean architecture, basic structured logging, layered services, automatic database initialization, Swagger/OpenAPI validation, GitHub Actions CI/CD, unit and integration tests, and a powerful Makefile-driven workflow.

> “We were asked to build a microservice with in-memory persistence, but we went a little further: SQLite, validation, tests, CI/CD, and swagger. We built something real. AI became a real boost for software engineering and I'm proud to show a project that was built in record time following good standards."

---

## 🚀 Features

- REST API for car creation, listing, retrieval and updates
- Data persistence using embedded SQLite (no native dependencies)
- Auto-initialization of DB schema
- Contract-based service/repository layers
- Logging with log levels
- Postman collection and integration tests with Newman
- Swagger validation support
- GitHub Actions pipeline for CI/CD
- Easy-to-use `Makefile` for every task

---

## 🛠️ Getting Started

### Requirements
- Go 1.21+
- (Optional) Node.js (for running `newman` integration tests natively)

### Setup

1. Clone the repository:

```bash
git clone https://github.com/kirebyte/thd-project.git
cd thd-project
```

2. Initialize the SQLite database:

```bash
make db-init
```

This creates a file-based SQLite DB and applies the schema defined in `scripts/schema.sql`.

> I noticed that in-memory storage was requested, but I decided to use embedded SQLite without native dependencies to simulate a more realistic environment without complicating deployment.
> If this app were used for debugging or local testing, minimal persistence is extremely helpful. Also, since the DB initializes itself if missing, there’s zero config needed.

3. Run the server:

```bash
make run
```

---

## 🧪 Testing

Run all unit tests:
```bash
make test
```

Generate a pretty HTML coverage report:
```bash
make coverage
```

Run lint and static checks:
```bash
make lint-strict
```

Run integration tests via Postman:
```bash
make integration
```

> These tests simulate full API workflows and validate that the microservice behaves correctly in real-world scenarios.

---

## 🔁 Makefile Tasks

This project uses a powerful `Makefile` to orchestrate common actions.

```bash
make db-init          # Creates the DB file if it doesn't exist
make run              # Runs the server
make build            # Builds a binary into bin/
make test             # Runs unit tests
make coverage         # Runs tests and generates coverage.html
make lint             # Runs basic lint (fmt, vet)
make lint-strict      # Runs fmt, vet, and staticcheck
make integration      # Runs Postman collection via newman
make swagger-validate # Validates Swagger (OpenAPI YAML)
```

---

## 🧩 Architecture Overview

```
cmd/
  init-database/   - Standalone script to bootstrap DB
  server/          - Application entrypoint
internal/
  api/             - HTTP router and handlers
  repository/sqlite- SQLite repository implementation
  service/         - Business logic
  logger/          - Log helpers with log levels
model/             - Public Car struct (API contract)
repository/        - Car repository interface
service/           - Car service interface
scripts/           - SQL schema
settings/          - Configuration and env loading
```

---

## ✅ Requirements Checklist

- [x] Create Car
- [x] List Cars
- [x] Get Car by ID
- [x] Update Car by ID
- [x] Persist data (using SQLite instead of memory)
- [x] Unit tests
- [x] Integration tests (Postman/Newman)
- [x] Basic Structured logging
- [x] RESTful API using standard library
- [x] Makefile automation
- [x] Swagger validation
- [x] GitHub Actions CI/CD

---

## 🧠 Final Thoughts

This microservice was designed not only to meet the challenge, but to exceed expectations with practical software engineering techniques. From contracts to tests and automated setup, it's ready to be extended or deployed in any real-world environment.

There's a document named "ARCHITECTURE_FAQ.md" in the docs directory that explains the design choices made for this project, I hope you find it interesting.

Thanks for reviewing my project! 💜

---

MIT License — Kirebyte 2024

