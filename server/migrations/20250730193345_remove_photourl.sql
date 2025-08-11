-- +goose Up
-- +goose StatementBegin
ALTER TABLE technologies DROP COLUMN IF EXISTS photourl
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
