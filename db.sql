CREATE DATABASE api_kostless;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR(50),
    username VARCHAR(50),
    password VARCHAR(100),
    email VARCHAR(50),
    phone_number VARCHAR(16),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
);

CREATE TABLE seekers(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50),
    password VARCHAR(100),
    fullname VARCHAR(50),
    email VARCHAR(50),
    phone_number VARCHAR(16),
    attitude_points INT,
    status VARCHAR(50),
    room_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

CREATE TABLE kos(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID,
    name VARCHAR(50),
    address TEXT,
    room_count INT,
    coordinate TEXT,
    desc TEXT,
    rules TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE TABLE rooms(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50),
    type VARCHAR(50),
    desc TEXT,
    avail ENUM('open', 'occupied') DEFAULT 'open',
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE bookings(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    room_id UUID,
    seeker_id UUID,
    start_date DATE,
    end_date DATE,
    discount INT,
    total INT,
    pay_later BOOLEAN DEFAULT FALSE,
    due_date DATE,
    payment_status ENUM('pending', 'paid', 'overdue') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (seeker_id) REFERENCES seekers(id)
);

CREATE TABLE vouchers(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50),
    percent_amount INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);