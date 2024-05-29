CREATE TABLE password_history (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id),
    old_password VARCHAR(150) NOT NULL UNIQUE,
    reason_status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);