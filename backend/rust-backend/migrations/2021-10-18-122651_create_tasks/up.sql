-- Your SQL goes here
CREATE TABLE "tasks"(
    id SERIAL PRIMARY KEY,
    label VARCHAR NOT NULL,
    date VARCHAR NOT NULL,
    done BOOLEAN NOT NULL DEFAULT 'f'
)