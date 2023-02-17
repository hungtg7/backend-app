-- +goose Up
-- +goose StatementBegin
INSERT INTO pet
SELECT generate_series(100,150), generate_series(100,150)::text, generate_series(100,150)::text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
