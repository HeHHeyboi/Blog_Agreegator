-- +goose Up
CREATE TABLE posts(
	id UUID PRIMARY KEY,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	title TEXT,
	url TEXT UNIQUE NOT NULL,
	description TEXT,
	published_at timestamp NOT NULL,
	feed_id UUID NOT NULL,
	FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
