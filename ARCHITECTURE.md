# Architecture Overview

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        Client / API                          │
│                    (HTTP REST Requests)                      │
└──────────────────────┬──────────────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────────────┐
│                    HTTP Server (Gorilla Mux)                 │
│  ┌────────────────────────────────────────────────────────┐ │
│  │                    Middleware Layer                     │ │
│  │  - Logging Middleware                                   │ │
│  │  - CORS Middleware                                      │ │
│  └────────────────────────────────────────────────────────┘ │
└──────────────────────┬──────────────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────────────┐
│                      Handlers Layer                          │
│  ┌──────────────┬──────────────┬──────────────┬──────────┐ │
│  │   Account    │     User     │    Order     │  Static  │ │
│  │   Handler    │   Handler    │   Handler    │ Handler  │ │
│  └──────────────┴──────────────┴──────────────┴──────────┘ │
└──────────────────────┬──────────────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────────────┐
│                    Repository Layer                          │
│  ┌──────────────┬──────────────┬──────────────────────────┐ │
│  │   Account    │     User     │        Order             │ │
│  │  Repository  │  Repository  │      Repository          │ │
│  └──────────────┴──────────────┴──────────────────────────┘ │
└──────────────────────┬──────────────────────────────────────┘
                       │
            ┌──────────┴──────────┐
            │                     │
┌───────────▼───────────┐ ┌──────▼──────────────┐
│   MySQL Database      │ │  MongoDB Database   │
│  ┌─────────────────┐  │ │  ┌──────────────┐  │
│  │  Account Table  │  │ │  │   Orders     │  │
│  │  - id           │  │ │  │  Collection  │  │
│  │  - email        │  │ │  │  - _id       │  │
│  │  - password     │  │ │  │  - account_id│  │
│  └─────────────────┘  │ │  │  - food_id   │  │
│  ┌─────────────────┐  │ │  │  - rest_id   │  │
│  │   User Table    │  │ │  │  - price     │  │
│  │  - id           │  │ │  │  - created_at│  │
│  │  - account_id   │  │ │  └──────────────┘  │
│  │  - name         │  │ │                     │
│  │  - address      │  │ │                     │
│  └─────────────────┘  │ │                     │
└───────────────────────┘ └─────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│              Constant Data (In-Memory)                       │
│  ┌──────────────────────┐  ┌──────────────────────────────┐│
│  │    Restaurants       │  │         Food Items           ││
│  │  - Pizza Palace      │  │  - Margherita Pizza          ││
│  │  - Sushi World       │  │  - Pepperoni Pizza           ││
│  │  - Burger House      │  │  - California Roll           ││
│  │  - Pasta Paradise    │  │  - Classic Burger            ││
│  │  - Taco Town         │  │  - Spaghetti Carbonara       ││
│  └──────────────────────┘  └──────────────────────────────┘│
└─────────────────────────────────────────────────────────────┘
```

## Data Flow

### Account Creation Flow
```
Client → Handler → Repository → MySQL → Response
```

### Order Creation Flow
```
Client → Handler → Validation (Restaurant & Food) → Repository → MongoDB → Response
```

### Get Restaurants Flow
```
Client → Handler → Static Data (Models) → Response
```

## Technology Stack

- **Language**: Go 1.21+
- **Web Framework**: Gorilla Mux (HTTP Router)
- **SQL Database**: MySQL 8.0+
- **NoSQL Database**: MongoDB 6.0+
- **Dependencies**:
  - `github.com/gorilla/mux` - HTTP router
  - `github.com/go-sql-driver/mysql` - MySQL driver
  - `go.mongodb.org/mongo-driver` - MongoDB driver
  - `github.com/joho/godotenv` - Environment variables
  - `golang.org/x/crypto` - Password hashing

## Design Patterns

1. **Repository Pattern**: Separates data access logic from business logic
2. **Handler Pattern**: Separates HTTP concerns from business logic
3. **Layered Architecture**: Clear separation between presentation, business, and data layers
4. **Dependency Injection**: Handlers receive repository instances

## Key Features

✅ **Dual Database Integration**: MySQL for structured data, MongoDB for flexible data
✅ **RESTful API**: Clean, resource-based endpoints
✅ **Password Hashing**: Secure password storage using bcrypt
✅ **CORS Support**: Cross-origin resource sharing enabled
✅ **Request Logging**: All requests are logged for debugging
✅ **Error Handling**: Consistent error responses
✅ **Data Validation**: Input validation at handler level
✅ **Graceful Shutdown**: Proper cleanup on server shutdown

## Database Strategy

### Why MySQL for Account & User?
- Structured data with clear relationships
- ACID compliance for account operations
- Foreign key constraints ensure data integrity
- Efficient joins for user-account queries

### Why MongoDB for Orders?
- High write throughput for order creation
- Flexible schema for order variations
- Easy to scale horizontally
- No complex relationships needed

### Why Static Data for Restaurants & Food?
- Demo/presentation purposes
- Fast access (no DB queries)
- Easy to modify for demos
- Can be migrated to DB later if needed
