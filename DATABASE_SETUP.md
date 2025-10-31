# Database Setup Guide

This document provides detailed instructions for setting up MySQL and MongoDB databases for the Food Ordering System.

## Prerequisites

- **MySQL 5.7+** or **MariaDB 10.3+**
- **MongoDB 4.0+**
- MySQL client (mysql command-line tool)
- MongoDB client (mongosh or mongo shell)

## MySQL Setup

### 1. Start MySQL Server

Make sure your MySQL server is running:

**Windows:**
```bash
# Start MySQL service
net start MySQL80  # Adjust service name based on your installation
```

**Linux/macOS:**
```bash
sudo systemctl start mysql
# or
sudo service mysql start
```

### 2. Create Database and Tables

Run the initialization script:

```bash
mysql -u root -p < sql/init.sql
```

Or connect to MySQL and run it manually:

```bash
mysql -u root -p
```

Then in the MySQL shell:
```sql
source sql/init.sql;
```

### 3. Verify MySQL Setup

Check that tables were created:

```sql
USE demo_db;
SHOW TABLES;
DESCRIBE Account;
DESCRIBE User;
```

Expected output:
- Account table with columns: id, email, password, created_at, updated_at
- User table with columns: id, account_id, name, address, created_at, updated_at

### 4. Configure MySQL Connection

Update the `.env` file with your MySQL credentials:

```env
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=demo_db
```

## MongoDB Setup

### 1. Start MongoDB Server

Make sure your MongoDB server is running:

**Windows:**
```bash
# Start MongoDB service
net start MongoDB
```

**Linux/macOS:**
```bash
sudo systemctl start mongod
# or
sudo service mongod start
```

### 2. Initialize MongoDB Database

Run the initialization script:

```bash
mongosh < mongodb/init.js
```

Or connect to MongoDB and run it manually:

```bash
mongosh
```

Then in the MongoDB shell:
```javascript
load('mongodb/init.js');
```

### 3. Verify MongoDB Setup

Check that the database and collection were created:

```javascript
use demo_db
show collections
db.orders.getIndexes()
```

Expected output:
- Collection: orders
- Indexes: _id, account_id, restaurant_id, created_at, compound index

### 4. Configure MongoDB Connection

Update the `.env` file with your MongoDB connection string:

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=demo_db
```

If using MongoDB with authentication:
```env
MONGODB_URI=mongodb://username:password@localhost:27017
MONGODB_DATABASE=demo_db
```

## Database Architecture

### MySQL (Relational Data)

Stores structured user data that requires ACID properties:

- **Account Table**: User authentication credentials
  - Primary Key: id (auto-increment)
  - Unique Key: email
  - Password stored as bcrypt hash

- **User Table**: User profile information
  - Primary Key: id (auto-increment)
  - Foreign Key: account_id → Account(id)
  - CASCADE delete when account is deleted

### MongoDB (Document Store)

Stores order data that benefits from flexible schema and scalability:

- **orders Collection**: Order history
  - Document ID: _id (ObjectId)
  - Fields: account_id, food_id, restaurant_id, total_price, created_at
  - Indexed by: account_id, restaurant_id, created_at

### Mock Data (In-Memory)

Static data loaded from Go code:

- **Restaurants**: 5 pre-defined restaurants
- **Foods**: 10 pre-defined menu items

## Troubleshooting

### MySQL Connection Issues

**Error: "Access denied for user"**
```bash
# Reset MySQL root password or create a new user
mysql -u root -p
CREATE USER 'demo_user'@'localhost' IDENTIFIED BY 'demo_password';
GRANT ALL PRIVILEGES ON demo_db.* TO 'demo_user'@'localhost';
FLUSH PRIVILEGES;
```

**Error: "Can't connect to MySQL server"**
- Check if MySQL service is running
- Verify port 3306 is not blocked by firewall
- Check hostname in .env file (use 127.0.0.1 instead of localhost if needed)

### MongoDB Connection Issues

**Error: "MongoServerError: Authentication failed"**
- MongoDB might be running without authentication
- Update MONGODB_URI to: `mongodb://localhost:27017` (no username/password)

**Error: "connect ECONNREFUSED"**
- Check if MongoDB service is running
- Verify port 27017 is not blocked by firewall
- Check MongoDB logs for errors

### Verification Commands

Test MySQL connection:
```bash
mysql -h localhost -P 3306 -u root -p demo_db -e "SELECT VERSION();"
```

Test MongoDB connection:
```bash
mongosh --eval "db.version()"
```

## Sample Data

To add test data for development, see the commented sections in:
- `sql/init.sql` - Sample accounts and users
- `mongodb/init.js` - Sample orders

## Database Migrations

When making schema changes:

1. **MySQL**: Update `sql/init.sql` with new columns/tables
2. **MongoDB**: Update validation schema in `mongodb/init.js`
3. Update corresponding Go models in `internal/models/`
4. Update repositories in `internal/repository/`

## Performance Tuning

### MySQL Indexes
Current indexes:
- Account: email, created_at
- User: account_id, created_at

### MongoDB Indexes
Current indexes:
- account_id (ascending)
- restaurant_id (ascending)
- created_at (descending)
- account_id + created_at (compound)

Add more indexes based on query patterns:
```javascript
db.orders.createIndex({ "food_id": 1 });
db.orders.createIndex({ "total_price": -1 });
```

## Backup and Restore

### MySQL Backup
```bash
mysqldump -u root -p demo_db > backup_mysql.sql
```

### MySQL Restore
```bash
mysql -u root -p demo_db < backup_mysql.sql
```

### MongoDB Backup
```bash
mongodump --db demo_db --out ./backup_mongodb
```

### MongoDB Restore
```bash
mongorestore --db demo_db ./backup_mongodb/demo_db
```

## Security Best Practices

1. **Never commit .env file** with real credentials
2. Use strong passwords for database users
3. Limit database user privileges (don't use root in production)
4. Enable SSL/TLS for database connections in production
5. Regularly backup your databases
6. Use prepared statements (already implemented in the code)
7. Hash passwords with bcrypt (already implemented)
