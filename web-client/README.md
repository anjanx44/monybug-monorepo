# MonyBug Web Client

React.js frontend for MonyBug expense tracker.

## Setup

```bash
# Install dependencies
npm install

# Start development server (requires Node.js 20+)
npm run dev

# Build for production
npm run build
```

## Features

- ✅ Expense form with category selection
- ✅ Summary cards (daily/monthly totals)
- ✅ Transaction list with date grouping
- ✅ Delete transactions
- ✅ Responsive design
- ✅ TypeScript support
- ✅ API integration with backend

## Components

- `ExpenseForm` - Add new transactions
- `SummaryCards` - Display spending summaries
- `TransactionList` - Show transactions by date
- `api.ts` - Backend API integration

## API Endpoints

- GET `/api/v1/categories` - Load categories
- GET `/api/v1/transactions` - Load transactions
- POST `/api/v1/transactions` - Create transaction
- DELETE `/api/v1/transactions/:id` - Delete transaction

## Tech Stack

- React 18 + TypeScript
- Vite (build tool)
- Axios (HTTP client)
- CSS Variables (styling)

## Note

Requires Node.js 20+ for Vite. If using Node.js 18, consider using Create React App instead.