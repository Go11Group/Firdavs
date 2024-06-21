CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    post_id INT REFERENCES posts(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);