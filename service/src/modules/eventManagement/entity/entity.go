package entity

import (
	"time"
)

type Organization struct {
	Id          uint64    `db:"id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Address     string    `db:"address"`
	PhoneNumber string    `db:"phone_number"`
	Logo        string    `db:"logo"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
