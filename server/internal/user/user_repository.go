package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

// returns a pointer to a new repository
func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertID int
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	// scan the result of the query into the lastInsertID variable
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertID)

	if err != nil {
		return &User{}, err
	}

	// set the user's ID to the last inserted ID from the database
	// (this is the ID of the user we just created)
	user.ID = int64(lastInsertID)

	return user, nil
}
