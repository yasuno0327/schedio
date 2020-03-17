package entity

const (
	// Sun Define enum value of Sunday
	Sun = iota
	// Mon Define enum value of Monday
	Mon
	// Tue Define enum value of Tuesday
	Tue
	// Wed Define enum value of Wednesday
	Wed
	// Thu Define enum value of Thursday
	Thu
	// Fri Define enum value of Friday
	Fri
	// Sat Define enum value of Saturday
	Sat
)

// Schedule plan entity for week
type Schedule struct {
	Day        int
	Title      string
	Activities []Activity
	UserID     string
}
