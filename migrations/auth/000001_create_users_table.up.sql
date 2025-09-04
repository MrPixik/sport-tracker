CREATE TABLE users
(
    id SERIAL PRIMARY KEY ,
    login   VARCHAR(256) UNIQUE not null unique ,
    password_hash varchar(256) not null
);