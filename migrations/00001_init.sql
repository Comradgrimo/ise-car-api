-- +goose Up
CREATE TABLE car (
      id SERIAL PRIMARY KEY,
      car_info TEXT NOT NULL,
      user_id BIGINT NOT NULL,
      total_price DECIMAL,
      risk_rate NUMERIC,
      circs_link TEXT,
      removed BOOL NOT NULL DEFAULT FALSE,
      created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated TIMESTAMP
);

CREATE TABLE car_event (
     id SERIAL PRIMARY KEY,
     car_id BIGINT references car(id),
     type TEXT,
     status TEXT,
     payload jsonb,
     updated TIMESTAMP
);

-- +goose Down
DROP TABLE car_event;

DROP TABLE car;

