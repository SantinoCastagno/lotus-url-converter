CREATE TABLE translations (
    id SERIAL PRIMARY KEY,
    source VARCHAR(100) UNIQUE NOT NULL,
    destination VARCHAR(100)
);
