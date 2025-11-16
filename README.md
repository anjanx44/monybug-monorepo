🔴 MonyBug: Full-Stack Personal Expense Manager

✨ Overview

MonyBug is a high-performance, full-stack personal finance application designed for meticulous expense tracking and financial analysis. It's structured as a Monorepo to efficiently manage the separate but cohesive backend API and the cross-platform frontend clients.

Our primary goal is to provide a reliable, fast, and secure tool, leveraging the power of Go for backend efficiency and PostgreSQL for rock-solid transactional integrity, ensuring accurate financial reporting every time.

🚀 Key Features

Transactional Integrity: Leveraging PostgreSQL for guaranteed ACID compliance in all financial data.

Cross-Platform Client: A single React/React Native codebase supports Web, iOS, and Android clients.

Custom Categories: Users can define custom income and expense categories.

Detailed Reporting: Generate monthly, quarterly, and annual summaries based on category and type (INCOME/EXPENSE).

High Performance: API endpoints optimized using the Go programming language and Gin framework.

🛠️ Technology Stack

MonyBug is built on a robust, type-safe, and high-performance stack, maximizing maintainability and scalability.

Backend (API)

Component

Technology

Rationale

Language

Go (Golang)

Excellent concurrency, compiled performance, and low resource usage.

Framework

Gin

A fast, minimalist HTTP web framework for building performant REST APIs.

Database

PostgreSQL

Industry standard for transactional systems, ensuring data reliability and complex querying capability.

Architecture

Clean Architecture

Clear separation of concerns (Models, Repositories, Handlers) for LLM-friendly code structure.

Frontend (Client)

Component

Technology

Rationale

Framework

React Native (via Expo)

Enables building native mobile apps (iOS/Android) and a web app from a single codebase.

Language

TypeScript

Adds static typing for better code quality, fewer runtime errors, and improved developer experience.

🏗️ Monorepo Structure

The project utilizes a standard monorepo layout for efficient development and deployment separation:

monybug-monorepo/
├── backend/            # Go/Gin API server
│   ├── cmd/main.go     # Application entry point
│   ├── internal/       # Core business logic (models, repositories, handlers)
│   └── go.mod          
└── frontend/           # React Native / Web client
    ├── src/            # Components, Screens, API clients
    └── package.json    


📚 Project Documentation & Schema

Detailed technical documentation, including the database schema and API specifications, is available in the dedicated docs/ folder.

Document

Description

Link

Database Schema

Detailed SQL definitions for all tables, constraints, and indexes.

View Schema

API Endpoints

Full specification of all available backend endpoints (Routes, Methods, Payloads).

View API Spec

System Flow

High-level data architecture and interaction flow between components.

View Architecture

📊 Entity-Relationship Diagram (ERD)

The core database design is visualized below using the Mermaid diagram standard:

erDiagram
    USERS ||--o{ CATEGORIES : has
    USERS ||--o{ TRANSACTIONS : records

    USERS {
        int user_id PK
        varchar username
        varchar password_hash
        timestamp created_at
    }

    CATEGORIES {
        int category_id PK
        int user_id FK "User who owns the category"
        varchar name "e.g., Groceries, Salary"
        varchar type "EXPENSE or INCOME"
    }

    TRANSACTIONS {
        int transaction_id PK
        int user_id FK "User who made the transaction"
        int category_id FK "Links to the category"
        date transaction_date
        numeric amount "10, 2"
        char currency "3-letter code"
        varchar description
        timestamp created_at
    }

    CATEGORIES ||--o{ TRANSACTIONS : classified_as


⚙️ Getting Started

Follow these instructions to set up and run the MonyBug API and Client locally.

Prerequisites

Ensure the following are installed on your machine:

Go (1.20+)

PostgreSQL (version 12+)

Node.js (LTS) and npm/Yarn

Git

Step 1: Clone the Repository

git clone [https://github.com/yourusername/monybug-monorepo.git](https://github.com/yourusername/monybug-monorepo.git)
cd monybug-monorepo


Step 2: Database Setup

Create Database: Create a new PostgreSQL database named monybug_db.

CREATE DATABASE monybug_db;


Apply Schema: Execute the SQL statements located in docs/db-design/schema.sql against the new database to create the necessary tables.

Step 3: Run the Backend (Go/Gin API)

Navigate to Backend:

cd backend


Configure Environment: Create a file named .env in the backend directory and add your database credentials:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=monybug_db
API_PORT=8080


Run the Server:

go run cmd/main.go
# The API should now be running at http://localhost:8080


Step 4: Run the Frontend (React Native Client)

Navigate to Frontend:

cd ../frontend


Install Dependencies:

npm install
# or yarn install


Start the Development Server:

npx expo start


This will open a browser window with the Expo CLI.

Press w to open the web client in your browser.

Scan the QR code with the Expo Go app on your mobile device to view the native client.

🤝 Contributing

We welcome all contributions! Whether it's reporting a bug, suggesting a feature, or submitting a pull request, please check our Contributing Guidelines (to be created) for more information.

📜 License

This project is licensed under the MIT License. See the LICENSE file for more details.

Developed with Go, Gin, PostgreSQL, and React Native.
