# FiberStarter

A clean architecture starter template for building REST APIs with [Go Fiber](https://gofiber.io/) framework.

## Features

- **Go Fiber Framework** - Fast and lightweight web framework
- **GORM ORM** - Elegant database abstraction layer
- **Multi-database Support** - PostgreSQL or SQLite
- **Handlebars Templating** - Server-side template rendering
- **Built-in Middlewares** - Recovery, request ID, structured logging, compression, idempotency, and CORS
- **Request Validation** - Integrated struct validation for request payloads
- **Docker Support** - Predefined Dockerfile for containerized deployment
- **Clean Architecture** - Organized into transport, service, and repository layers

## Prerequisites

- Go 1.20 or higher
- PostgreSQL or SQLite (depending on your choice)
- Docker (optional, for containerized deployment)

## Getting Started

### Installation

1. Clone the repository:
```bash
git clone https://github.com/jakubtomas-cz/fiber-starter.git
cd fiber-starter
```

2. Copy the environment file:
```bash
cp .env.example .env
```

3. Install dependencies:
```bash
go mod download
```

4. Set up your database:
```bash
# .env

# For PostgreSQL
DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable"

# For SQLite (just a filename of the DB)
DATABASE_URL="gorm.db"
```

4. Run the server:
```bash
go run cmd/server/main.go
```

### Docker Deployment

Build and run the application in a Docker container:

```bash
docker build -t fiber-starter .
docker run -e DATABASE_URL="your-database-url" -p 1080:1080 fiber-starter
```

## Why Clean Architecture?

This starter adopts **Clean Architecture** principles to provide several key benefits:

- **Separation of Concerns** - Each layer has a specific responsibility, making the codebase easier to understand and maintain
- **Flexibility** - Swapping databases (PostgreSQL ↔ SQLite) or changing HTTP framework details doesn't affect your business logic
- **Scalability** - Clear boundaries make it easy to refactor, extend, and organize code as your project grows
- **Reusability** - Service layer logic can be reused across different transports (REST, gRPC, etc.)

The three-layer architecture (Transport → Service → Repository) creates a unidirectional dependency flow where each layer only depends on layers below it, preventing circular dependencies and tight coupling.

## Folder Structure

```
fiber-starter/
├── cmd/
│   └── server/
│       └── main.go      # Entrypoint
├── internal/
│   ├── handlers/        # HTTP handlers
│   ├── pages/           # Handlers returning/rendering Handlebar views
│   ├── models/          # Data models (GORM entities and user defined structures)
│   ├── repository/      # Repository layer (database operations)
│   ├── service/         # Service layer (business logic)
│   ├── transport/       # Transport layer defining middlewares, HTTP handlers and pages
│   ├── middlewares/     # Custom middleware functions
│   └── utils/           # Utility functions and helpers
├── views/               # Handlebars templates
├── Dockerfile           # Docker configuration
├── go.mod               # Go module definition
└── go.sum               # Go module checksums
```

### Directory Descriptions

- **cmd/server** - Application entry point
- **internal/handlers** - HTTP request handlers and controllers for processing API requests
- **internal/pages** - Page handlers that render Handlebars templates for server-side rendered views
- **internal/models** - GORM database models with `DBModel` or `gorm.Model` base
- **internal/repository** - Data access layer with GORM queries and auto-migrations
- **internal/service** - Business logic layer handling core application operations
- **internal/transport** - HTTP route handlers, request parsing, and response formatting
- **internal/middlewares** - Custom middleware functions for request/response processing
- **internal/utils** - Shared utility functions used across layers
- **views** - Handlebars HTML templates for server-side rendering

## Architecture

The project follows **Clean Architecture** principles with three main layers:

### Transport Layer
Handles HTTP requests/responses using Fiber routes and middlewares.

### Service Layer
Contains business logic and orchestrates operations between transport and repository layers.

### Repository Layer
Manages all database operations using GORM, abstracting the data persistence logic. Also handles automatic schema migrations via GORM's `AutoMigrate` functionality.

## Models

User-defined models can be constructed using either:

- **GORM Model** - `gorm.Model` provides `ID`, `CreatedAt`, `UpdatedAt`, and `DeletedAt` fields
- **Custom DBModel** - Lightweight alternative with just `ID`, `CreatedAt`, and `UpdatedAt` fields

Both approaches support automatic timestamp management and primary key handling.

## Auto Migrations

After defining your GORM models, enable automatic schema migrations by updating the `AutoMigrate` function in `internal/repository/repository.go`:

```go
// internal/repository/repository.go
err = db.AutoMigrate(
    &models.User{},
    &models.Post{},
    // Add your other models here
)
if err != nil {
    return repository, err
}
```

Simply uncomment the `AutoMigrate` call and pass your model structs. GORM will automatically create or update database tables based on your model definitions when the application starts.

## Middlewares

The starter automatically sets up essential middlewares:

- **Recovery** - Gracefully handles panics and crashes
- **Request ID** - Tracks requests with unique identifiers
- **Structured Logging** - JSON-formatted logs with request details (time, method, path, status, latency, IP, ID)
- **Compression** - Gzip compression for responses
- **Idempotency** - Ensures safe retry of requests
- **CORS** - Cross-Origin Resource Sharing support

## Templating

This starter includes **Handlebars** templating engine for server-side rendering of HTML templates.

## Development

The project includes a `Makefile` with useful commands to streamline development:

- `make dev` - Run the server with hot reload using Nodemon (requires Node.js)
- `make run` - Run the server normally
- `make db` - Start a PostgreSQL Docker container with default credentials (user: `admin`, password: `admin`)

Example:
```bash
# Terminal 1: Start the database
make db

# Terminal 2: Run the server with hot reload
make dev
```

### Creating New Features

To create new page or handler, refer to the example implementations in the relevant folders:

- **New HTTP Handler** - Check `internal/handlers` for existing handler examples to follow the same pattern
- **New Page/View** - Look at example in `internal/pages` to understand how to create server-rendered pages with Handlebars templates
- **New Route** - Add your route definitions in `internal/transport/endpoints.go` following the existing route setup patterns

This ensures consistency with the established architecture and coding style across your project.

## Validation

The starter includes integrated request validation using the `go-playground/validator` package. Easily validate incoming request payloads with struct tags:

```go
type CreateUserRequest struct {
    Name  string `validate:"required,min=3,max=50"`
    Email string `validate:"required,email"`
}
```

Validation errors are automatically returned as structured HTTP responses, ensuring consistent error handling across your API.

## Environment Variables

- `DATABASE_URL` - Connection string for your database
  - PostgreSQL: `postgres://user:password@host:port/database?sslmode=disable`
  - SQLite: `filename.db`
- `PORT` - Server port (default: `1080`)
- `PROXY_HEADER` - Proxy header for X-Forwarded-For or similar headers (optional)

## Customization & Enhancement

This starter provides a solid foundation for building REST APIs with clean architecture principles. Feel free to enhance and customize it according to your project needs:

- Add additional middleware for authentication, rate limiting, or custom logic
- Extend the validation layer with custom validators
- Integrate additional packages for caching, messaging, or file handling
- Modify the folder structure to match your specific project requirements
- Add new layers (e.g., domain models, use cases) as your application grows

The modular design ensures that any enhancements integrate seamlessly with the existing architecture.

## Links

Explore the main technologies and dependencies used in this project:

- [Go Fiber](https://gofiber.io/) - Fast and lightweight web framework for Go
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [Handlebars](https://handlebarsjs.com/) - Logic-less templates
- [Go Playground Validator](https://github.com/go-playground/validator) - Struct and field validation for Go

## License

MIT
