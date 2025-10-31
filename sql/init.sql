-- Create database
CREATE DATABASE IF NOT EXISTS demo_db;
USE demo_db;

-- Account table
CREATE TABLE IF NOT EXISTS Account (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- User table
CREATE TABLE IF NOT EXISTS User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(id) ON DELETE CASCADE
);

-- Create indexes for better performance
CREATE INDEX idx_user_account_id ON User(account_id);
CREATE INDEX idx_account_email ON Account(email);
