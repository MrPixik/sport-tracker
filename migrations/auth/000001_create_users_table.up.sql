CREATE TABLE users
(
    id SERIAL PRIMARY KEY ,
    login   VARCHAR(256) UNIQUE not null ,
    password_hash varchar(256) not null
);