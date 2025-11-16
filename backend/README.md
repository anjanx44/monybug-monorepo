# MonyBug Backend API

Go/Gin REST API for personal expense management with PostgreSQL and GORM.

## Quick Start

```bash
# Install dependencies
go mod tidy

# Set environment variables in .env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=monybug_db
API_PORT=8080

# Run migrations and start server
go run cmd/main.go
```

## Available APIs

### Health Check
- `GET /health` - API status

### Categories (CRUD)
- `GET /api/v1/categories` - List all
- `POST /api/v1/categories` - Create
- `PUT /api/v1/categories/:id` - Update
- `DELETE /api/v1/categories/:id` - Delete

### Transactions (CRUD)
- `GET /api/v1/transactions` - List all
- `POST /api/v1/transactions` - Create
- `PUT /api/v1/transactions/:id` - Update
- `DELETE /api/v1/transactions/:id` - Delete

### Budgets (CRUD)
- `GET /api/v1/budgets` - List all
- `POST /api/v1/budgets` - Create
- `PUT /api/v1/budgets/:id` - Update
- `DELETE /api/v1/budgets/:id` - Delete

### Expense Insights
- `GET /api/v1/insights` - List all
- `POST /api/v1/insights` - Create

## Features
- ✅ GORM ORM with PostgreSQL
- ✅ Auto-migrations
- ✅ Input validation
- ✅ Soft deletes
- ✅ Foreign key relationships
- ✅ JSON API responses
- ✅ Error handling

## Tech Stack
- **Language**: Go 1.21+
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Architecture**: Clean Architecture (Models, Repositories, Handlers)

See `docs/api-endpoints.md` for detailed API documentation.