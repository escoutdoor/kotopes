-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS PETS(
    id uuid primary key default gen_random_uuid(),
    name text not null,
    description text not null,
    age integer not null,
    owner_id uuid not null,
    created_at timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PETS;
-- +goose StatementEnd
