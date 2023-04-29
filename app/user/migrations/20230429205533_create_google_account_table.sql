-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS google_account (
    google_id varchar,
    user_profile_id integer,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,

    PRIMARY KEY(google_id),
    CONSTRAINT fk_user_profile
      FOREIGN KEY(user_profile_id) 
	  REFERENCES user_profile(user_profile_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE google_account 
-- +goose StatementEnd

