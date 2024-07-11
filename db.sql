CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR(50),
    username VARCHAR(50),
    password VARCHAR(100),
    email VARCHAR(50),
    phone_number VARCHAR(16),
    photo_profile VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kos(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50),
    address TEXT,
    room_count INT,
    coordinate TEXT,
    description TEXT,
    rules TEXT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS rooms(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50),
    type VARCHAR(50),
    description TEXT,
    avail VARCHAR(15),
    price INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    kos_id UUID REFERENCES kos(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS seekers(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50),
    password VARCHAR(100),
    fullname VARCHAR(50),
    email VARCHAR(50),
    phone_number VARCHAR(16),
    attitude_points INT,
    status VARCHAR(50),
    photo_profile VARCHAR(255),
    room_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

CREATE TABLE IF NOT EXISTS bookings(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    room_id UUID,
    seeker_id UUID,
    start_date DATE,
    end_date DATE,
    discount INT,
    total INT,
    pay_later BOOLEAN DEFAULT FALSE,
    due_date DATE,
    payment_status VARCHAR(15),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (seeker_id) REFERENCES seekers(id)
);

CREATE TABLE IF NOT EXISTS vouchers(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50),
    expired_date DATE,
    seeker_id UUID,
    percent_amount INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (seeker_id) REFERENCES seekers(id)
);