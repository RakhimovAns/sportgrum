CREATE TABLE users(
    id bigserial primary key,
    name text not null,
    nickname text not null ,
    password text not null,
    created timestamp not null default current_timestamp
);