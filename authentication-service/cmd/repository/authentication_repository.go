package repository

import (
	"authentication_service/cmd/models/domain"
	"context"
	"database/sql"
)

type Repository interface {
	Insert(ctx context.Context, user domain.User) (domain.User, error)
	EmailIsAvailable(ctx context.Context, email string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *repository {
	return &repository{db}
}

// Insert a new user into the database, and return the newly user inserted row
func (r *repository) Insert(ctx context.Context, user domain.User) (domain.User, error) {
	stmt := `INSERT INTO 
				users (id, fullname, email, address, city, province, mobile, password, role)
			VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9) 
			RETURNING 
				id, fullname, email, address, city, province, mobile, password, role, created_at, updated_at`

	row := r.db.QueryRowContext(ctx, stmt,
		user.ID,
		user.Fullname,
		user.Email,
		user.Address,
		user.City,
		user.Province,
		user.Mobile,
		user.Password,
		user.Role,
	)

	var newUser domain.User
	err := row.Scan(
		&newUser.ID,
		&newUser.Fullname,
		&newUser.Email,
		&newUser.Address,
		&newUser.City,
		&newUser.Province,
		&newUser.Mobile,
		&newUser.Password,
		&newUser.Role,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return newUser, nil
}

// Find by email
func (r *repository) EmailIsAvailable(ctx context.Context, email string) bool {
	var user domain.User
	stmt := `SELECT 
				id, email 
			FROM 
				users
			WHERE 
				email = $1 
			LIMIT 1`

	row := r.db.QueryRowContext(ctx, stmt, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
	)

	// If email is not available, return false
	if err != nil {
		return false
	}

	// If email is available, return true
	return true
}
