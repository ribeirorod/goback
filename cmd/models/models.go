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
	Groups    []Group   `json:"groups omitempty" gorm:"many2many:group_accounts"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Group is the type for user groups
type Group struct {
	ID          string    `json:"gid" gorm:"primaryKey"`
	Description string    `json:"-"`
	Members     []User    `json:"members" gorm:"many2many:group_accounts"`
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

func (Item) TableName() string {
	return "products"
}

func (g *User) BeforeCreate(tx *gorm.DB) error {
	g.ID = guuid.New().String()
	return nil
}

func (u *Group) BeforeCreate(tx *gorm.DB) error {
	u.ID = guuid.New().String()
	return nil
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = guuid.Must(guuid.NewRandom()).String()
	return nil
}

type Session struct {
	ID     string `json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id" gorm:"foreignKey:UserID"`
	ShoppingCart
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Shopping Cart is the type for shopping carts
type ShoppingCart struct {
	Items     []Item    `json:"items" gorm:"many2many:cart_items"`
	Voucher   []Voucher `json:"voucher" gorm:"many2many:cart_vouchers"`
	Subtotal  float64   `json:"subtotal"`
	Discount  float64   `json:"discount"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Date      time.Time `json:"date"`
	Items     []Item    `json:"items" gorm:"many2many:order_items"`
	Voucher   []Voucher `json:"voucher" gorm:"many2many:order_vouchers"`
	Subtotal  float64   `json:"subtotal"`
	Discount  float64   `json:"discount"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Item is the type for items in shopping cart
type Item struct {
	ItemID      string  `json:"id" gorm:"primaryKey"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float64 `json:"rate"`
	Quantity    int     `json:"quantity"`
}

// Voucher is the type for vouchers in shopping cart
type Voucher struct {
	VID        string    `json:"id" gorm:"primaryKey"`
	Discount   float64   `json:"discount"`
	ExpiryDate time.Time `json:"expiry_date"`
}
