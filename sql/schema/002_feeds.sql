-- +goose Up
CREATE TABLE feeds(
	feed_id SERIAL PRIMARY KEY,
	name VARCHAR(1024) NOT NULL,
	url VARCHAR(1024) UNIQUE NOT NULL,
	user_id UUID REFERENCES users ON DELETE CASCADE NOT NULL,
	last_fetched_at TIMESTAMP
);
-- +goose Down
DROP TABLE feeds;
