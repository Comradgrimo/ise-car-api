-- +goose Up
CREATE TABLE car (
      id SERIAL PRIMARY KEY,
      car_info TEXT NOT NULL,
      user_id BIGINT NOT NULL,
      total_price DECIMAL NOT NULL,
      risk_rate NUMERIC NOT NULL,
      circs_link TEXT NOT NULL DEFAULT '',
      removed BOOL NOT NULL DEFAULT FALSE,
      created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE car_event (
     id SERIAL PRIMARY KEY,
     car_id BIGINT references car(id),
     type TEXT NOT NULL,
     status TEXT NOT NULL,
     payload jsonb NOT NULL,
     updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE car_event;

DROP TABLE car;

