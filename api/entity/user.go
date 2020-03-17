package entity

import "time"

// User App user
type User struct {
	FirstName string
	LastName  string
	Birthday  time.Time
	Schedule  Schedule
	Workouts  []Workout
}
