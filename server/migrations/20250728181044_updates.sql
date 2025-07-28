-- +goose Up
-- +goose StatementBegin
SELECT 'ALTER TABLE projects ADD COLUMN description TEXT NOT NULL';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
