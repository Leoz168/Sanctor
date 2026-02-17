# Sanctor Backend (Go)

A well-structured Go REST API following clean architecture principles.

## Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
│
├── internal/                  # Private application code
│   ├── auth/                  # Authentication module
│   │   ├── handler.go         # HTTP handlers
│   │   ├── service.go         # Business logic
│   │   ├── repository.go      # Data access
│   │   └── model.go           # Data models
│   │
│   ├── user/                  # User module
│   ├── post/                  # Post module
│   ├── ingestion/             # Data ingestion
│   ├── digestion/             # Scheduled processing
│   │   └── cron.go
│   │
│   ├── pubsub/                # Pub/Sub messaging
│   ├── middleware/            # HTTP middleware
│   ├── config/                # Configuration
│   └── database/              # Database connection
│
├── migrations/                # Database migrations
├── pkg/                       # Reusable packages
├── go.mod                     # Go dependencies
└── Dockerfile                 # Container definition
```

## Architecture

This project follows **Domain-Driven Design (DDD)** with a layered architecture:

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic and validation
- **Repository Layer**: Data persistence abstraction
- **Model Layer**: Domain entities and DTOs

## Development

### Local Development

```bash
go run cmd/api/main.go
```

### With Docker

```bash
docker build -t sanctor-backend .
docker run -p 8080:8080 sanctor-backend
```

### Testing

```bash
go test ./...
```

## API Endpoints

### Health
- `GET /health` - Health check
- `GET /api/health` - API health check

### Users
- `GET /api/users` - List all users
- `GET /api/users/get?id={id}` - Get user by ID
- `POST /api/users/create` - Create new user
- `PUT /api/users/update?id={id}` - Update user
- `DELETE /api/users/delete?id={id}` - Delete user

### Auth (TODO)
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration
- `POST /api/auth/refresh` - Refresh token

### Posts (TODO)
- `GET /api/posts` - List all posts
- `POST /api/posts/create` - Create new post

## Configuration

Environment variables:
- `PORT` - Server port (default: 8080)
- `GO_ENV` - Environment (development/production)
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET` - JWT signing secret

## Adding a New Module

1. Create directory under `internal/`
2. Create files: `model.go`, `handler.go`, `service.go`, `repository.go`
3. Register routes in `cmd/api/main.go`
4. Add tests

## License

MIT
