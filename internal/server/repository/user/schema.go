package user

import "time"

type User struct {
	ID        int
	CreatedAt time.Time
	Login     string
	Hash      string
	Archived  bool
}
