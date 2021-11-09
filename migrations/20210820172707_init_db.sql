-- +goose Up
CREATE TABLE car (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL
);

-- +goose Down
DROP TABLE car;
