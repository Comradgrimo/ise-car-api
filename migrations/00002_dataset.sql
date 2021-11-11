-- +goose Up
insert into car(car_info, user_id, total_price, risk_rate) values
     ('Skoda Kodiaq', 1, 3200000, 1),
     ('Toyota Highlander', 1, 1200000, 1),
     ('Ford Focus', 1, 1000000, 2),
     ('Suzuki Jimny', 1, 1500000, 1),
     ('Volvo XC90', 1, 4500000, 1),
     ('Mazda CX-5', 1, 2499999, 1.5);

-- +goose Down
truncate table car cascade;
