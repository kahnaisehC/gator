-- +goose Up
CREATE TABLE posts (
	id SERIAL PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL, 
	title VARCHAR(1024) NOT NULL,
	url VARCHAR(1024) NOT NULL UNIQUE,
	description TEXT,
	published_at TIMESTAMP,
	feed_id INTEGER REFERENCES feeds NOT NULL
);

-- +goose Down
DROP TABLE posts;
