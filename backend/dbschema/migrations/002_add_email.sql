-- +migrate Up
ALTER TABLE IF EXISTS accounts
    ADD COLUMN email VARCHAR(255);

-- +migrate Down
ALTER TABLE IF EXISTS accounts
    DROP COLUMN email;