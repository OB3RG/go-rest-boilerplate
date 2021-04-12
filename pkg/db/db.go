package db

import "database/sql"

// DB ...
type DB struct {
	Client *sql.DB
}

// Get ...
func Get(connStr string) (*DB, error) {
	db, err := get(connStr)

	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

// Close ...
func (d *DB) Close() error {
	return d.Client.Close()
}

func get(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
