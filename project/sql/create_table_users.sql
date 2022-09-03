-- Create type for enum
CREATE TYPE roles AS ENUM ('admin', 'seller', 'buyyer');

-- Create table users if not exists
CREATE TABLE IF NOT EXISTS users(
    id  uuid,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    province VARCHAR(255) NOT NULL,
    mobile VARCHAR(20) NOT NULL,
    password CHAR(225) NOT NULL,
    role roles NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (id)
);