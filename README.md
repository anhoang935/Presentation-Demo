# Presentation Demo - SQL & NoSQL Integration

A demo application showcasing the integration of MySQL (SQL) and MongoDB (NoSQL) databases using Go, with a modern web frontend.

## 🎯 Quick Links

- 🚀 **[Database Setup Summary](./DATABASE_SETUP_SUMMARY.md)** - Start here for database setup!
- 📗 **[MySQL Setup Guide](./MYSQL_SETUP.md)** - MySQL Workbench instructions
- 📙 **[MongoDB Setup Guide](./MONGODB_SETUP.md)** - MongoDB Compass instructions
- 🎨 **[Visual Guide](./VISUAL_GUIDE.md)** - Diagrams and flowcharts
- 🌐 **[Web Frontend Guide](./web/README.md)** - Frontend documentation
- ⚡ **[Quick Start](./web/QUICKSTART.md)** - 5-minute setup guide

## Architecture

- **MySQL**: Stores `Account` and `User` data (relational, ACID-compliant)
- **MongoDB**: Stores `Order` data (document-based, scalable)
- **Constants**: Restaurant and Food data are hardcoded (not in database)
- **Web Frontend**: HTML/CSS/JavaScript interface for user interaction

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
- MySQL 8.0 or higher (with MySQL Workbench recommended)
- MongoDB 6.0 or higher (with MongoDB Compass recommended)

## 🚀 Quick Setup

### Prerequisites Installation

1. **Install MySQL** (if not already installed)
   - Download from: https://dev.mysql.com/downloads/
   - Or use package manager:
     ```powershell
     # Windows with Chocolatey
     choco install mysql
     
     # macOS with Homebrew
     brew install mysql
     ```

2. **Install MongoDB** (if not already installed)
   - Download from: https://www.mongodb.com/try/download/community
   - Or use package manager:
     ```powershell
     # Windows with Chocolatey
     choco install mongodb
     
     # macOS with Homebrew
     brew install mongodb-community
     ```

3. **Start Database Services**
   ```powershell
   # Windows
   net start MySQL80
   net start MongoDB
   
   # macOS/Linux
   brew services start mysql
   brew services start mongodb-community
   ```

### Automated Database Setup (Recommended)

**Windows:**
```powershell
.\setup_databases.bat
```

**Linux/macOS:**
```bash
chmod +x setup_databases.sh
./setup_databases.sh
```

This script will:
- ✅ Initialize MySQL database and tables
- ✅ Initialize MongoDB database and collections
- ✅ Create necessary indexes
- ✅ Verify database connections

### Manual Database Setup

**Detailed guides available:**
- 📘 [Complete Database Setup Guide](./DATABASE_SETUP.md) - Comprehensive database documentation

**Quick version:**

1. **Setup MySQL Database**
   ```powershell
   mysql -u root -p < sql\init.sql
   ```

2. **Setup MongoDB Database**
   ```powershell
   mongosh < mongodb\init.js
   ```

3. **Configure environment variables**
   
   Copy `.env.example` to `.env` and update with your settings:
   ```env
   # MySQL Configuration
   MYSQL_HOST=localhost
   MYSQL_PORT=3306
   MYSQL_USER=root
   MYSQL_PASSWORD=your_password
   MYSQL_DATABASE=demo_db

   # MongoDB Configuration
   MONGODB_URI=mongodb://localhost:27017
   MONGODB_DATABASE=demo_db

   # Server Configuration
   PORT=8080
   ```

4. **Install Go dependencies**
   ```powershell
   go mod tidy
   ```

5. **Run the application**
   ```powershell
   go run cmd/server/main.go
   ```

   You should see:
   ```
   ✅ MySQL connected successfully
   ✅ MongoDB connected successfully
   🚀 Server starting on port 8080
   📝 API documentation available at http://localhost:8080/api
   🌐 Web interface available at http://localhost:8080
   ```

6. **Access the Web Interface**
   
   Open your browser and navigate to `http://localhost:8080`

### Troubleshooting

If you encounter database connection errors:

1. **MySQL Connection Failed**
   ```powershell
   # Test MySQL connection
   mysql -h localhost -P 3306 -u root -p -e "SELECT VERSION();"
   
   # Verify MySQL is running
   net start | findstr MySQL
   ```

2. **MongoDB Connection Failed**
   ```powershell
   # Test MongoDB connection
   mongosh --eval "db.version()"
   
   # Verify MongoDB is running
   net start | findstr MongoDB
   ```

3. **Check DATABASE_SETUP.md** for detailed troubleshooting steps

📚 **More Resources:**
- [Database Setup Guide](./DATABASE_SETUP.md) - Detailed database configuration
- [API Examples](./API_EXAMPLES.md) - Test API endpoints

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
├── web/
│   ├── index.html            # Web frontend HTML
│   ├── styles.css            # Frontend styling
│   ├── app.js                # Frontend JavaScript
│   └── README.md             # Frontend documentation
├── sql/
│   └── init.sql              # MySQL schema
├── .env.example              # Environment variables template
├── .gitignore
├── go.mod
└── README.md
```

## Technologies Used

### Backend
- **Go** - Programming language
- **MySQL** - Relational database
- **MongoDB** - Document database
- **Gorilla Mux** - HTTP router
- **godotenv** - Environment variables

### Frontend
- **HTML5** - Structure
- **CSS3** - Styling (Grid, Flexbox, Animations)
- **JavaScript (ES6+)** - Logic and API communication
- **Fetch API** - HTTP requests

## Features

- ✅ User registration and login
- ✅ View restaurants and menus
- ✅ Place food orders
- ✅ View order history
- ✅ Update user profile
- ✅ Responsive web design
- ✅ Toast notifications
- ✅ Session persistence
- ✅ CORS enabled
- ✅ Database integration indicators (MySQL/MongoDB badges)
