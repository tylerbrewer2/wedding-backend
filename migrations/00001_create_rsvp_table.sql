-- +goose Up
CREATE TABLE rsvp (
    id SERIAL,
    email text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    attending boolean NOT NULL,
    food_choice text,
    guest_name text,
    note text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE rsvp;

