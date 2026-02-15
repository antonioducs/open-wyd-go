package entity

import "time"

type Account struct {
	ID           uint32
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}
