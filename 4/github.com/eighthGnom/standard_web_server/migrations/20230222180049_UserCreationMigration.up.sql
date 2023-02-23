create table users (
    id integer not null primary key,
    user_name varchar not null,
    emil varchar not null  unique,
    age integer not null,
    password_hash varchar not null
);