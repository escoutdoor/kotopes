-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS FAVORITES(
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null,
    pet_id uuid not null,
    created_at timestamptz not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS FAVORITES;
-- +goose StatementEnd
