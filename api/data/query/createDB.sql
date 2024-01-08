-- create_accounts_table.sql

-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user'
);

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

-- insert_users.sql

-- Insert sample users
INSERT INTO accounts (email, password , role) VALUES ('gurvan', 'gurvan' , 'admin');
INSERT INTO accounts (email, password , role ) VALUES ('lenny', 'lenny' , 'admin');
INSERT INTO accounts (email, password , role ) VALUES ('rayen', 'rayen' , 'admin');

-- Insert sample campus
INSERT INTO campus (name) VALUES ('Nantes');
INSERT INTO campus (name) VALUES ('Paris');
INSERT INTO campus (name) VALUES ('Montcuq');

-- Insert sample bde
INSERT INTO bde (name, campus_id) VALUES ('Peacocktail', 1);
INSERT INTO bde (name, campus_id) VALUES ('ParisBDE', 2);
INSERT INTO bde (name, campus_id) VALUES ('MontcuqBDE', 3);

-- Insert sample events
INSERT INTO events (name, description, date, price, bde_id) VALUES ('Soirée de fin année', '', '2024-05-23', 10, 1);