# Quick Start Guide

## 1. Install Dependencies

### MySQL
Download and install from: https://dev.mysql.com/downloads/mysql/

### MongoDB
Download and install from: https://www.mongodb.com/try/download/community

### Go
Download and install from: https://go.dev/download/

## 2. Setup Databases

### MySQL Setup
```powershell
# Start MySQL service (if not running)
net start MySQL80

# Create database and tables
mysql -u root -p < sql/init.sql
```

### MongoDB Setup
```powershell
# Start MongoDB service
net start MongoDB

# MongoDB will automatically create the database when first used
```

## 3. Configure Environment

```powershell
# Copy the example environment file
copy .env.example .env

# Edit .env with your database credentials
notepad .env
```

## 4. Install Go Dependencies

```powershell
go mod download
```

## 5. Run the Application

```powershell
go run cmd/server/main.go
```

Or build and run:

```powershell
go build -o server.exe ./cmd/server
.\server.exe
```

## 6. Test the API

The server will be available at `http://localhost:8080`

### Quick Test
```powershell
curl http://localhost:8080/health
curl http://localhost:8080/api/restaurants
```

See `API_EXAMPLES.md` for complete API documentation and examples.

## Troubleshooting

### MySQL Connection Error
- Check if MySQL service is running: `net start MySQL80`
- Verify credentials in `.env` file
- Ensure database exists: `mysql -u root -p -e "SHOW DATABASES;"`

### MongoDB Connection Error
- Check if MongoDB service is running: `net start MongoDB`
- Verify MongoDB URI in `.env` file
- Default URI: `mongodb://localhost:27017`

### Port Already in Use
- Change the PORT in `.env` file
- Or stop the process using port 8080

### Build Errors
- Run `go mod tidy` to sync dependencies
- Ensure Go version is 1.21 or higher: `go version`
