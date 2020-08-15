CREATE TABLE IF NOT EXISTS settings(
    id serial PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    category VARCHAR(250) NOT NULL,
    value TEXT,
    created_at timestamp, 
    updated_at timestamp
)
-- #ex
-- name     |   category    |    value
-- url      |   canvas_api  |   canvas.com
-- token    |   canvas_api  |   super-secret-token
-- cron     |   job         |   ********