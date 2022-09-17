package models

import (
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// User is the type for users
type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"-"`
	UserGroups []int     `json:"groups"`
}

// Group is the type for user groups
type UserGroup struct {
	ID        int       `json:"gid"`
	GroupName string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
