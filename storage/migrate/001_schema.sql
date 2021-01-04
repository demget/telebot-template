-- +goose Up

create table users (
    created_at    timestamp     not null default now(),
    id            bigint(20)    not null primary key,
    lang          varchar(2)    not null default '',
    ref           varchar(64)   not null default ''
);
