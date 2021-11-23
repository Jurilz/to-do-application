DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks(
    id serial PRIMARY KEY,
    label VARCHAR(255),
    date VARCHAR(255),
    done BOOLEAN,
     )