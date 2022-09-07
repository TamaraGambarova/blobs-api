-- +migrate Up

create table blob
(
    id      bigserial not null primary key,
    owner   text      not null,
    content jsonb      not null
);

-- +migrate Down
drop table blob cascade;
