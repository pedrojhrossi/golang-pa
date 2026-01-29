# Go Appointment System (Multi-tenant)

A robust, production-grade appointment scheduling API built with **Go**, utilizing **Ports and Adapters (Hexagonal) Architecture**. This project serves as a demonstration of clean code principles, multi-tenancy implementation, and decoupled system design.

## Architectural Goal
The primary goal of this repository is to showcase a scalable backend system where the core business logic (The Hexagon) is isolated from infrastructure concerns. By using the Ports and Adapters pattern, the application remains agnostic of specific databases (MongoDB) or delivery mechanisms (HTTP/Chi).

### Key Features
* **Hexagonal Architecture:** Strict separation between Domain, Application, and Infrastructure layers.
* **Multi-tenancy:** Isolated data access patterns to support multiple organizations/clients within a single instance.
* **RESTful API:** Routing handled via `go-chi/chi` for lightweight, idiomatic performance.
* **NoSQL Persistence:** MongoDB integration with repository pattern abstractions.
* **Dependency Injection:** Clean initialization and wiring of components.

## Tech Stack
* **Language:** Go (Golang)
* **Routing:** [chi](https://github.com/go-chi/chi)
* **Database:** MongoDB
* **Containerization:** Docker & Docker Compose (for easy environment setup)

## Project Structure
```text
.
├── cmd/                # Main entry point (Wiring)
├── internal/
│   ├── core/           # The "Hexagon"
│   │   ├── domain/     # Entities (Appointment, Tenant)
│   │   └── services/   # Business logic / Use cases
│   ├── ports/          # Interfaces (Input/Output definitions)
│   └── adapters/       # Implementation (Infrastructure)
│       ├── handler/    # HTTP Adapters (Chi)
│       └── repository/ # Database Adapters (MongoDB)
└── config/             # Configuration management
