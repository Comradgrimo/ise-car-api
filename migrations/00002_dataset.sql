-- +goose Up
insert into car(car_info, user_id, total_price, risk_rate) values
     ('Skoda Kodiaq', 1, 3200000, 1),
     ('Toyota Highlander', 1, 1200000, 1),
     ('Ford Focus', 1, 1000000, 2),
     ('Suzuki Jimny', 1, 1500000, 1),
     ('Volvo XC90', 1, 4500000, 1),
     ('Mazda CX-5', 1, 2499999, 1.5);

insert into car_event(car_id, type, payload) values
     (1, 'created', '{"id":1,"car_info":"Skoda Kodiaq","user_id":1,"total_price":3200000,"risk_rate":1,"circs_link":""}'),
     (2, 'created', '{"id":2,"car_info":"Toyota Highlander","user_id":1,"total_price":1200000,"risk_rate":1,"circs_link":""}'),
     (3, 'created', '{"id":3,"car_info":"Ford Focus","user_id":1,"total_price":1000000,"risk_rate":2,"circs_link":""}'),
     (4, 'created', '{"id":4,"car_info":"Suzuki Jimny","user_id":1,"total_price":1500000,"risk_rate":1,"circs_link":""}'),
     (5, 'created', '{"id":5,"car_info":"Volvo XC90","user_id":1,"total_price":4500000,"risk_rate":1,"circs_link":""}'),
     (6, 'created', '{"id":6,"car_info":"Mazda CX-5","user_id":1,"total_price":2499999,"risk_rate":1.5,"circs_link":""}');

-- +goose Down
truncate table car cascade;
