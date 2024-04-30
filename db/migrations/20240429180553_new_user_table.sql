-- +goose Up
-- +goose StatementBegin
create table users(
    id bigserial primary key,
    name varchar(100) unique,
    email text unique,
    pass text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
