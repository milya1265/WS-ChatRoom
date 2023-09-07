package user

import (
	"context"
	"database/sql"
	"log"
)

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id;"

	var userId int

	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&userId)
	if err != nil {
		return &User{}, err
	}

	user.Id = userId

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1;"

	var u User

	log.Println(email)
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
