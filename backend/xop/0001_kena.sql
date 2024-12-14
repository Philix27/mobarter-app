-- +goose Up
CREATE TABLE kana (
    id INTEGER PRIMARY KEY,
    username text NOT NULL
);

-- +goose Down
DROP TABLE kana;