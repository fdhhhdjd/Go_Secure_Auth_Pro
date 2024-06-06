CREATE TABLE social_logins (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    provider SMALLINT  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);