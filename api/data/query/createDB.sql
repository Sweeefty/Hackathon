-- createDB.sql

-- Create campus table
CREATE TABLE IF NOT EXISTS campus (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- Create bde table
CREATE TABLE IF NOT EXISTS bde (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    campus_id INTEGER NOT NULL,
    FOREIGN KEY (campus_id) REFERENCES campus(id)
);

-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    SessionCookie TEXT,
    campus_id INTEGER NOT NULL,
    bde_id INTEGER,
    FOREIGN KEY (bde_id) REFERENCES bde(id),
    FOREIGN KEY (campus_id) REFERENCES campus(id)
);

-- Create events table
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    date TEXT NOT NULL,
    price INTEGER NOT NULL,
    bde_id INTEGER NOT NULL,
    FOREIGN KEY (bde_id) REFERENCES bde(id)
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    price INTEGER NOT NULL
);

-- Create inventory table
CREATE TABLE IF NOT EXISTS inventory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    accounts_id INTEGER NOT NULL,
    products_id INTEGER NOT NULL,
    FOREIGN KEY (accounts_id) REFERENCES accounts(id),
    FOREIGN KEY (products_id) REFERENCES products(id)
);

