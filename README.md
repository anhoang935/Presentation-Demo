# Presentation Demo - SQL & NoSQL Integration

A demo application showcasing the integration of MySQL (SQL) and MongoDB (NoSQL) databases using Go.

## Architecture

- **MySQL**: Stores `Account` and `User` data
- **MongoDB**: Stores `Order` data
- **Constants**: Restaurant and Food data are hardcoded (not in database)

## Database Schema

### MySQL Tables

**Account**
- id (PK)
- email
- password (hashed)

**User**
- id (PK)
- account_id (FK)
- name
- address

### MongoDB Collection

**Orders**
- id (PK)
- account_id
- food_id
- restaurant_id
- total_price
- created_at

## Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- MongoDB 6.0 or higher

## Setup

1. **Clone the repository**

2. **Install dependencies**
```bash
go mod download
```

3. **Setup MySQL Database**
```bash
mysql -u root -p < sql/init.sql
```

4. **Configure environment variables**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

5. **Start MongoDB**
```bash
mongod
```

6. **Run the application**
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Accounts
- `POST /api/accounts` - Create a new account
- `POST /api/accounts/login` - Login
- `GET /api/accounts/{id}` - Get account by ID

### Users
- `POST /api/users` - Create a new user
- `GET /api/users/{id}` - Get user by ID
- `GET /api/users/account/{account_id}` - Get user by account ID
- `PUT /api/users/{id}` - Update user

### Orders
- `POST /api/orders` - Create a new order
- `GET /api/orders/{id}` - Get order by ID
- `GET /api/orders/account/{account_id}` - Get orders by account ID

### Restaurants & Food
- `GET /api/restaurants` - Get all restaurants (constant data)
- `GET /api/restaurants/{id}` - Get restaurant by ID
- `GET /api/foods` - Get all food items (constant data)
- `GET /api/foods/{id}` - Get food by ID

## Example Requests

### Create Account
```bash
curl -X POST http://localhost:8080/api/accounts \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

### Create User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"account_id":1,"name":"John Doe","address":"123 Main St"}'
```

### Create Order
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{"account_id":1,"food_id":1,"restaurant_id":1,"total_price":25.99}'
```

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── database/
│   │   ├── mysql.go          # MySQL connection
│   │   └── mongodb.go        # MongoDB connection
│   ├── models/
│   │   ├── account.go        # Account model
│   │   ├── user.go           # User model
│   │   ├── order.go          # Order model
│   │   ├── restaurant.go     # Restaurant model (constants)
│   │   └── food.go           # Food model (constants)
│   ├── repository/
│   │   ├── account_repo.go   # Account database operations
│   │   ├── user_repo.go      # User database operations
│   │   └── order_repo.go     # Order database operations
│   └── handlers/
│       ├── account.go        # Account HTTP handlers
│       ├── user.go           # User HTTP handlers
│       ├── order.go          # Order HTTP handlers
│       └── static.go         # Restaurant & Food handlers
├── sql/
│   └── init.sql              # MySQL schema
├── .env.example              # Environment variables template
├── .gitignore
├── go.mod
└── README.md
```

## Technologies Used

- **Go** - Programming language
- **MySQL** - Relational database
- **MongoDB** - Document database
- **Gorilla Mux** - HTTP router
- **godotenv** - Environment variables
