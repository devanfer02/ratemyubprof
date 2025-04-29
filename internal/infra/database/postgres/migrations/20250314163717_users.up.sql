CREATE TABLE users (
    id VARCHAR(250) PRIMARY KEY,
    nim VARCHAR(50) CONSTRAINT users_nim_unique UNIQUE NOT NULL,
    username VARCHAR(200) CONSTRAINT users_username_unique UNIQUE NOT NULL,
    password VARCHAR(250) NOT NULL,
    forgot_password_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
