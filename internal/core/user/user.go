package user

import "time"

type User struct {
	ID      int
	Balance float64
}

// TODO: Add transaction history and store user items))
type Transaction struct {
	ID           int
	UserID       int
	ItemHashName string
	Amount       float64
	CreatedAt    time.Time
}
