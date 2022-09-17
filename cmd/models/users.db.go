package models

import (
	"context"
	"fmt"
	"time"
)

func (m *DBModel) GetUserByUsername(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Query user from DB
	query := `select id, email, username, password from public.accounts where email = $1`

	row := m.DB.QueryRowContext(ctx, query, email)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
	)

	fmt.Println("We go email ", user.Email)

	if err != nil {
		return nil, err
	}

	// get user groups, if any

	return &user, nil
}

// Get returns one user and error, if any
func (m *DBModel) GetUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, email, username, password, created_on from public.accounts where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// get user groups, if any

	return &user, nil
}

func (m *DBModel) UpdateUser(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update public.accounts set email = $1, username = $2, password = $3, updated_at = $4 where id = $5`

	_, err := m.DB.ExecContext(ctx, query, user.Email, user.Username, user.Password, time.Now(), user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) InsertUser(u *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into public.accounts 
		(email, username, phone, password, created_at, updated_at) 
		values ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx,
		query,
		u.Email,
		u.Username,
		u.Phone,
		u.Password,
		time.Now(),
		time.Now())

	if err != nil {
		return err
	}

	return nil
}
