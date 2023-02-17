-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pet (
    id integer
    name varchar
    pet_type varchar
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pet
-- +goose StatementEnd
