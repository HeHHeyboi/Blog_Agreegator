-- +goose Up
CREATE TABLE feed (
	name VARCHAR(50) NOT NULL,
	url VARCHAR(100) UNIQUE NOT NULL,
	user_id UUID NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed;
