-- +goose Up
CREATE TABLE projects (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    github_url TEXT NOT NULL,
    demo_url TEXT,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE projects;
