-- +goose Up
CREATE TABLE IF NOT EXISTS feed_follows(
    id UUID PRIMARY KEY,
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE IF EXISTS feed_follows;
