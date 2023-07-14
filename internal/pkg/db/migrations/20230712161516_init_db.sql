-- +goose Up
-- +goose StatementBegin
create table orders (
    id         BIGSERIAL PRIMARY KEY    NOT NULL,
    item_id bigint,
    user_id bigint,
    orderpoint_id bigint,
    orderstate varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
