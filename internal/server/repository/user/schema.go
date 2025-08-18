package user

import "time"

// User database table struct
type User struct {
	ID        int
	CreatedAt time.Time
	Login     string
	Hash      string
	Archived  bool
}
