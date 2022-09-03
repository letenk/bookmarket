package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Fullname  string
	Email     string
	Address   string
	City      string
	Province  string
	Mobile    string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
