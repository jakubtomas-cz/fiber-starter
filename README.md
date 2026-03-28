# FiberStarter

A production-ready starter template for building REST APIs and server-rendered web apps with [Go Fiber v3](https://gofiber.io/), following Clean Architecture principles.

## Features

- **Go Fiber v3** — Fast HTTP framework built on Fasthttp
- **Clean Architecture** — Transport → Service → Repository layering
- **GORM** — Multi-database support (PostgreSQL, SQLite)
- **Handlebars Templating** — Server-side rendering out of the box
- **Middleware Stack** — Recovery, request ID, structured JSON logging, compression, idempotency, CORS, rate limiting
- **Request Validation** — Struct-based validation via `go-playground/validator`
- **Docker Ready** — Multi-stage build with Alpine runtime
- **No CGO** — Pure Go, no C dependencies required

## Prerequisites

- Go 1.25+
- PostgreSQL or SQLite
- Docker (optional)

## Getting Started

### 1. Clone

```bash
git clone https://github.com/jakubtomas-cz/fiber-starter.git
cd fiber-starter
```

### 2. Configure environment

```bash
cp .env.example .env
```

Edit `.env`:

```env
# PostgreSQL
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable

# SQLite (just a filename)
DATABASE_URL=gorm.db

PORT=8090
```

### 3. Run

```bash
# Start database (PostgreSQL via Docker)
make db

# Development with hot reload
make dev

# Or run directly
make run
```

## Project Structure

```
fiber-starter/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── handlers/                # API request handlers
│   ├── pages/                   # Server-rendered page handlers
│   ├── models/                  # GORM models
│   ├── repository/              # Data access layer
│   ├── service/                 # Business logic layer
│   ├── transport/               # Fiber app setup (routes, middleware, config)
│   ├── middlewares/             # Custom middleware
│   └── utils/                   # Shared utilities and constants
├── views/                       # Handlebars templates
├── Dockerfile
├── Makefile
├── go.mod
└── .env.example
```

## Architecture

The project enforces a **unidirectional dependency flow**:

```
Transport → Service → Repository
```

Each layer only depends on the layer below it, preventing circular dependencies and keeping business logic isolated from HTTP and database concerns.

| Layer | Responsibility |
|-------|---------------|
| **Transport** | HTTP routing, middleware, request parsing, response formatting |
| **Service** | Business logic, orchestration between layers |
| **Repository** | Database access via GORM, auto-migrations |

## Routes

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check (rate limited) |
| `GET` | `/` | Homepage (server-rendered) |
| `GET` | `/api/hello` | Example JSON endpoint |

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile binary to `bin/app` |
| `make run` | Run the server |
| `make dev` | Run with hot reload (requires Node.js) |
| `make db` | Start PostgreSQL in Docker |

## Docker

```bash
docker build -t fiber-starter .
docker run -e DATABASE_URL="your-database-url" -p 8090:8090 fiber-starter
```

## Adding Features

**New API endpoint** — add a method to `internal/handlers/`, register it in `internal/transport/endpoints.go`

**New page** — add a handler in `internal/pages/`, create a template in `views/`, register the route in `internal/transport/endpoints.go`

**New database model** — define the struct in `internal/models/`, add it to `AutoMigrate` in `internal/repository/repository.go`

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | Database connection string | — |
| `PORT` | Server port | `8090` |
| `PROXY_HEADER` | Proxy header (e.g. `X-Forwarded-For`) | — |

## Validation

Validate request payloads using struct tags:

```go
type CreateUserRequest struct {
    Name  string `validate:"required,min=3,max=50"`
    Email string `validate:"required,email"`
}
```

Validation errors are returned as structured JSON responses automatically.

## Auto Migrations

Register your models for automatic schema migration on startup:

```go
// internal/repository/repository.go
err = db.AutoMigrate(
    &models.User{},
    &models.Post{},
)
```

## Dependencies

- [Go Fiber v3](https://gofiber.io/) — HTTP framework
- [GORM](https://gorm.io/) — ORM library
- [GORM PostgreSQL Driver](https://github.com/go-gorm/postgres) — PostgreSQL support
- [GORM SQLite Driver](https://github.com/glebarez/sqlite) — SQLite support (pure Go, no CGO)
- [Handlebars](https://handlebarsjs.com/) — Templating engine
- [godotenv](https://github.com/joho/godotenv) — `.env` file loading
- [go-playground/validator](https://github.com/go-playground/validator) — Struct validation

## License

MIT
