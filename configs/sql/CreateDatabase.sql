CREATE DATABASE hussat;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE users;

CREATE TABLE users (
    uid UUID NOT NULL PRIMARY KEY,
    id VARCHAR(10) UNIQUE,
    name VARCHAR(50) NOT NULL,
    family_name VARCHAR(50) NOT NULL,
    given_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    date_of_birth DATE,
    religion VARCHAR(30),
    ethnic_group VARCHAR(30),
    address VARCHAR(100),
    resident VARCHAR(100),
    faculty VARCHAR(50),
    class VARCHAR(50),
    academic_year INT,
    department VARCHAR(30),
    position VARCHAR(30),
    role VARCHAR(10) NOT NULL,
    party_member VARCHAR(10),
    avatar VARCHAR(150),
    signature VARCHAR(150),
    created_at DATE NOT NULL,
    updated_at DATE,
    verified_at DATE,
    approved_at DATE,
    deleted_at DATE,
    blocked_at DATE
);

INSERT INTO users (
    uid, name, given_name, family_name, email,
    date_of_birth, religion, ethnic_group,
    address, resident, faculty, class,
    department, position, role, party_member,
    avatar, signature, created_at, updated_at,
    verified_at, approved_at
)
VALUES (
    uuid_generate_v4(),
    'Trần Đăng Quang Minh (Help Web)',
    'Dragon',
    'Fury',
    'kuanmin.bigdragon.56@gmail.com',
    '2005-06-06',
    'Không',
    'Kinh',
    'Xã Mão Điền, Huyện Thuận Thành, Tỉnh Bắc Ninh',
    '105B, Tổ dân phố số 1, Xã Vĩnh Quỳnh, Huyện Thanh Trì, TP. Hà Nội',
    'toan-co-tinhoc',
    'K68 Toán Tin',
    'hanhchinh-ketoan',
    'thanhvien',
    'root',
    'doanvien',
    'https://res.cloudinary.com/dc77voqvd/image/upload/v1735388376/maychenang/36fc1f83bb921884_b8k7ea.png',
    NULL,
    '2024-03-21',
    '2024-06-11',
    '2024-12-28',
    '2024-03-21'
);
