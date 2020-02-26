-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE rsvp
ADD COLUMN guest_food text;

ALTER TABLE rsvp
ALTER COLUMN attending TYPE text;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE rsvp
DROP COLUMN guest_food;

ALTER TABLE rsvp
ALTER COLUMN attending TYPE bool USING CAST(attending AS bool);