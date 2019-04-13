package domain

import "time"

// Input is Temperature & humidity data from DB
type Input struct {
	ID          uint
	Temperature float64
	Humidity    float64
	CreatedAt   time.Time
}
