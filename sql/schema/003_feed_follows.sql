-- +goose Up
CREATE TABLE feed_follows (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID REFERENCES users ON DELETE CASCADE NOT NULL,
	feed_id INTEGER REFERENCES feeds ON DELETE CASCADE NOT NULL,
	UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
