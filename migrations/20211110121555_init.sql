-- +goose Up
-- +goose StatementBegin
create table book(
    id bigserial primary key,
    name text not null,
    published_year int,
    description text not null default '',
    price int not null,
    image bytea not null
);

create table author(
    id bigserial primary key,
    name text not null,
    surname text not null,
    middle_name text not null,
    birth_year int,
    description text not null default ''
);

create table book_author(
  book_id bigint not null ,
  author_id bigint not null,
  primary key (book_id, author_id)
);

create table cart(
  id bigserial primary key,
  session_id text not null unique,
  created_time timestamp not null
);

create table cart_book(
    cart_id bigint not null,
    book_id bigint not null,
    count int not null default 0,
    primary key (book_id, cart_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table book;
drop table author;
drop table book_author;
drop table cart;
drop table cart_book;
-- +goose StatementEnd
