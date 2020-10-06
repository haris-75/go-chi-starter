package models

type User struct {
	Name string `json:"name"`
	Role int64  `json:"role"`
}

// USER Type
const (
	USER_ADMIN = iota
	USER_REGULAR
	NONUSER
)
