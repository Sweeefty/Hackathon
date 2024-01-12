-- createAuthorization.sql

-- Create campus table
CREATE TABLE IF NOT EXISTS authorization (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    hash_key TEXT NOT NULL UNIQUE,
    date DATE NOT NULL
);

INSERT INTO authorization (hash_key,date) VALUES ('$2a$10$VpWUA16JVEg1EvVei8JvRuwyj8ZNX2M.Gpu7b0qw0/x6oeU7cEdQu', '2020-01-01');

