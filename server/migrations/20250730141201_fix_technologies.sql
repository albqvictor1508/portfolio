-- +goose Up
-- +goose StatementBegin
ALTER TABLE technologies ADD COLUMN photo_url TEXT NOT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
