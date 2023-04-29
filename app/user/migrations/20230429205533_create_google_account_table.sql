-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS google_account (
    user_profile_id varchar,
    google_id varchar,

    created_at timestamp
    updated_at timestamp
    deleted_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE google_account 
-- +goose StatementEnd

