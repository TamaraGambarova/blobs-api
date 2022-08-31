-- +migrate Up

create table blob
(
    id      bigserial not null primary key,
    owner   text      not null,
    content text      not null
);

-- +migrate Down
drop table blob cascade;
