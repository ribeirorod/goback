package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type DBModel struct {
	DB *gorm.DB
}

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModel(db *gorm.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// User is the type for users
type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Username  string    `json:"name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Groups    []Group   `json:"groups omitempty" gorm:"many2many:group_accounts;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Group is the type for user groups
type Group struct {
	ID          string    `json:"gid"`
	Description string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by the Model to `accouns`
func (User) TableName() string {
	return "accounts"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = guuid.New().String()
	return nil
}

func (u *Group) BeforeCreate(tx *gorm.DB) error {
	u.ID = guuid.New().String()
	return nil
}
