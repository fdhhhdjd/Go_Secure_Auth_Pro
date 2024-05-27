CREATE TABLE otps (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    otp_code VARCHAR(6) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);
