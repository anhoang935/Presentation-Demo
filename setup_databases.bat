@echo off
REM Database Setup Script for Windows
REM This script helps initialize MySQL and MongoDB databases

echo ========================================
echo Food Ordering System - Database Setup
echo ========================================
echo.

REM Check if .env file exists
if not exist .env (
    echo [ERROR] .env file not found!
    echo Please create a .env file with your database configuration.
    echo You can copy .env.example to .env and update the values.
    pause
    exit /b 1
)

echo [1/4] Loading environment variables from .env...
echo.

REM MySQL Setup
echo ========================================
echo MySQL Database Setup
echo ========================================
echo.
echo This will create the demo_db database and tables in MySQL.
echo Make sure MySQL server is running before continuing.
echo.
set /p MYSQL_SETUP="Do you want to set up MySQL database? (Y/N): "

if /i "%MYSQL_SETUP%"=="Y" (
    echo.
    echo Please enter your MySQL root password when prompted...
    mysql -u root -p < sql\init.sql
    
    if %errorlevel% equ 0 (
        echo [SUCCESS] MySQL database initialized successfully!
    ) else (
        echo [ERROR] Failed to initialize MySQL database.
        echo Please check your MySQL connection and credentials.
    )
) else (
    echo [SKIPPED] MySQL setup skipped.
)

echo.

REM MongoDB Setup
echo ========================================
echo MongoDB Database Setup
echo ========================================
echo.
echo This will create the demo_db database and orders collection in MongoDB.
echo Make sure MongoDB server is running before continuing.
echo.
set /p MONGO_SETUP="Do you want to set up MongoDB database? (Y/N): "

if /i "%MONGO_SETUP%"=="Y" (
    echo.
    echo Initializing MongoDB database...
    mongosh < mongodb\init.js
    
    if %errorlevel% equ 0 (
        echo [SUCCESS] MongoDB database initialized successfully!
    ) else (
        echo [ERROR] Failed to initialize MongoDB database.
        echo Please check your MongoDB connection.
    )
) else (
    echo [SKIPPED] MongoDB setup skipped.
)

echo.
echo ========================================
echo Setup Complete!
echo ========================================
echo.
echo Next steps:
echo 1. Verify .env file has correct database credentials
echo 2. Run 'go run cmd/server/main.go' to start the server
echo 3. Open http://localhost:8080 in your browser
echo.
pause
