// internal/repository/user.go
package repository

import (
	"database/sql"
	"sample-health/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	user := &model.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return user, nil
}
