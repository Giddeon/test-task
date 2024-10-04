-- +goose Up
-- +goose StatementBegin
create table rates
(
    id         serial primary key,
    market     varchar   not null,
    ask        float     not null check ( ask >= 0 ),
    bid        float     not null check ( ask >= 0 ),
    created_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists rates;
-- +goose StatementEnd
