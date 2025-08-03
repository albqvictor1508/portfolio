-- +goose Up
-- +goose StatementBegin
ALTER TABLE projects ADD COLUMN photo_url TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE projects DROP COLUMN photo_url;
-- +goose StatementEnd
