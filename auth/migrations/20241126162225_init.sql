-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role as ENUM('admin', 'user');

CREATE TABLE IF NOT EXISTS USERS(
    id uuid primary key default gen_random_uuid(),
    first_name text not null,
    last_name text not null,
    email text unique not null,
    password text not null,
    role user_role not null default 'user',
    phone text,
    city text,
    country text,
    created_at timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS USERS;
DROP TYPE IF EXISTS user_role;
-- +goose StatementEnd
