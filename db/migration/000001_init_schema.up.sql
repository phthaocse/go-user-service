CREATE TABLE IF NOT EXISTS user (
    _id INT4 AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name  VARCHAR(30),
    last_name VARCHAR(30),
    password VARCHAR(255),
    is_admin BOOLEAN DEFAULT FALSE,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);