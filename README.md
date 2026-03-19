# Sanctor Monorepo

A full-stack application with Go backend and React frontend, organized as a monorepo and containerized with Docker.

## Architecture

This monorepo contains:
- **Backend** (`apps/api`): Go REST API with clean architecture
  - Domain-driven design with handler → service → repository layers
  - Modular structure: auth, user, post, ingestion, digestion
  - Middleware for CORS, logging, authentication
  - Database abstraction layer
  - Pub/Sub messaging system
  
- **Frontend** (`apps/web`): React SPA with Vite
  - Modern React with hooks and functional components
  - Zustand for state management
  - Axios for API integration with interceptors
  - React Router for client-side routing
  - Path aliases for clean imports
  
- **Shared** (`packages/`): Shared code and type definitions
  - Synchronized TypeScript and Go types
- **Orchestration**: Docker Compose for local development and production

## Prerequisites

- Docker and Docker Compose installed
- (Optional) Node.js 18+ and Go 1.21+ for local development

## Quick Start

### Using npm scripts:

```bash
npm run dev          # Start development with hot reload
npm run dev:build    # Rebuild and start development
npm run start        # Start production
npm run down         # Stop all containers
npm run logs         # View logs
```

### Using Make:

```bash
make dev            # Start development with hot reload
make dev-build      # Rebuild and start development
make start          # Start production
make stop           # Stop all containers
make logs           # View logs
make help           # Show all available commands
```

### Using Docker (Production)

Build and run both services with Docker Compose:

```bash
docker-compose up --build
docker compose -f docker-compose.dev.yml up --build # Docker dev mode
```

Access the application:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

### Using Docker (Development with Hot Reload)

For development with volume mounting and hot reloading:

```bash
docker-compose -f docker-compose.dev.yml up
```
apps/
│   ├── backend/              # Go REST API
│   │   ├── handlers/         # HTTP handlers
│   │   ├── models/           # Data models
│   │   ├── Dockerfile
make backend-dev
# or
npm run backend:dev
# or manually:
cd apps/api && go run main.go
```

**Frontend:**
```bash
make frontend-dev
# or
npm run frontend:dev
# or manually:
cd apps/web && npm install && ├── docker-compose.yml        # Production
├── docker-compose.dev.yml    # Development
├── package.json              # Monorepo scripts
├── Makefile                  # Build commands
├── backend/
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── README.md
├── frontend/
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── package.json
│   ├── public/
│   │   └── index.html
│   ├── src/
│   │   ├── App.js
│   │   ├── App.css
│   │   ├── index.js
│   │   └── index.css
│   └── README.md
├── docker-compose.yml         # Production
├── docker-compose.dev.yml     # Development
└── README.md
```

## Docker Commands

**Build images:**
```bash
docker-compose build
```

**Start services:**
```bash
docker-compose up
```

**Start in detached mode:**
```bash
docker-compose up -d
```

**Stop services:**
```bash
docker-compose down
```

**View logs:**
```bash
docker-compose logs -f
```

**Rebuild and restart:**
```bash
docker-compose up --build --force-recreate
```

## API Endpoints

- `GET /health` - Health check
- `GET /api/health` - API health check with JSON response

## Environment Variables

### Backend
- `PORT` - Server port (default: 8080)
- `GO_ENV` - Environment mode (development/production)

### Frontend
- `REACT_APP_API_URL` - Backend API URL (default: http://localhost:8080)

## Development

The backend includes CORS headers for local development. The frontend is configured to communicate with the backend API through environment variables.

### Backend Stack
- Go 1.21 with standard library HTTP server
- Clean architecture (handler/service/repository)
- In-memory storage (ready for database integration)
- UUID for ID generation

### Frontend Stack
- React 18 with Vite for blazing-fast development
- Zustand for lightweight state management
- Axios for HTTP requests with interceptors
- React Router v6 for routing
- Path aliases for clean imports

## Documentation

- 📘 [Backend Structure Guide](BACKEND_STRUCTURE.md) - Complete backend architecture
- 📗 [Frontend Structure Guide](FRONTEND_STRUCTURE.md) - Frontend organization & usage

## License

MIT