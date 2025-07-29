-- +goose Up
-- +goose StatementBegin
DELETE FROM projects;
ALTER TABLE projects ALTER COLUMN category_id SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
