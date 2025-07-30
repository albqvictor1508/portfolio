-- +goose Up
-- +goose StatementBegin
ALTER TABLE technologies DROP COLUMN photourl
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
