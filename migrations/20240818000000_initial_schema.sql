-- Table for houses
CREATE TABLE houses (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    developer VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_flat_added TIMESTAMP
);

-- Table for flats
CREATE TABLE flats (
    id SERIAL PRIMARY KEY,
    house_id INT NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    price INT NOT NULL CHECK (price >= 0),
    rooms INT NOT NULL CHECK (rooms > 0),
    status VARCHAR(50) NOT NULL CHECK (status IN ('created', 'approved', 'declined', 'on moderation')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    user_type VARCHAR(50) NOT NULL CHECK (user_type IN ('client', 'moderator')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
