CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title TEXT,
    body TEXT,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);