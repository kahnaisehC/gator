-- +goose Up
CREATE TABLE feeds(
	name CHAR[1024],
	url CHAR[1024] UNIQUE,
	user_id UUID REFERENCES users ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds;
