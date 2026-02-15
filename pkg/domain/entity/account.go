package entity

import "time"

type Account struct {
	ID           int32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
