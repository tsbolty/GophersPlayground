CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO todos (task, done) VALUES
    ('Learn SQL', true),
    ('Learn GraphQL', false),
    ('Build a cool app', false);