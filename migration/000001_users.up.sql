CREATE TABLE IF NOT EXISTS users (
    id            serial       not null unique,
    login          varchar(255) not null unique,
    email      varchar(255) not null unique,
    password_hash varchar(255) not null,
    phone_number varchar(255) not null
);