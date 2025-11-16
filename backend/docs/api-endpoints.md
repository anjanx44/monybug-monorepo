# MonyBug API Endpoints

## Base URL
`http://localhost:8080`

## Health Check
- **GET** `/health` - Check API and database status

## Categories API
- **GET** `/api/v1/categories` - Get all active categories
- **POST** `/api/v1/categories` - Create new category
- **PUT** `/api/v1/categories/:id` - Update category
- **DELETE** `/api/v1/categories/:id` - Delete category (soft delete)

### Category JSON Structure
```json
{
  "id": 1,
  "name": "Food",
  "type": "EXPENSE",
  "priority": "ESSENTIAL",
  "color_code": "#FF5733",
  "is_active": true
}
```

## Transactions API
- **GET** `/api/v1/transactions` - Get all transactions with category details
- **POST** `/api/v1/transactions` - Create new transaction
- **PUT** `/api/v1/transactions/:id` - Update transaction
- **DELETE** `/api/v1/transactions/:id` - Delete transaction (soft delete)

### Transaction JSON Structure
```json
{
  "id": 1,
  "category_id": 1,
  "date": "2024-01-15T00:00:00Z",
  "amount": 250.50,
  "currency": "BDT",
  "description": "Grocery shopping",
  "tags": "food, weekly",
  "created_at": "2024-01-15T10:30:00Z",
  "category": {
    "id": 1,
    "name": "Food",
    "type": "EXPENSE"
  }
}
```

## Budgets API
- **GET** `/api/v1/budgets` - Get all active budgets with category details
- **POST** `/api/v1/budgets` - Create new budget
- **PUT** `/api/v1/budgets/:id` - Update budget
- **DELETE** `/api/v1/budgets/:id` - Delete budget (soft delete)

### Budget JSON Structure
```json
{
  "id": 1,
  "category_id": 1,
  "monthly_limit": 5000.00,
  "month": "2024-01-01T00:00:00Z",
  "is_active": true,
  "created_at": "2024-01-01T00:00:00Z",
  "category": {
    "id": 1,
    "name": "Food",
    "type": "EXPENSE"
  }
}
```

## Expense Insights API
- **GET** `/api/v1/insights` - Get all expense insights
- **POST** `/api/v1/insights` - Create new insight

### Insight JSON Structure
```json
{
  "id": 1,
  "insight_type": "OVERSPEND",
  "message": "Food spending exceeded budget by 20%",
  "period": "2024-01",
  "data": "{\"budget\": 5000, \"actual\": 6000}",
  "generated_at": "2024-01-31T23:59:59Z"
}
```

## Validation Rules

### Categories
- `name`: Required, unique
- `type`: Required, must be "INCOME" or "EXPENSE"
- `priority`: Optional, must be "ESSENTIAL", "IMPORTANT", or "LUXURY"

### Transactions
- `category_id`: Required, must exist
- `amount`: Required, must be > 0
- `date`: Optional, defaults to current date

### Budgets
- `category_id`: Required, must exist
- `monthly_limit`: Required, must be > 0
- `month`: Required

### Insights
- `insight_type`: Required, must be "OVERSPEND", "TREND", or "PATTERN"
- `message`: Required
- `period`: Required

## Error Responses
```json
{
  "error": "Error message description"
}
```

## HTTP Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error)
- `500` - Internal Server Error