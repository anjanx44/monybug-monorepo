# MonyBug Backend API Documentation

## Overview
Go + Gin REST API for personal expense management with PostgreSQL database.

## Technology Stack
- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **Architecture**: Clean Architecture (Models, Handlers, Routes)

## Project Structure
```
backend/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── models/              # Database models
│   │   ├── category.go
│   │   ├── transaction.go
│   │   └── budget.go
│   ├── handlers/            # HTTP handlers
│   │   ├── category.go
│   │   ├── transaction.go
│   │   └── budget.go
│   ├── database/            # Database connection
│   │   └── db.go
│   └── routes/              # Route definitions
│       └── routes.go
├── go.mod                   # Go module
└── .env                     # Environment variables
```

## API Endpoints

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | Get all categories |
| POST | `/api/categories` | Create new category |
| PUT | `/api/categories/:id` | Update category |
| DELETE | `/api/categories/:id` | Delete category |

### Transactions
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/transactions` | Get all transactions |
| POST | `/api/transactions` | Add new expense/income |
| GET | `/api/transactions/:id` | Get transaction by ID |
| PUT | `/api/transactions/:id` | Update transaction |
| DELETE | `/api/transactions/:id` | Delete transaction |

### Analytics
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/analytics/monthly` | Monthly spending summary |
| GET | `/api/analytics/categories` | Category-wise breakdown |
| GET | `/api/analytics/trends` | Spending trends |

## Request/Response Examples

### Add Transaction
```http
POST /api/transactions
Content-Type: application/json

{
  "category_id": 1,
  "amount": 100.00,
  "currency": "BDT",
  "description": "baby toy",
  "tags": "toy, kids"
}
```

### Response
```json
{
  "id": 1,
  "category_id": 1,
  "date": "2024-01-15T10:30:00Z",
  "amount": 100.00,
  "currency": "BDT",
  "description": "baby toy",
  "tags": "toy, kids",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### Get Monthly Summary
```http
GET /api/analytics/monthly?month=2024-01
```

```json
{
  "month": "2024-01",
  "total_expenses": 5000.00,
  "total_income": 15000.00,
  "net_amount": 10000.00,
  "categories": [
    {
      "name": "Food",
      "amount": 2000.00,
      "percentage": 40
    }
  ]
}
```

## Database Models

### Transaction Model
```go
type Transaction struct {
    ID          int       `json:"id" db:"transaction_id"`
    CategoryID  int       `json:"category_id" db:"category_id"`
    Date        time.Time `json:"date" db:"transaction_date"`
    Amount      float64   `json:"amount" db:"amount"`
    Currency    string    `json:"currency" db:"currency"`
    Description string    `json:"description" db:"description"`
    Tags        string    `json:"tags" db:"tags"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
```

### Category Model
```go
type Category struct {
    ID       int    `json:"id" db:"category_id"`
    Name     string `json:"name" db:"name"`
    Type     string `json:"type" db:"type"`
    Priority string `json:"priority" db:"priority"`
    Color    string `json:"color_code" db:"color_code"`
    IsActive bool   `json:"is_active" db:"is_active"`
}
```

## Environment Variables
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=monybug_db
API_PORT=8080
```

## Development Setup

### 1. Initialize Go Module
```bash
cd backend
go mod init monybug-backend
go mod tidy
```

### 2. Install Dependencies
```bash
go get github.com/gin-gonic/gin
go get github.com/lib/pq
go get github.com/joho/godotenv
```

### 3. Run Server
```bash
go run cmd/main.go
```

## Learning Milestones

### Phase 1: Basic Setup ✅
- [x] Project structure
- [x] Go module initialization
- [x] Basic Gin server

### Phase 2: Database Integration
- [ ] PostgreSQL connection
- [ ] Database models
- [ ] Basic CRUD operations

### Phase 3: Core API
- [ ] Transaction endpoints
- [ ] Category management
- [ ] Error handling

### Phase 4: Analytics
- [ ] Monthly summaries
- [ ] Category breakdowns
- [ ] Spending trends

### Phase 5: Advanced Features
- [ ] Budget tracking
- [ ] Expense insights
- [ ] Data validation

## Error Handling
```json
{
  "error": "Invalid request",
  "message": "Amount must be greater than 0",
  "code": 400
}
```

## Testing
```bash
# Run tests
go test ./...

# Test coverage
go test -cover ./...
```