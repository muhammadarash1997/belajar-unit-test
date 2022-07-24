package entity

import "time"

type Product struct {
	Id          int
	Description string
	CreatedAt   time.Time
}
