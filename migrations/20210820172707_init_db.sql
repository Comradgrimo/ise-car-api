-- +goose Up
CREATE TABLE car (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE car;
