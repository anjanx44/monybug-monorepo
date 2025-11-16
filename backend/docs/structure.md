# MonyBug Backend Structure

## Project Architecture

```
backend/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── database/
│   │   ├── db.go              # Raw SQL connection
│   │   ├── gorm.go            # GORM connection
│   │   ├── health.go          # Health check
│   │   └── migrate.go         # Migration runner
│   ├── handlers/
│   │   ├── category_handler.go    # Category CRUD endpoints
│   │   ├── transaction_handler.go # Transaction CRUD endpoints
│   │   ├── budget_handler.go      # Budget CRUD endpoints
│   │   └── insight_handler.go     # Insight endpoints
│   ├── models/
│   │   ├── category.go        # Category model with GORM tags
│   │   ├── transaction.go     # Transaction model with associations
│   │   ├── budget.go          # Budget model
│   │   └── expense_insight.go # ExpenseInsight model
│   ├── repositories/
│   │   ├── category_repository.go    # Category data access
│   │   ├── transaction_repository.go # Transaction data access
│   │   ├── budget_repository.go      # Budget data access
│   │   └── insight_repository.go     # Insight data access
│   └── routes/
│       └── routes.go          # Route definitions
├── migrations/
│   ├── 001_initial_schema.sql     # Initial database schema
│   └── 002_enhanced_schema.sql    # Enhanced features
├── docs/
│   ├── api-endpoints.md       # API documentation
│   └── structure.md           # This file
├── .env                       # Environment variables
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
└── README.md                  # Project overview
```

## Architecture Layers

### 1. **Entry Point** (`cmd/`)
- `main.go` - Application bootstrap, database connections, server startup

### 2. **HTTP Layer** (`internal/routes/`, `internal/handlers/`)
- **Routes**: URL mapping and middleware
- **Handlers**: HTTP request/response processing, validation, JSON binding

### 3. **Business Logic** (`internal/repositories/`)
- **Repositories**: Data access layer using GORM
- Database queries, CRUD operations, business rules

### 4. **Data Layer** (`internal/models/`, `internal/database/`)
- **Models**: GORM entities with tags and associations
- **Database**: Connection management, migrations, health checks

### 5. **Infrastructure** (`migrations/`, `.env`)
- **Migrations**: Database schema versioning
- **Config**: Environment-based configuration

## Data Flow

```
HTTP Request → Routes → Handlers → Repositories → GORM → PostgreSQL
                ↓
HTTP Response ← JSON ← Models ← Database Results ← SQL Query
```

## Key Design Patterns

### Clean Architecture
- **Separation of Concerns**: Each layer has single responsibility
- **Dependency Inversion**: Handlers depend on repository interfaces
- **Testability**: Each layer can be tested independently

### Repository Pattern
- **Data Access Abstraction**: Repositories hide database implementation
- **GORM Integration**: Leverages ORM for type-safe queries
- **Error Handling**: Consistent error propagation

### RESTful API Design
- **Resource-based URLs**: `/api/v1/categories`, `/api/v1/transactions`
- **HTTP Methods**: GET, POST, PUT, DELETE for CRUD operations
- **JSON Communication**: Structured request/response format

## Database Schema

### Core Entities
- **Categories**: Expense/Income classification with priorities
- **Transactions**: Financial records linked to categories
- **Budgets**: Monthly spending limits per category
- **ExpenseInsights**: AI-generated spending analysis

### Relationships
- `Categories` → `Transactions` (One-to-Many)
- `Categories` → `Budgets` (One-to-Many)
- GORM handles foreign keys and preloading

## Technology Stack

| Component | Technology | Purpose |
|-----------|------------|---------|
| Language | Go 1.21+ | Performance, concurrency |
| Framework | Gin | HTTP routing, middleware |
| Database | PostgreSQL | ACID compliance, reliability |
| ORM | GORM | Type-safe queries, migrations |
| Config | godotenv | Environment management |

## Development Workflow

1. **Models**: Define GORM entities
2. **Migrations**: Create SQL schema files
3. **Repositories**: Implement data access methods
4. **Handlers**: Create HTTP endpoints
5. **Routes**: Wire handlers to URLs
6. **Testing**: Verify API functionality

## Scalability Considerations

- **Horizontal Scaling**: Stateless handlers support load balancing
- **Database Optimization**: Indexes on frequently queried columns
- **Caching**: Can add Redis for session/query caching
- **Microservices**: Clean architecture enables service extraction