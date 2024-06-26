-- +goose Up
CREATE TABLE IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL,
    filename TEXT NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

-- +goose Down
DROP TABLE IF EXISTS images;