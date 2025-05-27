package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, u User) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.QueryContext(ctx, "Select id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *repository) Create(ctx context.Context, u User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(name, email) VALUES($1, $2)", u.Name, u.Email)

	return err
}
