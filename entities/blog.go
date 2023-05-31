package entities

import "time"

type Blog struct {
	Id        int
	Title     string
	Date      time.Time
	Author    string
	Tags      string  // []string is alias for array
	Content   []uint8 // []byte is alias for longtext
	CreatedAt time.Time
	UpdatedAt time.Time
}
