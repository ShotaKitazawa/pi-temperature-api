package domain

import "time"

// Output will be state of ArrCon
type Output struct {
	ID        int
	State     string
	CreatedAt *time.Time
}
