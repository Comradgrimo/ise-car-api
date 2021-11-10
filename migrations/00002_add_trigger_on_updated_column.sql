-- +goose Up
CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER change_cars_updated_column
    BEFORE UPDATE ON car
    FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();

CREATE TRIGGER change_events_updated_column
    BEFORE UPDATE ON car
    FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();

-- +goose Down
DROP TRIGGER change_events_updated_column ON car;
DROP TRIGGER change_cars_updated_column ON car;

