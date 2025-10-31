#!/bin/bash
# Database Setup Script for Linux/macOS
# This script helps initialize MySQL and MongoDB databases

echo "========================================"
echo "Food Ordering System - Database Setup"
echo "========================================"
echo ""

# Check if .env file exists
if [ ! -f .env ]; then
    echo "[ERROR] .env file not found!"
    echo "Please create a .env file with your database configuration."
    echo "You can copy .env.example to .env and update the values."
    exit 1
fi

echo "[1/4] Loading environment variables from .env..."
source .env
echo ""

# MySQL Setup
echo "========================================"
echo "MySQL Database Setup"
echo "========================================"
echo ""
echo "This will create the demo_db database and tables in MySQL."
echo "Make sure MySQL server is running before continuing."
echo ""
read -p "Do you want to set up MySQL database? (y/n): " MYSQL_SETUP

if [ "$MYSQL_SETUP" = "y" ] || [ "$MYSQL_SETUP" = "Y" ]; then
    echo ""
    echo "Please enter your MySQL password when prompted..."
    mysql -u "${MYSQL_USER:-root}" -p < sql/init.sql
    
    if [ $? -eq 0 ]; then
        echo "[SUCCESS] MySQL database initialized successfully!"
    else
        echo "[ERROR] Failed to initialize MySQL database."
        echo "Please check your MySQL connection and credentials."
    fi
else
    echo "[SKIPPED] MySQL setup skipped."
fi

echo ""

# MongoDB Setup
echo "========================================"
echo "MongoDB Database Setup"
echo "========================================"
echo ""
echo "This will create the demo_db database and orders collection in MongoDB."
echo "Make sure MongoDB server is running before continuing."
echo ""
read -p "Do you want to set up MongoDB database? (y/n): " MONGO_SETUP

if [ "$MONGO_SETUP" = "y" ] || [ "$MONGO_SETUP" = "Y" ]; then
    echo ""
    echo "Initializing MongoDB database..."
    mongosh < mongodb/init.js
    
    if [ $? -eq 0 ]; then
        echo "[SUCCESS] MongoDB database initialized successfully!"
    else
        echo "[ERROR] Failed to initialize MongoDB database."
        echo "Please check your MongoDB connection."
    fi
else
    echo "[SKIPPED] MongoDB setup skipped."
fi

echo ""
echo "========================================"
echo "Setup Complete!"
echo "========================================"
echo ""
echo "Next steps:"
echo "1. Verify .env file has correct database credentials"
echo "2. Run 'go run cmd/server/main.go' to start the server"
echo "3. Open http://localhost:8080 in your browser"
echo ""
