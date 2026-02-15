package entity

import "time"

type Account struct {
	ID           uint32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
