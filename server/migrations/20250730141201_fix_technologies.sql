-- +goose Up
-- +goose StatementBegin
ALTER TABLE technologies ADD COLUMN photoUrl TEXT NOT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
