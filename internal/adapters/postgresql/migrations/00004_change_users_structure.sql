-- +goose Up
UPDATE users SET first_name = 'Not assigned' WHERE first_name IS NULL;
UPDATE users SET last_name = 'Not assigned' WHERE last_name IS NULL;

ALTER TABLE users ALTER COLUMN first_name SET NOT NULL;
ALTER TABLE users ALTER COLUMN last_name SET NOT NULL;

-- +goose Down
ALTER TABLE users ALTER COLUMN first_name DROP NOT NULL;
ALTER TABLE users ALTER COLUMN last_name DROP NOT NULL;