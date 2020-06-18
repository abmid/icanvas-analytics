CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email VARCHAR(250) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    name VARCHAR(250) NOT NULL,
    created_at timestamp, 
    updated_at timestamp,
    deleted_at timestamp
)