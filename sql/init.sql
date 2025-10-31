-- Create database with proper character set
CREATE DATABASE IF NOT EXISTS demo_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE demo_db;

-- Drop tables if they exist (for clean reinstall)
-- Uncomment the lines below if you want to reset the database
-- DROP TABLE IF EXISTS User;
-- DROP TABLE IF EXISTS Account;

-- Account table (stores user authentication data)
CREATE TABLE IF NOT EXISTS Account (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_account_email (email),
    INDEX idx_account_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- User table (stores user profile data)
CREATE TABLE IF NOT EXISTS User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(id) ON DELETE CASCADE,
    INDEX idx_user_account_id (account_id),
    INDEX idx_user_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample data for testing (optional)
-- Uncomment the lines below to add test accounts
-- Note: Password is 'password123' hashed with bcrypt
-- You can generate bcrypt hashes using online tools or the application itself

-- INSERT IGNORE INTO Account (email, password) VALUES
-- ('test@example.com', '$2a$10$XYZ...'), -- Replace with actual bcrypt hash
-- ('demo@example.com', '$2a$10$ABC...'); -- Replace with actual bcrypt hash

-- INSERT IGNORE INTO User (account_id, name, address) VALUES
-- (1, 'Test User', '123 Test Street, Test City'),
-- (2, 'Demo User', '456 Demo Avenue, Demo Town');

-- Verify tables were created successfully
SHOW TABLES;

-- Display table structures
DESCRIBE Account;
DESCRIBE User;
