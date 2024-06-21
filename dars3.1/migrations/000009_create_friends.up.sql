CREATE TABLE friends (
    user_id INT REFERENCES users(id),
    friend_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_id)
);