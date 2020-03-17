package entity

import "time"

// Workout Daily workout of user
type Workout struct {
	Title     string
	Activity  Activity
	WorkedAt  time.Time
	CreatedAt time.Time
}
