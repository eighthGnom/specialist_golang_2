create table users (
    user_id bigserial not null primary key,
    login varchar not null unique,
    password varchar not null
);

create table articles (
    article_id bigserial not null primary key,
    title varchar not null unique,
    author varchar not null,
    content varchar not null
);