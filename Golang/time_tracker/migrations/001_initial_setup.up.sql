-- 001_initial_setup.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_number VARCHAR(10) NOT NULL,
    name VARCHAR(100),
    surname VARCHAR(100),
    patronymic VARCHAR(100),
    address VARCHAR(255)
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP
);
