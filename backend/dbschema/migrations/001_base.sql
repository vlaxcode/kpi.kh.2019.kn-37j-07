-- +migrate Up
CREATE TABLE IF NOT EXISTS accounts
(
    id       BIGSERIAL PRIMARY KEY,
    name     VARCHAR(255),
    surname  VARCHAR(255),
    login    VARCHAR(255),
    password VARCHAR(255),
    sex      VARCHAR(255),
    date     VARCHAR(255)
);

-- +migrate Down
DROP TABLE IF EXISTS accounts;