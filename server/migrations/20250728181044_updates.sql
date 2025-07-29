-- +goose Up
-- +goose StatementBegin
 ALTER TABLE projects ADD COLUMN description TEXT NOT NULL;
 ALTER TABLE projects ALTER COLUMN github_url TYPE TEXT
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
