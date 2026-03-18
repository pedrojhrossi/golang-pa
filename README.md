# 📅 Go Appointment System (Multi-tenant)

A robust, production-grade appointment scheduling API built with **Go**, utilizing **Ports and Adapters (Hexagonal) Architecture**. This project serves as a demonstration of clean code principles, multi-tenancy implementation, and decoupled system design.

## 🏗 Architectural Approach

The primary goal of this repository is to showcase a scalable backend system where the core business logic (**The Hexagon**) is isolated from infrastructure concerns. By using the Ports and Adapters pattern, the application remains agnostic of specific databases (MongoDB) or delivery mechanisms (HTTP/Chi).



### Key Benefits:
* **Independence:** The "Core" doesn't care if it's talking to a database, a mock, or a CLI.
* **Testability:** Business logic can be tested in isolation without needing a database or a web server.
* **Flexibility:** Swap MongoDB for PostgreSQL or Chi for Echo by only changing the Adapter layer.
* **Maintainability:** Database schema changes or library updates do not "leak" into the business rules.

---

## 🌟 Key Features

* **Hexagonal Architecture:** Strict separation between Domain, Application, and Infrastructure layers.
* **Multi-tenancy:** Isolated data access patterns to support multiple organizations within a single instance.
* **RESTful API:** Lightweight, idiomatic performance using `go-chi/chi`.
* **NoSQL Persistence:** MongoDB integration using the Repository Pattern.
* **Dependency Injection:** Clean initialization and wiring of components in the application entry point.
* **Conflict Prevention:** The service enforces a "no-overlap" policy by checking existing schedules via the Repository Port before booking.

---

## 📦 Project Structure

```text
.
├── cmd/
│   └── app/                # Application Entry Point (The "Composer" / Wiring)
├── config/                 # Configuration Management (Environment-based)
├── internal/               # Private code (The Hexagon)
│   ├── core/               # The "Brain"
│   │   ├── domain/         # Entities & Factory Functions (Tenant, Appointment)
│   │   └── services/       # Use Cases & Business Rules
│   ├── ports/              # The "Contracts" (Go Interfaces)
│   └── adapters/           # The "Plugs" (Infrastructure)
│       ├── handler/        # Primary Adapters: HTTP/REST (Chi Router)
│       └── repository/     # Secondary Adapters: Persistence (MongoDB)
├── docker-compose.yml      # Infrastructure Orchestration (MongoDB)
└── go.mod                  # Dependency Management
```
---

## 🛠 Deep Dive: Layer Explanations

### 1. The Core (Domain & Services)
* **Domain Entities:** We use **Self-Validating Entities**. For example, the `Tenant` struct uses `google/uuid` for identity, ensuring that IDs are generated at the domain level.
* **Services:** The service layer coordinates logic. It talks only to **Ports** (interfaces), never concrete implementations. This ensures the business flow is protected from external infrastructure changes.

### 2. The Ports (Interfaces)
Ports define the contracts for how the core interacts with the outside world:
* **Input Ports:** Methods the Service exposes to the API (e.g., `RegisterTenant`, `Schedule`, `CancelAppointment`).
* **Output Ports:** Methods the Service requires (e.g., `FindOverlapping`, `Update`, `Create`, `GetByID`).

### 3. The Adapters (Infrastructure)
* **HTTP Handler (Chi):** A **Primary Adapter** that converts JSON requests into domain-friendly data and triggers the service.
* **MongoDB Repository:** A **Secondary Adapter** that implements the Output Port. It handles BSON mapping and database-specific queries.

---

## 🚦 Getting Started

### Prerequisites
* **Go 1.22+**
* **Docker & Docker Compose**

### Installation & Run

1. **Start Infrastructure:**
```bash
   docker-compose up -d
```

2. **Set Environment Variables:**
```bash
  export MONGO_USER=admin
  export MONGO_PASS=12345
  export MONGO_HOST=localhost
  export MONGO_PORT=27017
  export DB_NAME=appointment_db
  export APP_PORT=9080
```

3. **Run the Application:**
```bash
  go run cmd/app/main.go
```

---

## 🧪 Testing the API
### Tenants
* **Create a Tenant (POST):**
  `curl -X POST http://localhost:9080/tenants -H "Content-Type: application/json" -d '{"name": "Health Clinic Alpha", "email": "admin@alpha.com"}'`

* **List Active Tenants (GET):**
  `curl -X GET http://localhost:9080/tenants`
 
* **List Tenants including Archived (GET):**
  `curl -X GET "http://localhost:9080/tenants?archive=true"`
 
* **Get Tenant by ID (GET):**
  `curl -X GET http://localhost:9080/tenants/<TENANT_ID>`
 
* **Create a Tenant (DELETE):**
  `curl -X DELETE http://localhost:9080/tenants/<TENANT_ID>`
 
### Appointments
* **Schedule an Appointment (POST):**
  `curl -X POST http://localhost:9080/tenants/<TENANT_ID>/appointments -H "Content-Type: application/json" -d '{"patient_id": "<USER_UUID>", "patient_name": "John Doe", "start_time": "2026-03-30T10:00:00Z", "end_time": "2026-03-30T11:00:00Z"}'`

* **List Tenant Appointments (GET):**
  `curl -X GET http://localhost:9080/tenants/<TENANT_ID>/appointments`

* **Get Specific Appointment (GET):**
  `curl -X GET http://localhost:9080/tenants/<TENANT_ID>/appointments/<APPOINTMENT_ID>`

* **Cancel an Appointment (DELETE):**
  `curl -X DELETE http://localhost:9080/tenants/<TENANT_ID>/appointments/<APPOINTMENT_ID>/cancel`

---

## ✅ Tech Stack
* **Language:** Go (Golang)
* **Routing:** Chi v5
* **Database:** MongoDB v2 Driver
* **ID Generation:** Google UUID
* **Containerization:** Docker & Docker Compose

## 🗺 Roadmap
* [x] Core Hexagonal Architecture Setup
* [x] Tenant Management Vertical Slice
* [x] Appointment Domain (Scheduling & Conflicts)
* [ ] **Next:** Multi-tenant isolation middleware
* [ ] Unit Testing with Mocking (Testify/GoMock)
* [ ] JWT Authentication per Tenant
