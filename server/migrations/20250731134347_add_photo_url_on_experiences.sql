-- +goose Up
-- +goose StatementBegin
ALTER TABLE experiences ADD COLUMN photo_url TEXT;
ALTER TABLE experiences ADD COLUMN category_id INTEGER REFERENCES categories(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE experiences DROP COLUMN category_id;
ALTER TABLE experiences DROP COLUMN photo_url;
-- +goose StatementEnd
