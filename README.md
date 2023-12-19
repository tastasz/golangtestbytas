# Build Project by Golang

## About the project

### Design

The project is created base on Domain-Driven Design (DDD) [Crean Architecture], OOP (Design Patterns) and SOLID Design principles you can see more basic infromation (thai language) on this referance

- [Domain-Driven Design (DDD)](https://medium.com/devspree/domain-driven-design-934c495e7e0b)

- [Object Oriented Programming (OOP)](https://www.saladpuk.com/beginner-1/oop)

- [Design Patterns](https://www.saladpuk.com/software-design/designpatterns)

- [SOLID Design](https://www.saladpuk.com/basic/solid)

### Layout (Example)

```tree
project-root/
|-- cmd/
|   |-- main.go
|
|-- pkg/
|   |-- app/
|   |   |-- cart.go
|
|   |-- domain/
|   |   |-- model/
|   |   |   |-- user.go
|   |   |   |-- order.go
|   |   |
|   |   |-- service/
|   |       |-- user_service.go
|   |       |-- order_service.go
|   |
|   |-- infrastructure/
|       |-- persistence/
|       |   |-- user_repository.go
|       |   |-- order_repository.go
|       |
|       |-- handlers/
|           |-- cart_handelers.go
|
|-- scripts/
|   |-- db_migration_script.sql
|
|-- go.mod
|-- go.sum
|-- README.md
```

A brief description of the layout:

Explanation:

- `cmd:` This directory contains the main entry point for your application. The `main.go` file initializes the application and starts the server.

- `pkg/app:` This directory contains the application layer. It includes the application service and any other application-specific logic.

- `pkg/domain:` This directory contains the domain layer, where you define your business logic. It includes the domain models, repositories, and services.

  - `model:` This subdirectory contains the domain models (e.g., User, Order).

  - `service:` This subdirectory contains domain services that encapsulate business logic.

- `pkg/infrastructure:` This directory contains the infrastructure layer, responsible for interacting with external systems (e.g., databases, external services).

  - `persistence:` This subdirectory contains implementations of repositories, responsible for data storage.

  - `handlers:` This subdirectory contains the HTTP server and handlers.

- `scripts:` This directory includes any scripts related to your project, such as database migration scripts.

- `go.mod, go.sum:` These files manage the project's dependencies.

- `README.md:` Documentation for your project.

## Getting started

1. **Initialize Go Modules:**
   If the project uses Go modules, you need to initialize them:

   ```bash
   go mod download
   ```

   This command downloads the dependencies specified in the `go.mod` file.

2. **Run the Application:**

   ```bash
   go run .
   ```
