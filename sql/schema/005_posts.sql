-- +goose Up
CREATE TABLE posts (
	ID UUID PRIMARY KEY,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	title text NOT NULL,
	url TEXT UNIQUE NOT NULL,
	description TEXT NOT NULL,
	published_at TIMESTAMP NOT NULL,
	feed_id UUID NOT NULL,
	FOREIGN KEY(feed_id) REFERENCES feeds(id)
);


-- +goose Down
DROP TABLE posts;
