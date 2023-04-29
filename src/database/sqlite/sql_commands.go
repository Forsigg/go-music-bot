package sqlite

const (
	schemaSQL = `
CREATE TABLE IF NOT EXISTS users (
    id INT UNIQUE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    username VARCHAR(50)
);

CREATE INDEX IF NOT EXISTS users_id ON users(id);`
	addUserSQL = `
INSERT INTO users(
	id, first_name, last_name, username
) VALUES (
    ?, ?, ?, ?
);`
	allUsersSQL = `SELECT * FROM users;`
	userByIdSQL = `SELECT * FROM users WHERE id=? LIMIT 1`
)
