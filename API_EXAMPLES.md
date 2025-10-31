# API Testing Examples

This file contains example curl commands to test all API endpoints.

## Prerequisites
- MySQL must be running with the database initialized (run sql/init.sql)
- MongoDB must be running
- Server must be running: `go run cmd/server/main.go`

## Account Endpoints

### Create Account
```powershell
curl -X POST http://localhost:8080/api/accounts `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"john@example.com\",\"password\":\"password123\"}'
```

### Get Account
```powershell
curl http://localhost:8080/api/accounts/1
```

### Login
```powershell
curl -X POST http://localhost:8080/api/accounts/login `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"john@example.com\",\"password\":\"password123\"}'
```

## User Endpoints

### Create User
```powershell
curl -X POST http://localhost:8080/api/users `
  -H "Content-Type: application/json" `
  -d '{\"account_id\":1,\"name\":\"John Doe\",\"address\":\"123 Main St\"}'
```

### Get User by ID
```powershell
curl http://localhost:8080/api/users/1
```

### Get User by Account ID
```powershell
curl http://localhost:8080/api/users/account/1
```

### Update User
```powershell
curl -X PUT http://localhost:8080/api/users/1 `
  -H "Content-Type: application/json" `
  -d '{\"name\":\"John Smith\",\"address\":\"456 Oak Ave\"}'
```

## Restaurant Endpoints (Static Data)

### Get All Restaurants
```powershell
curl http://localhost:8080/api/restaurants
```

### Get Restaurant by ID
```powershell
curl http://localhost:8080/api/restaurants/1
```

### Get Foods by Restaurant
```powershell
curl http://localhost:8080/api/restaurants/1/foods
```

## Food Endpoints (Static Data)

### Get All Foods
```powershell
curl http://localhost:8080/api/foods
```

### Get Food by ID
```powershell
curl http://localhost:8080/api/foods/1
```

## Order Endpoints

### Create Order
```powershell
curl -X POST http://localhost:8080/api/orders `
  -H "Content-Type: application/json" `
  -d '{\"account_id\":1,\"food_id\":1,\"restaurant_id\":1,\"total_price\":12.99}'
```

### Get Order by ID
```powershell
curl http://localhost:8080/api/orders/[MONGODB_OBJECT_ID]
```

### Get Orders by Account ID
```powershell
curl http://localhost:8080/api/orders/account/1
```

### Get All Orders
```powershell
curl http://localhost:8080/api/orders
```

## Health Check
```powershell
curl http://localhost:8080/health
```

## Complete Flow Example

```powershell
# 1. Create an account
curl -X POST http://localhost:8080/api/accounts `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"demo@example.com\",\"password\":\"demo123\"}'

# 2. Login
curl -X POST http://localhost:8080/api/accounts/login `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"demo@example.com\",\"password\":\"demo123\"}'

# 3. Create user profile
curl -X POST http://localhost:8080/api/users `
  -H "Content-Type: application/json" `
  -d '{\"account_id\":1,\"name\":\"Demo User\",\"address\":\"123 Demo St\"}'

# 4. View available restaurants
curl http://localhost:8080/api/restaurants

# 5. View foods at a restaurant
curl http://localhost:8080/api/restaurants/1/foods

# 6. Place an order
curl -X POST http://localhost:8080/api/orders `
  -H "Content-Type: application/json" `
  -d '{\"account_id\":1,\"food_id\":1,\"restaurant_id\":1,\"total_price\":12.99}'

# 7. View order history
curl http://localhost:8080/api/orders/account/1
```
