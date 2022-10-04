package models

import (
	"context"
	"time"
)

// GORM Querys

// Query user by ID returns one user and error, if any
func (m *DBModel) GetUserByID(id string) (*User, error) {
	var user User

	if result := m.DB.Model(&user).First(&user, id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Query User by email, returns one user and error, if any
func (m *DBModel) GetUserByEmail(email string) (*User, error) {
	var user User

	result := m.DB.Model(&user).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *DBModel) UpdateOneUser(u *User) error {
	var user User
	result := m.DB.Model(&user).Updates(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *DBModel) CreateUser(u *User) error {
	var db = m.DB
	var user User
	result := db.Model(&user).Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SQL Querys
// Get returns one user and error, if any
func (m *DBModel) GetUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, _ := m.DB.DB()
	defer db.Close()

	query := `select id, email, username, password, created_on from public.accounts where id = $1`
	row := db.QueryRowContext(ctx, query, id)

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

	db, _ := m.DB.DB()
	defer db.Close()

	query := `update public.accounts set email = $1, username = $2, password = $3, updated_at = $4 where id = $5`
	_, err := db.ExecContext(ctx, query, user.Email, user.Username, user.Password, time.Now(), user.ID)
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

	db, _ := m.DB.DB()
	defer db.Close()
	_, err := db.ExecContext(ctx,
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
