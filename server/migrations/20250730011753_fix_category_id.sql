-- +goose Up
-- +goose StatementBegin
ALTER TABLE projects ALTER COLUMN category_id DROP NOT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
